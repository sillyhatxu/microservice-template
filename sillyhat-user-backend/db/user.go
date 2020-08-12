package db

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

// Model Struct
type User struct {
	Id   int    `orm:"auto"`
	Name string `orm:"size(100)"`
}

func init() {
	// register model
	orm.RegisterModel(new(User))

	// set default database
	orm.RegisterDataBase("default", "mysql", "root:root@/my_db?charset=utf8", 30)

	// create table
	orm.RunSyncdb("default", false, true)
}

func TestORM() {
	o := orm.NewOrm()

	user := User{Name: "slene"}

	// insert
	id, err := o.Insert(&user)
	if err != nil {
		panic(err)
	}
	fmt.Println(id)
	// update
	user.Name = "astaxie"
	num, err := o.Update(&user)
	if err != nil {
		panic(err)
	}
	fmt.Println(num)
	// read one
	u := User{Id: user.Id}
	err = o.Read(&u)
	// delete
	num, err = o.Delete(&u)
}
