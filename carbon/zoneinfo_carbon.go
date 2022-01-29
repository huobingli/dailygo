package main

import (
	"fmt"
	"log"
	"time"

	"github.com/uniplaces/carbon"
)

func main() {
	c, err := carbon.Create(2020, time.July, 24, 20, 0, 0, 0, "Asia/Shanghai")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", c)
}
