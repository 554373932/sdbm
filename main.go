package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "sdbm/routers"
)

type Sys_user struct {
	UserId   int    `orm:"pk"`
	Username string `orm:"size(100)"`
}

func main() {
	//beego.Run()
	orm.Debug = true
	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库
	user := new([]Sys_user)
	//err := o.Read(&user)
	//fmt.Println(err)
	//if err == orm.ErrNoRows {
	//
	//} else {
	//	fmt.Println(user.Username)
	//}
	num, err := o.QueryTable("sys_user").All(user)
	fmt.Println(num)
	fmt.Println(err)
	fmt.Println(user)

}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	// set default database
	orm.RegisterDataBase("default", "mysql", "root:root123@tcp(127.0.0.1:3306)/renren_security?charset=utf8", 30)

	orm.RegisterModel(new(Sys_user))
}
