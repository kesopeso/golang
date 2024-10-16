package main

import (
	"ch16/unsafe/unexported"
	"fmt"
	"math/bits"
	"reflect"
	"runtime"
	"unsafe"
)

type Data struct {
	value  uint32
	label  [10]byte
	active bool
}

const dataSize = unsafe.Sizeof(Data{})

func main() {
	WrapInFunctionName(SizeOf, "SizeOf")()
	WrapInFunctionName(IsLE, "IsLE")()
	WrapInFunctionName(BytesFromDataToSliceToData, "BytesFromDataToSliceToData")()
	WrapInFunctionName(ModifyUnexported, "ModifyUnexported")()
}

func ModifyUnexported() {
	s := unexported.HasUnexported{
		A: "some data",
	}

	fmt.Println("is activated", s.IsActivated())

	st := reflect.TypeOf(s)
	bField, ok := st.FieldByName("b")
	if !ok {
		fmt.Println("Can't find field named 'b'")
		return
	}
	bOffset := bField.Offset
	stPt := unsafe.Pointer(&s)

	bVal := (*bool)(unsafe.Add(stPt, bOffset))
	*bVal = true
	fmt.Println("is activated", s.IsActivated())
}

func BytesFromDataToSliceToData() {
	var label [10]byte
	copy(label[:], []byte("labellabel"))
	data := Data{
		value:  0xAABBCCDD,
		label:  label,
		active: true,
	}
	if IsLE() {
		data.value = bits.ReverseBytes32(data.value)
	}

	// data array
	da := *(*[dataSize]byte)(unsafe.Pointer(&data))
	// data slice
	ds := unsafe.Slice((*byte)(unsafe.Pointer(&data)), dataSize)
	// print it
	fmt.Println(da, ds)

	// data from bytes slice
	sd := *(*Data)((unsafe.Pointer)(unsafe.SliceData(ds)))
	if IsLE() {
		sd.value = bits.ReverseBytes32(sd.value)

	}
	// data from bytes array
	ad := *(*Data)(unsafe.Pointer(&da))
	if IsLE() {
		ad.value = bits.ReverseBytes32(ad.value)

	}
	// print it
	fmt.Println(sd, ad, string(sd.label[:]), string(ad.label[:]))
}

func IsLE() bool {
	var x uint16 = 0xFF00
	xb := *(*[2]byte)(unsafe.Pointer(&x))
	isLE := xb[0] == 0x00 && xb[1] == 0xFF
	fmt.Println("is LittleEndian", isLE)
	return isLE
}

func SizeOf() {
	// 0 8 16 = 24 total
	type boolIntBool struct {
		b  bool
		i  int
		b2 bool
	}

	// 0 1 8 = 16 total
	type boolBoolInt struct {
		b  bool
		b2 bool
		i  int
	}

	// 0 8 9 = 16 total
	type intBoolBool struct {
		i  int
		b  bool
		b2 bool
	}

	fmt.Println(unsafe.Sizeof(boolIntBool{}), unsafe.Offsetof(boolIntBool{}.b), unsafe.Offsetof(boolIntBool{}.i), unsafe.Offsetof(boolIntBool{}.b2))
	fmt.Println(unsafe.Sizeof(boolBoolInt{}), unsafe.Offsetof(boolBoolInt{}.b), unsafe.Offsetof(boolBoolInt{}.b2), unsafe.Offsetof(boolBoolInt{}.i))
	fmt.Println(unsafe.Sizeof(intBoolBool{}), unsafe.Offsetof(boolBoolInt{}.i), unsafe.Offsetof(boolBoolInt{}.b), unsafe.Offsetof(boolBoolInt{}.b2))
}

func WrapInFunctionName[T any](f T, funcName string) T {
	ft := reflect.TypeOf(f)
	if ft.Kind() != reflect.Func {
		panic("you can only pass functions to AddTiming wrapper")
	}

	fv := reflect.ValueOf(f)
	fo := runtime.FuncForPC(fv.Pointer())
	if fo != nil {
		funcName = fo.Name()
	}

	wrappedFunction := reflect.MakeFunc(ft, func(in []reflect.Value) []reflect.Value {
		delimiter := ""
		for i := 0; i < len(funcName); i++ {
			delimiter += "="
		}
		fmt.Println(delimiter)
		fmt.Println(funcName)
		fmt.Println(delimiter)

		out := fv.Call(in)

		fmt.Println(delimiter)
		fmt.Println()

		return out
	})

	return wrappedFunction.Interface().(T)
}
