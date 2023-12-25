package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	//"micro-service/go-common-utils/utils/logger/gorm_logger"
	"gin-template/conf"
	//"gorm.io/gorm/logger"
)

var TestDB *gorm.DB

func InitDB(config *conf.MySqlConfig, runMode string) {
	TestDB = setupMySqlDb(config.Test, runMode)
}

func setupMySqlDb(config *conf.MySqLDB, runMode string) *gorm.DB {
	c := &gorm.Config{}
	// 非线上环境开启log
	if runMode != "prod" {
		c = &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		}
	}
	db, err := gorm.Open(mysql.Open(config.DSN), c)
	if err != nil {
		panic(fmt.Errorf("Fatal error connection mysql: %s \n", err))
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(fmt.Errorf("Fatal error connection mysql: %s \n", err))
	}
	sqlDB.SetMaxOpenConns(config.MaxOpenConns)
	sqlDB.SetMaxIdleConns(config.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(config.ConnMaxLifetime)

	return db
}
