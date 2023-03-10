package util

import (
	"github.com/spf13/viper"
)

type Config struct {
	DbDriver      string `mapstructure:"POSTGRES_DRIVER"`
	DbSource      string `mapstructure:"DATABASE_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

func LoadConfig(path string) (cfg Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&cfg)
	return
}
