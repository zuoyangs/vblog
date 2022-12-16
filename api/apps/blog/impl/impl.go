package impl

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"https://github.com/zuoyangs/vblog/api/conf"
)

type Impl struct {
	db *gorm.DB
}

func NewImpl() *Impl {
	return &Impl{}
}

dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

func (i Impl) Init() error{
	i.db = conf.C
}