package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type User struct {
	Username string `orm:"size(100)"`
	Department string
	Created string
}

func (u User) TableName() string {
	return "userinfo"
}

func main()  {
	//username :="root"
	//password :="123456"
	//host :="127.0.0.1"
	//port :=3306
	//Dbname :="test"
	//dsn :=fmt.Sprint("%s:%s@tcp(%s:%d)/%s?charset=utf8", username, password, host, port, Dbname)
	dsn :="root:123456@(127.0.0.1)/test?charset=utf8&parseTime=True&loc=Local"
	db,err :=gorm.Open("mysql",dsn)
	if err!=nil{
		panic("连接数据库失败,error="+err.Error())
	}
	defer db.Close()
	db.AutoMigrate(User{})
	//u:=User{
	//	Username: "Mark59",
	//	Department: "软件开发部",
	//	Created: "2020-12-13",
	//}
	//会自动生成SQL：insert into `userinfo`(`username`,`department`,`created`) values("Mark","软件开发部","2020-10-12")
	//if err := db.Create(&u).Error; err != nil {
	//	fmt.Println("插入失败",err)
	//	return
	//}
	//查询全部信息
	users :=[]User{}
	result :=db.Find(&users)
	for _,v :=range users{
		fmt.Println(v)
	}
	println(result.RowsAffected)
	u :=[]User{}
	isNotFound :=db.Where("username=?","Mark59").Find(&u).RecordNotFound()
	if isNotFound {
		fmt.Println("找不到记录")
		return
	}
	for _,v := range u{
		fmt.Println(v.Username,v.Department)
	}

	//db.Model(User{}).Where("username=?","Mark59").Update("department","销售部")
	db.Where(map[string]interface{}{"username":"mark59"}).First(&u)
	db.Where("username=?","mark59").Delete(User{})
}
