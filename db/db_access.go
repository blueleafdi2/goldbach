package main

import (
	"fmt"
	"git.garena.com/shopee/loan-service/credit_backend/Goldbach/prime"
	_ "github.com/go-sql-driver/mysql" //千万不要忘记导入 默认会执行init初始化一些操作的
	"github.com/go-xorm/xorm"
	"xorm.io/core"
)

func main() {
	//执行一下两条命令  安装mysql引擎 以及 安装xorm库
	//go get -u github.com/go-sql-driver/mysql
	//go get github.com/go-xorm/xorm

	//1.创建数据库引擎对象
	engine, err := xorm.NewEngine("mysql", "root:wd@(127.0.0.1:3306)/goldbach?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
	//2.数据库引擎关闭
	defer engine.Close()
	//3.数据库引擎设置
	engine.ShowSQL(true)                     //设置显示sql语句  会在控制台当中展示
	engine.Logger().SetLevel(core.LOG_DEBUG) //设置日志级别
	engine.SetMaxOpenConns(10)               //设置最大链接数量

	//3.简单的一些使用 具体看手册
	//查询表的所有的数据
	session := engine.Table("goldbach")
	count, err := session.Count()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(count)

	/*//使用原生sql语句进行查询
	result, err := engine.Query("select * from user")
	if err != nil {
		panic(err.Error())
	}
	for key,value := range result {
		fmt.Println(key,value)
	}*/

	for num := int64(4); num <= 100; num += 2 {
		goldbach := prime.SumOfTwoPrimes(num)
		engine.Insert(goldbach)
	}

}
