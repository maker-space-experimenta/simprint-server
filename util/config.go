package util

import (
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Port                      string `mapstructure:"PORT"`
	TempFileDir               string `mapstructure:"TEMPDIR"`
	FileDeleteDurationMinutes int    `mapstructure:"FILE_DELETE_DURATION_MINUTES"`

	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBSource      string `mapstructure:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)

	envFile, var_exists := os.LookupEnv("ENVFILE")
	if !var_exists {
		envFile = "app"
	}

	viper.SetConfigName(envFile)
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
