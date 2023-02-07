package change

import (
	"errors"
	"reflect"
)

//暂时只能处理一层的结构体, 多层的时候有需求再改
func SqlObjToObj(sqlObj any, obj any) error {
	objVal := reflect.ValueOf(obj)
	if objVal.Kind() != reflect.Ptr {
		return errors.New("must pass a pointer, not a value, to StructScan destination")
	}
	sqlVal := reflect.ValueOf(sqlObj)
	for sqlVal.Kind() == reflect.Ptr {
		sqlVal = sqlVal.Elem()
	}
	objElemVal := objVal.Elem()
	objElemType := objElemVal.Type()
	count := objElemVal.NumField()

	for i := 0; i < count; i++ {
		field := objElemVal.Field(i)
		fieldName := objElemType.Field(i).Name
		sqlField := sqlVal.FieldByName(fieldName)
		if fieldName[0] < 'A' || fieldName[0] > 'Z' || sqlField.IsZero() {
			continue
		}
		if sqlField.Kind() == field.Kind() {
			field.Set(sqlField)
		} else if sqlField.Field(0).Kind() == field.Kind() {
			field.Set(sqlField.Field(0))
		}
	}
	return nil
}
