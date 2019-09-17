package models

import (
	"github.com/jinzhu/gorm"
)

type Categories struct {
	ID int64 `gorm:"primary_key" json:"id" form:"id"`
	Created  string `json:"created" form:"created"`
	Owner string `json:"owner" form:"owner"`
	Name string `json:"name" form:"name"`
}

func initCategories(db *gorm.DB) error{
	return db.Set("gorm:table_options","ENGINE=InnoDB").AutoMigrate(&Categories{}).Error
}

func(*Categories) tableName() string{
	return "categories"
}

func (this *Categories) CheckName(db *gorm.DB) (int64,error){
	var count int64
	err := db.Table(this.tableName()).Where("name = ? ",this.Name).Count(&count).Error
	if err !=nil{
		return 0,err
	}
	return count,err
}

func (this *Categories) Insert(db *gorm.DB) error{
	return db.Table(this.tableName()).Create(this).Error
}

func (this *Categories) Delete(db *gorm.DB) error{
	return db.Table(this.tableName()).Delete("id = ?",this.ID).Error
}

func (this *Categories) Update(db *gorm.DB,params map[string]interface{}) error{
	return db.Table(this.tableName()).Where("id = ?",this.ID).Update(params).Error
}

func (this *Categories) List(db *gorm.DB) (list []*Categories,err error){
	err = db.Table(this.tableName()).Find(&list).Error
	return
}






