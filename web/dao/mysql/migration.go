package mysql

import (
	"go.uber.org/zap"
	"microHome/web/model"
	"os"
)

// Migration 自动迁移数据库
func Migration() {
	err := DB.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(new(model.User),
			new(model.Area),
			new(model.House),
			new(model.HouseImage),
			new(model.OrderHouse),
			new(model.Facility))
	if err != nil {
		zap.L().Error("register table fail", zap.Error(err))
		os.Exit(0)
	}
	zap.L().Info("register table success")
}
