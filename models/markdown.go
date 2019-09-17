package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type MarkDown struct {
	Id 			uint   	`json:"id"  			gorm:"primary_key"`
	CreateTime 	time.Time  	`json:"create_time"		form:"create_time"`                   //创建时间
	Title 		string  `json:"title"           form:"title"`
	Content 	string 	`json:"content"         form:"content"`
}

func initMarkDown(db *gorm.DB) error{
	return	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&MarkDown{}).Error
}

func (mk *MarkDown) TableName() string{
	return `markdown`
}

//文章存储
func (mk *MarkDown) Add() error{
	return NewConn().Table(mk.TableName()).Create(mk).Error
}

//删除
func (mk *MarkDown) Del() error{
	return nil
}

//改
func (mk *MarkDown) Update() error{
	return nil
}

//查
func (mk *MarkDown) Get() (*[]MarkDown,error){
	var markdowns []MarkDown
	conn := NewConn().Table(mk.TableName())
	return &markdowns,conn.Error
}
