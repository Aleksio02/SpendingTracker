package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var Config appConfig

type appConfig struct {
	Application struct {
		Name    string
		Version string
		Port    int
	}
}

func LoadConfig(configPaths ...string) error {
	v := viper.New()
	v.SetConfigName("application")
	v.SetConfigType("yml")
	v.AutomaticEnv()
	for _, path := range configPaths {
		v.AddConfigPath(path)
	}
	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to read the configuration file: %s", err)
	}
	return v.Unmarshal(&Config)
}
