package schema

import (
	"geeorm/dialect"
	"go/ast"
	"reflect"
)

// 实现结构体和数据库表结构的映射

type Field struct {
	Name string
	Type string
	Tag  string
}

type Schema struct {
	Model      interface{}
	Name       string
	Fields     []*Field
	FieldNames []string
	fieldMap   map[string]*Field
}

func (schema *Schema) GetFiled(name string) *Field {
	return schema.fieldMap[name]
}

// 解析
func Parse(dest interface{}, d dialect.Dialect) *Schema {
	modelType := reflect.Indirect(reflect.ValueOf(dest)).Type()
	schema := &Schema{
		Model:    dest,
		Name:     modelType.Name(),
		fieldMap: make(map[string]*Field),
	}

	for i := 0; i < modelType.NumField(); i++ {
		p := modelType.Field(i)
		if !p.Anonymous && ast.IsExported(p.Name) {
			field := &Field{
				Name: p.Name,
				Type: d.DataTypeOf(reflect.Indirect(reflect.New(p.Type))),
			}
			if v, ok := p.Tag.Lookup("geeorm"); ok {
				field.Tag = v
			}
			schema.Fields = append(schema.Fields, field)
			schema.FieldNames = append(schema.FieldNames, p.Name)
			schema.fieldMap[p.Name] = field
		}
	}
	return schema
}

// 根据数据库中列的顺序，从对象中找到对应的值，按顺序平铺放入interface{}数组中。
// 因为sql语句格式为INSERT INTO user (name, age) VALUES ("Tom", 18)，需要字段名和字段的值顺序一致
func (schema *Schema) RecordValues(dest interface{}) []interface{} {
	destValue := reflect.Indirect(reflect.ValueOf(dest))
	var fieldsValues []interface{}
	for _, field := range schema.Fields {
		fieldsValues = append(fieldsValues, destValue.FieldByName(field.Name).Interface())
	}
	return fieldsValues
}
