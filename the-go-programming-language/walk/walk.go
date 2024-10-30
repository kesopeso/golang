package walk

import "reflect"

func Walk(r any, fn func(string, string)) {
	rValue := reflect.ValueOf(r)
	if rValue.Kind() != reflect.Struct {
		return
	}

	for i := 0; i < rValue.NumField(); i++ {
		fValue := rValue.Field(i)
		if fValue.Kind() != reflect.String {
			continue
		}
		fn(reflect.TypeOf(r).Field(i).Name, fValue.String())
	}
}
