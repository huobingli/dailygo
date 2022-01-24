package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	// 省略部分代码

	fmt.Println("GOPATH: ", viper.Get("GOPATH"))
}

func init() {
	// 绑定环境变量
	viper.AutomaticEnv()
}