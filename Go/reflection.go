package main

import (
	"fmt"
	"go/types"
	"reflect"
)

func main() {
	var x float64 = 3.4
	fmt.Println("type:", reflect.TypeOf(x))

	fmt.Println("type:", TypeOf(x))

	var y = [3]int{1, 2, 3}

	fmt.Println("type:", reflect.TypeOf(y))
	fmt.Println("type:", TypeOf(y))

	var z = map[string]int{"a": 1, "b": 2}
	fmt.Println("type:", reflect.TypeOf(z))
	fmt.Println("type:", TypeOf(z))

	var x2 float64 = 3.4
	v := reflect.ValueOf(x2)
	fmt.Println("type:", v.Type())
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	fmt.Println("value:", v.Float())

	fmt.Println("------")

	var x3 uint8 = 'x'
	v2 := reflect.ValueOf(x3)
	fmt.Println("type:", v2.Type())                            // uint8.
	fmt.Println("kind is uint8: ", v2.Kind() == reflect.Uint8) // true.
	x3 = uint8(v2.Uint())                                      // v.Uint returns a uint64.

	fmt.Println("------")

	fmt.Println("interface{} type:", v2.Interface().(uint8)) // interface{} type: uint8

}

func TypeOf(i interface{}) types.Type {
	t := reflect.TypeOf(i)
	if t == nil {
		return nil
	}
	switch t.Kind() {
	case reflect.Ptr:
		return types.NewPointer(TypeOf(reflect.Zero(t.Elem()).Interface()))
	case reflect.Slice:
		return types.NewSlice(TypeOf(reflect.Zero(t.Elem()).Interface()))
	case reflect.Map:
		return types.NewMap(TypeOf(reflect.Zero(t.Key()).Interface()), TypeOf(reflect.Zero(t.Elem()).Interface()))
	case reflect.Float64:
		return types.Typ[types.Float64]
	default:
		return nil
	}
}
