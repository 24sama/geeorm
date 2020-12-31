package dialect

import "reflect"

// 将go语言中的数据类型和数据库中的数据类型进行映射
// 屏蔽数据库差异

var dialectsmap = map[string]Dialect{}

type Dialect interface {
	DataTypeOf(typ reflect.Value) string                    // 用于将go语言的类型转换为该数据库的数据类型
	TableExistSQL(tableName string) (string, []interface{}) // 返回某个表是否存在的sql语句
}

func RegisterDialect(name string, dialect Dialect) {
	dialectsmap[name] = dialect
}

func GetDialect(name string) (dialect Dialect, ok bool) {
	dialect, ok = dialectsmap[name]
	return
}
