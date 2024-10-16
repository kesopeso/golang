package main

/*
   extern int mini_calc(char *op, int a, int b);
*/
import "C"

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"unsafe"
)

type OrderInfo struct {
	OrderCode   rune
	Amount      int
	OrderNumber uint16
	Items       []string
	IsReady     bool
}

type SmallOrderInfo struct {
	Items       []string
	Amount      int
	OrderCode   rune
	OrderNumber uint16
	IsReady     bool
}

func main() {
	FieldLenCheck()
	CheckOrderInfoStruct()
	CheckSmallOrderInfoStruct()
	MiniCalc()
}

func MiniCalc() {
	val1 := 10
	val2 := 5
	op := "/"

	result := C.mini_calc(C.CString(op), C.int(val1), C.int(val2))
	fmt.Println("Result from c:", result)
}

func CheckSmallOrderInfoStruct() {
	oi := SmallOrderInfo{}
	ois := unsafe.Sizeof(oi)
	oit := unsafe.Offsetof(oi.Items)
	oa := unsafe.Offsetof(oi.Amount)
	oco := unsafe.Offsetof(oi.OrderCode)
	oon := unsafe.Offsetof(oi.OrderNumber)
	oir := unsafe.Offsetof(oi.IsReady)

	fmt.Println(ois, oit, oa, oco, oon, oir)
}

func CheckOrderInfoStruct() {
	oi := OrderInfo{}
	ois := unsafe.Sizeof(oi)
	oco := unsafe.Offsetof(oi.OrderCode)
	oa := unsafe.Offsetof(oi.Amount)
	oon := unsafe.Offsetof(oi.OrderNumber)
	oit := unsafe.Offsetof(oi.Items)
	oir := unsafe.Offsetof(oi.IsReady)

	//56 0 8 16 24 48
	fmt.Println(ois, oco, oa, oon, oit, oir)
}

func FieldLenCheck() {
	obj1 := ValidateStrLenStruct{"ba", "doesntmatter", "odye", true, "badbadba", 10}
	obj2 := ValidateStrLenStruct{"okok", "doesntmatter", "goodye", true, "badbadba", 10}
	obj3 := ValidateStrLenStruct{"okok", "doesntmatter", "goodye", true, "badidemobadba", 10}
	PorcessFieldLen(obj1)
	PorcessFieldLen(obj2)
	PorcessFieldLen(obj3)
}

func PorcessFieldLen(obj ValidateStrLenStruct) {
	err := ValidateStringLength(obj)
	if err != nil {
		fmt.Println("Invalid fields detected")
		fmt.Println(err)
		return
	}
	fmt.Println("Todo bien!")
}

type ValidateStrLenStruct struct {
	ToValidate1   string `minStrLen:"3"`
	DontValidate  string
	ToValidate2   string `minStrLen:"6"`
	DontValidate2 bool
	ToValidate3   string `minStrLen:"9"`
	DontValidate3 int
}

func ValidateStringLength(obj any) error {
	ot := reflect.TypeOf(obj)
	if ot.Kind() != reflect.Struct {
		return errors.New("only struct allowed as parameters")
	}

	var errs []error
	ov := reflect.ValueOf(obj)

	for i := 0; i < ot.NumField(); i++ {
		tf := ot.Field(i)
		if tf.Type.Kind() != reflect.String {
			continue
		}
		minLen := tf.Tag.Get("minStrLen")
		if minLen == "" {
			continue
		}
		minLenInt, err := strconv.Atoi(minLen)
		if err != nil {
			continue
		}
		vf := ov.Field(i)
		vfv := vf.String()
		vfvLen := len(vfv)
		if vfvLen < minLenInt {
			errs = append(errs, errors.New(fmt.Sprintf("field %s has value %s of length %d, min length is %d", tf.Name, vfv, vfvLen, minLenInt)))
		}
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}
	return nil
}
