package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

var DB *gorm.DB

func Init() (err error) {
	//连接mysql数据库
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		"root", "mysql", "127.0.0.1", "3306", "ihome")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true, //跳过默认事务
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false, // 单数表名
			NoLowerCase:   false, // 关闭小写转换
		},
	})
	if err != nil {
		return
	}
	DB = db
	sqlDB, _ := DB.DB()
	sqlDB.SetMaxIdleConns(20)                  //
	sqlDB.SetMaxOpenConns(100)                 //
	sqlDB.SetConnMaxLifetime(time.Second * 60) //设置连接的最大空闲时间
	return
}
func Close() (err error) {
	sqlDB, _ := DB.DB()
	err = sqlDB.Close()
	if err != nil {
		return
	}
	return
}
