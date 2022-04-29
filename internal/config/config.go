package config

import (
	"sync"

	"github.com/spf13/viper"
)

const (
	FtAPIKey = "FT_API_KEY"
)

var once sync.Once
var vipercfg = viper.New()
var currentConfig config

type config struct {
	FtAPIKey string
}

func Get() config {
	once.Do(func() {
		vipercfg.BindEnv(FtAPIKey)

		currentConfig.FtAPIKey = vipercfg.GetString(FtAPIKey)
	})

	return currentConfig
}
