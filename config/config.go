package config

import (
	"fmt"
	"path"
	"runtime"

	"github.com/spf13/viper"
)

var config *viper.Viper

func init() {
	config = viper.New()
	config.AutomaticEnv()

	config.AddConfigPath(settingPath())
	config.SetConfigType("yaml")
	config.SetDefault("APP_ENV", "development")
	config.SetConfigName(config.GetString("APP_ENV"))

	err := config.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func settingPath() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Join(path.Dir(filename), "/settings")
}

func GetConfig() *viper.Viper {
	return config
}
