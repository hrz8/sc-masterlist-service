package database

import (
	"fmt"

	"github.com/hrz8/sc-masterlist-service/src/config"
)

type MysqlInterface interface {
	Connect() string
}

type mysql struct {
	appConfig config.AppConfig
}

func (m *mysql) Connect() string {
	strConn := fmt.Sprintf("%s:%d", m.appConfig.DB_HOST, m.appConfig.DB_PORT)
	return strConn
}

func NewMysql(appConfig *config.AppConfig) MysqlInterface {
	return &mysql{
		appConfig: *appConfig,
	}
}
