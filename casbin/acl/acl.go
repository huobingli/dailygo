package main

import (
	"fmt"
	"log"

	"github.com/casbin/casbin/v2"
)

func check(e *casbin.Enforcer, sub, obj, act string) {
	ok, _ := e.Enforce(sub, obj, act)
	if ok {
		fmt.Printf("%s CAN %s %s\n", sub, act, obj)
	} else {
		fmt.Printf("%s CANNOT %s %s\n", sub, act, obj)
	}
}

func main() {
	e, err := casbin.NewEnforcer("./model.conf", "./policy.csv")
	if err != nil {
		log.Fatalf("NewEnforecer failed:%v\n", err)
	}

	check(e, "hh", "data1", "read")
	check(e, "gg", "data2", "write")
	check(e, "hh", "data1", "write")
	check(e, "hh", "data2", "read")

	check(e, "root", "data1", "read")
	check(e, "root", "data2", "write")
	check(e, "root", "data1", "execute")
	check(e, "root", "data3", "rwx")
}
