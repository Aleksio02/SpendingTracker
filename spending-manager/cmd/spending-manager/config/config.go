package config

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/spf13/viper"
	"os"
)

var Config appConfig

type appConfig struct {
	Application struct {
		Name    string
		Version string
		Port    int
	}
	Db struct {
		Name     string
		Username string
		Password string
		Host     string
		Port     int
	}
	Connector struct {
		Auth string
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

func CreateDatabaseConnection() (*pgx.Conn, error) {
	connString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		Config.Db.Username,
		Config.Db.Password,
		Config.Db.Host,
		Config.Db.Port,
		Config.Db.Name)
	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return conn, err
}
