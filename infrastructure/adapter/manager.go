package adapter

import (
	"hex-architecture-boilerplate/infrastructure/adapter/api"
	"hex-architecture-boilerplate/infrastructure/adapter/config"
)

func Start() {
	configSrv := config.NewConfig()
	if configSrv.ApiActive() {
		api.Router()
	}
}
