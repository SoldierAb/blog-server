package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"runtime"
)

type MysqlDB struct {
	dbName string
	gorm *gorm.DB
}


type Config struct {
	User string
	Password string
	Dbname string
}

var mysqlDB *MysqlDB

func InitMysqlDB() error{

	db, err := gorm.Open("mysql", "root:123456@/blog?charset=utf8&parseTime=True&loc=Local")

	if err !=nil{
		fmt.Println("-----------------------> open err")
		return err
	}

	if db.DB().Ping() != nil{
		return err
	}

	if err = initAdmin(db); err != nil {
		return err
	}

	mysqlDB = &MysqlDB{
		gorm:   db,
		dbName: `blog`,
	}

	runtime.SetFinalizer(db,func(db *gorm.DB){
			db.Close()
	})

	return nil
}


//新连接
func NewConn() *gorm.DB{
	return mysqlDB.gorm.New()
}


