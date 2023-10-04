package structurecheck

import "reflect"

func GetFirstFieldValue(obj interface{}) interface{} {
	v := reflect.ValueOf(obj)

	if v.Kind() == reflect.Struct {
		if v.NumField() > 0 {
			return v.Field(0).Interface()
		}
	}

	return nil
}
