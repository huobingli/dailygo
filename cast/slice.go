package main

import (
	"fmt"

	"github.com/spf13/cast"
)

func main() {
	sliceOfInt := []int{1, 3, 7}
	arrayOfInt := [3]int{8, 12}
	// ToIntSlice
	fmt.Println(cast.ToIntSlice(sliceOfInt)) // [1 3 7]
	fmt.Println(cast.ToIntSlice(arrayOfInt)) // [8 12 0]

	sliceOfInterface := []interface{}{1, 2.0, "hahaha"}
	sliceOfString := []string{"abc", "hhh", "ggg"}
	stringFields := " abc  def hij   "
	any := interface{}(37)
	// ToStringSliceE
	fmt.Println(cast.ToStringSlice(sliceOfInterface)) // [1 2 hahaha]
	fmt.Println(cast.ToStringSlice(sliceOfString))    // [abc hhh ggg]
	fmt.Println(cast.ToStringSlice(stringFields))     // [abc def hij]
	fmt.Println(cast.ToStringSlice(any))              // [37]
}
