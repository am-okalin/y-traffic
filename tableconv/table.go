package tableconv

import (
	"errors"
	"reflect"
)

var ObjError = errors.New("objs' type error")

func createStruct(t reflect.Type, m map[string]interface{}) reflect.Value {
	p := reflect.New(t)

	if t.Kind() == reflect.Struct {
		for k, v := range m {
			field := p.Elem().FieldByName(k)
			if field.IsValid() {
				field.Set(reflect.ValueOf(v))
			}
		}
	}
	return p.Elem()
}

func createStruct1(t reflect.Type, keys, values []string) reflect.Value {

	p := reflect.New(t)
	for i := 0; i < len(keys); i++ {
		field := p.Elem().FieldByName(keys[i])
		if field.IsValid() {
			field.Set(reflect.ValueOf(values[i]))
		}
	}
	return p.Elem()
}
