package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"redEnvelope/config"
)

// Conn 数据库连接句柄
var Conn *gorm.DB

func Init(dbConf config.Database) {
	if Conn != nil {
		return
	}

	dsn := dbConf.GenConnArgs()
	gormConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			// 全局禁用表名复数
			SingularTable: true,
		},
	}
	db, err := gorm.Open(mysql.Open(dsn), gormConfig)
	if err != nil {
		panic("Conn DB failed, err: " + err.Error())
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic("get sqlDB failed, err: " + err.Error())
	}
	err = sqlDB.Ping()
	if err != nil {
		panic("ping DB failed, err: " + err.Error())
	}

	Conn = db
}
