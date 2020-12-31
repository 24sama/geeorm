package session

import (
	"geeorm/clause"
)

// 1.通过多次调用clause.Set()构造sql的子句
// 2.调用一次clause.Build()按照传入顺序构造最终sql语句
func (s *Session) Insert(values ...interface{}) (int64, error) {
	recordValues := make([]interface{}, 0)
	for _, value := range values {
		table := s.Model(value).RefTable()
		s.clause.Set(clause.INSERT, table.Name, table.FieldNames)
		recordValues = append(recordValues, table.RecordValues(value))
	}

	s.clause.Set(clause.VALUES, recordValues)
	sql, vars := s.clause.Build(clause.INSERT, clause.VALUES)
	result, err := s.Raw(sql, vars...).Exec()
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

//func (s *Session) Find(values interface{}) error {
//	destSlice := reflect.Indirect(reflect.ValueOf(values))
//	destType := destSlice.Type().Elem()
//	table := s.Model(reflect.New(destType).Elem().Interface()).RefTable()
//
//	s.clause.Set(clause.SELECT, table.Name, table.FieldNames)
//	sql, vars := s.clause.Build(clause.SELECT, clause.WHERE, clause.ORDERBY, clause.LIMIT)
//	rows, err := s.Raw(sql, vars...).QueryRows()
//	if err != nil {
//		return err
//	}
//}
