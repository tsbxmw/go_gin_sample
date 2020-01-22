package config

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var (
	ALL_HERE = ""
)

func LoadConfig(model string) {
	configPath := "./project/config/other.json"
	if model == gin.ReleaseMode {
		configPath = "/config/other.json"
	}
	viper.SetConfigFile(configPath)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	ALL_HERE = viper.GetString("all_here")
}
