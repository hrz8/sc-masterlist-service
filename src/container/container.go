package container

import (
	"github.com/fgrosse/goldi"
	Config "github.com/hrz8/sc-masterlist-service/src/config"
	Database "github.com/hrz8/sc-masterlist-service/src/database"
)

func NewAppContainer() *goldi.Container {
	goldiRegistry := goldi.NewTypeRegistry()
	goldiConfig := make(map[string]interface{})
	container := goldi.NewContainer(goldiRegistry, goldiConfig)

	appConfig := Config.NewConfig()

	container.InjectInstance("shared.config", appConfig)
	container.RegisterType("shared.mysql", Database.NewMysql, "@shared.config")

	return container
}
