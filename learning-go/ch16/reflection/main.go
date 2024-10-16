package main

import (
	"errors"
	"fmt"
	"reflect"
	"runtime"
	"strconv"
	"time"
)

type Foo struct {
	A int `myTag:"value"`
	B int `myTag:"value2"`
}

type MyData struct {
	Name   string `csv:"name"`
	Age    int    `csv:"age"`
	HasPet bool   `csv:"has_pet"`
}

// v should be an empty slice of MyData objects
func Unmarshal(data [][]string, v any) error {
	sliceValPointer := reflect.ValueOf(v)
	if sliceValPointer.Kind() != reflect.Pointer {
		return errors.New("second param should be of type slice of structs")
	}
	sliceVal := sliceValPointer.Elem()
	if sliceVal.Kind() != reflect.Slice {
		return errors.New("second param should be of type slice of structs")
	}
	structType := sliceVal.Type().Elem()
	if structType.Kind() != reflect.Struct {
		return errors.New("second param should be of type slice of structs")
	}

	if len(data) < 2 {
		return nil
	}

	csvMap := make(map[string]int, len(data[0]))
	for i, tag := range data[0] {
		csvMap[tag] = i
	}

	for i := 1; i < len(data); i++ {
		row := reflect.New(structType).Elem()
		for j := 0; j < structType.NumField(); j++ {
			if curTag, ok := structType.Field(j).Tag.Lookup("csv"); ok {
				valueIdx, ok := csvMap[curTag]
				if !ok {
					continue
				}
				valueToSet := data[i][valueIdx]
				rowField := row.Field(j)
				fieldKind := rowField.Kind()
				fieldName := structType.Field(j).Name
				switch fieldKind {
				case reflect.Int:
					intVal, err := strconv.ParseInt(valueToSet, 10, 64)
					if err != nil {
						return errors.New(fmt.Sprintf("invalid int value on line %d, tag %s", i, curTag))
					}
					rowField.SetInt(intVal)
				case reflect.Bool:
					boolVal, err := strconv.ParseBool(valueToSet)
					if err != nil {
						return errors.New(fmt.Sprintf("invalid bool value on line %d, tag %s", i, curTag))
					}
					rowField.SetBool(boolVal)
				case reflect.String:
					rowField.SetString(valueToSet)
				default:
					return errors.New(fmt.Sprintf("field %s has unsupported field type %s", fieldName, fieldKind))
				}
			}
		}
		sliceVal.Set(reflect.Append(sliceVal, row))
	}

	return nil
}

func Marshal(v any) ([][]string, error) {
	sliceVal := reflect.ValueOf(v)
	if sliceVal.Kind() != reflect.Slice {
		return nil, errors.New("must be a slice of structs")
	}
	structType := sliceVal.Type().Elem()
	if structType.Kind() != reflect.Struct {
		return nil, errors.New("must be a slice of structs")
	}
	var out [][]string
	header := marshalHeader(structType)
	out = append(out, header)
	for i := 0; i < sliceVal.Len(); i++ {
		row, err := marshalOne(sliceVal.Index(i))
		if err != nil {
			return nil, err
		}
		out = append(out, row)
	}
	return out, nil
}

func marshalHeader(structType reflect.Type) []string {
	var tags []string
	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		if curTag, ok := field.Tag.Lookup("csv"); ok {
			tags = append(tags, curTag)
		}
	}

	return tags
}

func marshalOne(stuctVal reflect.Value) ([]string, error) {
	var row []string
	structType := stuctVal.Type()
	fieldCount := stuctVal.NumField()
	for i := 0; i < fieldCount; i++ {
		fieldVal := stuctVal.Field(i)
		fieldKind := fieldVal.Kind()
		fieldName := structType.Field(i).Name
		switch fieldKind {
		case reflect.Int:
			strVal := strconv.Itoa(int(fieldVal.Int()))
			row = append(row, strVal)
		case reflect.Bool:
			boolVal := fieldVal.Bool()
			row = append(row, strconv.FormatBool(boolVal))
		case reflect.String:
			row = append(row, fieldVal.String())
		default:
			return nil, errors.New(fmt.Sprintf("%s field is of unsupported type %s", fieldName, fieldKind))
		}
	}
	return row, nil
}

func main() {
	// TagChecker()
	// CreateStringSliceAndAddStringWithReflection()
	//CsvMarshaler()

	//timeChecker := AddTiming(GetDate)
	//d := timeChecker(true)
	//fmt.Println(d)

	RunMemoizer()
}

func CsvMarshaler() {
	data := []MyData{
		{"Keso", 39, false},
		{"Reno", 38, false},
		{"Sifo", 37, true},
	}
	marshaled, err := Marshal(data)
	if err != nil {
		fmt.Println("marshaling failed", err)
	}
	fmt.Println("marshaled slice", marshaled)

	data2 := [][]string{
		{"name", "age", "has_pet"},
		{"Keso", "39", "false"},
		{"Reno", "38", "false"},
		{"Sifo", "37", "true"},
	}
	var unmarshaled []MyData
	err = Unmarshal(data2, &unmarshaled)
	if err != nil {
		fmt.Println("unmarshaling failed", err)
	}
	fmt.Println("unmarshaled slice", unmarshaled)
}

