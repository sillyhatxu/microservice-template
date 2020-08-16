package db

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
	"time"
)

// Model Struct
type User struct {
	Id               int64
	LoginName        string
	Password         string
	UserName         string
	Status           bool
	Platform         string
	Age              *int
	Amount           *float64
	Description      *string
	Birthday         *time.Time
	CreatedTime      time.Time
	LastModifiedTime time.Time
}

func InitialDatabase() {
	// register model
	orm.RegisterModel(new(User))

	// set default database
	//db.SetMaxOpenConns(10)
	//db.SetMaxIdleConns(5)
	//db.SetConnMaxLifetime(5*time.Minute)
	err := orm.RegisterDataBase("default", "mysql", "sillyhat_xu:sillyhat_xu_password@/sillyhat_xu_db?charset=utf8", 5, 10)
	if err != nil {
		panic(err)
	}

	// create table
	//orm.RunSyncdb("default", false, true)
}

func Insert(user User) error {
	o := orm.NewOrm()
	id, err := o.Insert(&user)
	if err != nil {
		return err
	}
	fmt.Println(id)
	return nil
}

func Update(user User) error {
	o := orm.NewOrm()
	num, err := o.Update(&user)
	if err != nil {
		return err
	}
	fmt.Println(num)
	return nil
}

func FindById(id int64) (*User, error) {
	o := orm.NewOrm()
	u := User{Id: id}
	err := o.Read(&u)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func TestORM() {
	o := orm.NewOrm()

	user := User{
		LoginName:        "LoginName",
		Password:         "Password",
		UserName:         "UserName",
		Status:           false,
		Platform:         "Platform",
		CreatedTime:      time.Time{},
		LastModifiedTime: time.Time{},
	}

	// insert
	id, err := o.Insert(&user)
	if err != nil {
		panic(err)
	}
	fmt.Println(id)
	// update
	user.LoginName = "LoginNameUpdate"
	user.Password = "PasswordUpdate"
	user.UserName = "UserNameUpdate"
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
