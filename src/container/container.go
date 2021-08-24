package container

import (
	"fmt"

	"github.com/fgrosse/goldi"
	Config "github.com/hrz8/sc-masterlist-service/src/config"
	"github.com/hrz8/sc-masterlist-service/src/database"
)

func NewContainer() *goldi.Container {
	goldiRegistry := goldi.NewTypeRegistry()
	goldiConfig := make(map[string]interface{})
	container := goldi.NewContainer(goldiRegistry, goldiConfig)

	appConfig, err := Config.NewConfig()
	if err != nil {
		fmt.Println(err)
	}

	container.InjectInstance("shared.config", appConfig)
	container.RegisterType("shared.mysql", database.NewMysql, "@shared.config")

	return container
}
