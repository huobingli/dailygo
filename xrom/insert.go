package main

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

type User struct {
	Id      int64
	Name    string
	Salt    string
	Age     int
	Passwd  string    `xorm:"varchar(200)"`
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}

func main() {
	engine, _ := xorm.NewEngine("mysql", "root:@/test?charset=utf8")
	user := &User{Name: "hbl", Age: 50}

	// 插入一条数据
	affected, _ := engine.Insert(user)
	fmt.Printf("%d records inserted, user.id:%d\n", affected, user.Id)

	// 插入多条数据
	users := make([]*User, 2)
	users[0] = &User{Name: "yxf", Age: 41}
	users[1] = &User{Name: "xjf", Age: 12}

	affected, _ = engine.Insert(&users)
	fmt.Printf("%d records inserted, id1:%d, id2:%d", affected, users[0].Id, users[1].Id)
}