func TagChecker() {
	var f Foo
	ft := reflect.TypeOf(f)

	fields := map[string]reflect.StructField{}
	tags := map[string]string{}

	for i := 0; i < ft.NumField(); i++ {
		curField := ft.Field(i)
		fields[curField.Name] = curField
		tags[curField.Name] = curField.Tag.Get("myTag")
	}

	fmt.Println(fields, tags)
}

func CreateStringSliceAndAddStringWithReflection() {
	stringType := reflect.TypeOf((*string)(nil)).Elem()
	stringSliceType := reflect.TypeOf(([]string)(nil))

	stringSlice := reflect.MakeSlice(stringSliceType, 0, 10)
	stringVal := reflect.New(stringType).Elem()
	stringVal.SetString("somestring")
	stringSlice = reflect.Append(stringSlice, stringVal)
	stringSlice = reflect.Append(stringSlice, stringVal)
	stringSlice = reflect.Append(stringSlice, stringVal)

	stringSliceConverted := stringSlice.Interface().([]string)
	fmt.Println(stringSliceConverted)
}

func AddTiming[T any](f T) T {
	ft := reflect.TypeOf(f)
	if ft.Kind() != reflect.Func {
		panic("you can only pass functions to AddTiming wrapper")
	}
	fv := reflect.ValueOf(f)
	fo := runtime.FuncForPC(fv.Pointer())
	funcName := "<name>"
	if fo != nil {
		funcName = fo.Name()
	}

	wrapperF := reflect.MakeFunc(ft, func(in []reflect.Value) []reflect.Value {
		start := time.Now()
		out := fv.Call(in)
		duration := time.Since(start)
		fmt.Printf("Function %s was running for: %v\n", funcName, duration)
		return out
	})
	return wrapperF.Interface().(T)
}

func GetDate(real bool) time.Time {
	if !real {
		return time.Now()
	}

	return time.Now()
}

type outExp struct {
	out    []reflect.Value
	expiry time.Time
}

// buildInStruct creates a dynamic struct whose fields match the input parameters for the
// memoized function.
func buildInStruct(ft reflect.Type) (reflect.Type, error) {
	if ft.NumIn() == 0 {
		return nil, errors.New("must have at least one param")
	}
	// to create a dynamic struct, we create a slice of reflect.StructField
	sf := make([]reflect.StructField, 0, ft.NumIn())
	for i := 0; i < ft.NumIn(); i++ {
		ct := ft.In(i)
		// since this struct will be used as the key in a map, the struct must be comparable.
		// for a struct to be comparable, all of its fields must also be comparable.
		if !ct.Comparable() {
			return nil, fmt.Errorf("parameter %d of type %s and kind %v is not comparable", i+1, ct.Name(), ct.Kind())
		}
		// we add a struct field to sf for the input parameter,
		// making up a name and using the type from the input parameter
		sf = append(sf, reflect.StructField{
			Name: fmt.Sprintf("F%d", i),
			Type: ct,
		})
	}
	// this creates our dynamic struct type from our struct fields
	s := reflect.StructOf(sf)
	return s, nil
}

// Memoizer takes in a function and returns a wrapper function that caches the results of
// running the function for the specified duration.
//
// There are limitations on the functions that can be passed in to Memoizer.
//  1. The function should be long-running. Otherwise, there's no point in caching its results.
//  2. The function shouldn't have side effects. If it does, the side effects will only run when the
//     results for the provided parameters are not cached.
//  3. The input paramaters for the function must be comparable.
func Memoizer[T any](f T, expiration time.Duration) (T, error) {
	ft := reflect.TypeOf(f)
	if ft.Kind() != reflect.Func {
		var zero T
		return zero, errors.New("only for functions")
	}

	// we use a dynamic struct type to represent the input parameters for the function
	inType, err := buildInStruct(ft)
	if err != nil {
		var zero T
		return zero, err
	}

	if ft.NumOut() == 0 {
		var zero T
		return zero, errors.New("must have at least one returned value")
	}

	m := map[any]outExp{}
	fv := reflect.ValueOf(f)

	// we use the reflect.MakeFunc function to create a function with the same
	// input and output parameters as the provided function
	memo := reflect.MakeFunc(ft, func(args []reflect.Value) []reflect.Value {
		// create a key for our map
		iv := reflect.New(inType).Elem()
		for k, v := range args {
			iv.Field(k).Set(v)
		}
		ivv := iv.Interface()

		// check to see if the key is in the map and hasn't expired
		ov, ok := m[ivv]
		now := time.Now()
		if !ok || ov.expiry.Before(now) {
			// if the key isn't in the map, or the result has expired,
			// run the function and cache the results in the map
			ov.out = fv.Call(args)
			ov.expiry = now.Add(expiration)
			m[ivv] = ov
		}
		// return the value in the cache
		return ov.out
	})
	// return the memoized function
	return memo.Interface().(T), nil
}

func AddSlowly(a, b int) int {
	time.Sleep(100 * time.Millisecond)
	return a + b
}

func RunMemoizer() {
	addSlowly, err := Memoizer(AddSlowly, 2*time.Second)
	if err != nil {
		panic(err)
	}
	for i := 0; i < 5; i++ {
		start := time.Now()
		result := addSlowly(1, 2)
		end := time.Now()
		fmt.Println("got result", result, "in", end.Sub(start))
	}
	time.Sleep(3 * time.Second)
	start := time.Now()
	result := addSlowly(1, 2)
	end := time.Now()
	fmt.Println("got result", result, "in", end.Sub(start))
}
