package database

import (
	"fmt"

	"github.com/hrz8/sc-masterlist-service/src/config"
	MysqlDriver "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlInterface interface {
	Connect() (*gorm.DB, error)
}

type mysql struct {
	appConfig config.AppConfig
}

func (m *mysql) Connect() (*gorm.DB, error) {
	DSN := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		m.appConfig.DB_USER,
		m.appConfig.DB_PASSWORD,
		m.appConfig.DB_HOST,
		m.appConfig.DB_PORT,
		m.appConfig.DB_NAME,
	)
	db, err := gorm.Open(MysqlDriver.Open(DSN), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func NewMysql(appConfig *config.AppConfig) MysqlInterface {
	return &mysql{
		appConfig: *appConfig,
	}
}
