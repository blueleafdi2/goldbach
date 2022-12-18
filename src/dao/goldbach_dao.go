package dao

import (
	"git.garena.com/shopee/loan-service/credit_backend/Goldbach/src/model"
	_ "github.com/go-sql-driver/mysql" //千万不要忘记导入 默认会执行init初始化一些操作的
	"github.com/go-xorm/xorm"
	"github.com/sirupsen/logrus"
	"xorm.io/core"
)

var engine *xorm.Engine

func init() {
	var err error
	engine, err = xorm.NewEngine("mysql", "root:123456@(127.0.0.1:3306)/goldbach?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
	engine.ShowSQL(true)
	engine.Logger().SetLevel(core.LOG_DEBUG)
	engine.SetMaxOpenConns(10)
}

func Insert(goldbachs []*model.Goldbach) {
	count, err := engine.Table("goldbach").Insert(goldbachs)
	if int(count) != len(goldbachs) || err != nil {
		logrus.Errorf("Insert Goldbach err %s", err)
	}
}

func Close() {
	engine.Close()
}
