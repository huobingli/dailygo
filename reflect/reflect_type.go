package main

import (
	"fmt"
	"reflect"
)

//反射获取interface类型信息

func reflect_type(a interface{}) {
	t := reflect.TypeOf(a)
	fmt.Println("类型是：", t)
	// kind()可以获取具体类型
	k := t.Kind()
	fmt.Println(k)
	switch k {
	case reflect.Float64:
		fmt.Printf("a is float64\n")
	case reflect.String:
		fmt.Println("string")
	}
}

func reflect_value(a interface{}) {
	t := reflect.ValueOf(a)
	fmt.Println("类型是：", t)
	// kind()可以获取具体类型
	k := t.Kind()
	fmt.Println(k)
	switch k {
	case reflect.Float64:
		fmt.Printf("a is float64\n")
	case reflect.String:
		fmt.Println("string")
	}
}

func main() {
	var x float64 = 3.4
	x1 := "asdfafd"
	reflect_type(x)
	reflect_type(x1)

	y := 3.4
	reflect_value(y)
}
