package main

import (
	"fmt"
	"log"

	"github.com/casbin/casbin/v2"
)

func check(e *casbin.Enforcer, sub, domain, obj, act string) {
	ok, _ := e.Enforce(sub, domain, obj, act)
	if ok {
		fmt.Printf("%s CAN %s %s in %s\n", sub, act, obj, domain)
	} else {
		fmt.Printf("%s CANNOT %s %s in %s\n", sub, act, obj, domain)
	}
}

func main() {
	e, err := casbin.NewEnforcer("model.conf", "policy.csv")
	if err != nil {
		log.Fatalf("NewEnforecer failed:%v\n", err)
	}

	check(e, "hh", "tenant1", "data1", "read")
	check(e, "hh", "tenant1", "data1", "write")
	check(e, "gg", "tenant1", "data1", "read")
	check(e, "gg", "tenant1", "data1", "write")
	check(e, "gg", "tenant2", "data2", "read")
	check(e, "gg", "tenant2", "data2", "write")
}
