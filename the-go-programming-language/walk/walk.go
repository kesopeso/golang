package walk

import "reflect"

func Walk(r any, fn func(string, string), namespace string) {
	rValue := reflect.ValueOf(r)

	if rValue.Kind() == reflect.Pointer {
		rValue = rValue.Elem()
		r = rValue.Interface()
	}

	if rValue.Kind() != reflect.Struct {
		return
	}

	for i := 0; i < rValue.NumField(); i++ {
		fValue := rValue.Field(i)
		fName := reflect.TypeOf(r).Field(i).Name
		if fValue.Kind() != reflect.String {
			Walk(fValue.Interface(), fn, fName)
			continue
		}
		if len(namespace) > 0 {
			fName = namespace + "." + fName
		}
		fn(fName, fValue.String())
	}
}

func WalkAlt(r any, fn func(string)) {
	rValue := getValue(r)

	switch rValue.Kind() {
	case reflect.Struct:
		for i := 0; i < rValue.NumField(); i++ {
			f := rValue.Field(i)
			WalkAlt(f.Interface(), fn)
		}
	case reflect.Chan:
		for v, ok := rValue.Recv(); ok; v, ok = rValue.Recv() {
			WalkAlt(v.Interface(), fn)
		}
	case reflect.String:
		fn(rValue.String())
	}
}

func getValue(r any) reflect.Value {
	rValue := reflect.ValueOf(r)

	if rValue.Kind() == reflect.Pointer {
		rValue = rValue.Elem()
	}

	return rValue
}
