package main

import (
	"fmt"

	"github.com/reactivex/rxgo/v2"
)

// 第一次Observe()消耗了所有的数据，第二个就没有数据输出了。
func main() {
	ch := make(chan rxgo.Item)
	go func() {
		for i := 0; i < 3; i++ {
			ch <- rxgo.Of(i)
		}
		close(ch)
	}()

	observable := rxgo.FromChannel(ch)

	// 输出后 第二个输出已经没有数据
	for item := range observable.Observe() {
		fmt.Println(item.V)
	}

	for item := range observable.Observe() {
		fmt.Println(item.V)
	}
}
