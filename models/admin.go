package models

import "github.com/jinzhu/gorm"

type Admin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u *Admin) TableName() string {
	return `admin`
}

func initAdmin(db *gorm.DB) error{
	return db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Admin{}).Error
}

//根据用户名查找
func (admin *Admin) GetUserByUsername() error{
	return NewConn().Table(admin.TableName()).Where("username = ? ",admin.Username).First(admin).Error
}

