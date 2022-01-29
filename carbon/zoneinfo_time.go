package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	loc, err := time.LoadLocation("Japan")
	if err != nil {
		log.Fatal("failed to load location: ", err)
	}

	d := time.Date(2021, time.July, 24, 20, 0, 0, 0, loc)
	fmt.Printf("%s\n", d)
}
