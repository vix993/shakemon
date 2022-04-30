package config

import (
	"sync"

	"github.com/spf13/viper"
)

const (
	FtAPIKey = "FT_API_KEY"
	Port     = "PORT"
)

var once sync.Once
var vipercfg = viper.New()
var currentConfig config

type config struct {
	FtAPIKey string
	Port     string
}

func Get() config {
	once.Do(func() {
		vipercfg.BindEnv(FtAPIKey)

		currentConfig.FtAPIKey = vipercfg.GetString(FtAPIKey)
		currentConfig.Port = viper.GetString(Port)
	})

	return currentConfig
}
