package main

import (
	"fmt"
	"reflect"
)

// 定义结构体
type User struct {
	Id   int
	Name string
	Age  int
}

// 绑方法
func (u User) Hello() {
	fmt.Println("Hello")
}

func (u User) Echo() {
	fmt.Println("echo")
}

// 传入interface{}
func Poni(o interface{}) {
	t := reflect.TypeOf(o)
	fmt.Println("类型：", t)
	fmt.Println("字符串类型：", t.Name())
	// 获取值
	v := reflect.ValueOf(o)
	fmt.Println(v)
	// 可以获取所有属性
	// 获取结构体字段个数：t.NumField()
	for i := 0; i < t.NumField(); i++ {
		// 取每个字段
		f := t.Field(i)
		fmt.Printf("%s : %v", f.Name, f.Type)
		// 获取字段的值信息
		// Interface()：获取字段对应的值
		val := v.Field(i).Interface()
		fmt.Println("val :", val)
	}
	fmt.Println("=================方法====================")
	mn := t.NumMethod()
	fmt.Println("方法个数：", mn)
	for i := 0; i < mn; i++ {
		m := t.Method(i)
		fmt.Println(m.Name)
		fmt.Println(m.Type)
	}
}

func main() {
	u := User{1, "zs", 20}
	Poni(u)
}
