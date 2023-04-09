package mysql

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"microHome/web/setting"
	"time"
)

var DB *gorm.DB

func Init(cfg *setting.MysqlConfig) (err error) {
	//连接mysql数据库
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.MysqlUser, cfg.MysqlPass, cfg.MysqlHost, cfg.MysqlPort, cfg.MysqlDB)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true, //跳过默认事务
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false, // 单数表名
			NoLowerCase:   false, // 关闭小写转换
		},
	})
	if err != nil {
		zap.L().Error("mysql连接失败", zap.Error(err))
		return
	}
	DB = db
	sqlDB, _ := DB.DB()
	sqlDB.SetMaxIdleConns(cfg.MysqlMaxConn)    //
	sqlDB.SetMaxOpenConns(cfg.MysqlMaxOpen)    //
	sqlDB.SetConnMaxLifetime(time.Second * 60) //设置连接的最大空闲时间
	zap.L().Info("mysql连接成功")
	Migration()
	return
}
func Close() (err error) {
	sqlDB, _ := DB.DB()
	err = sqlDB.Close()
	if err != nil {
		zap.L().Error("mySQL关闭失败", zap.Error(err))
		return
	}
	return
}
