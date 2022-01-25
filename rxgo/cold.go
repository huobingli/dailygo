package main

import (
	"context"
	"fmt"

	"github.com/reactivex/rxgo/v2"
)

// Cold Observable 第二个println可以再次输出，因为它创建的流是独立于每个观察者的。
// 即每次调用Observe()都创建一个新的 channel。
func main() {
	observable := rxgo.Defer([]rxgo.Producer{func(_ context.Context, ch chan<- rxgo.Item) {
		for i := 0; i < 3; i++ {
			ch <- rxgo.Of(i)
		}
	}})

	// 输出后 第二个可以输出数据
	for item := range observable.Observe() {
		fmt.Println(item.V)
	}

	for item := range observable.Observe() {
		fmt.Println(item.V)
	}
}
