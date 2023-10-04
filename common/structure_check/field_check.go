package structurecheck

import (
	"reflect"
)

func CheckField(p interface{}) bool {
	v := reflect.ValueOf(p)

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)

		if field.Interface() != reflect.Zero(field.Type()).Interface() {
			return false
		}
	}
	return true
}
