package initializers

import (
	"github.com/Abdulrahman-Tayara/notes-app/users-service/configs"
	"github.com/spf13/viper"
)

func LoadConfig(path string, filename string) (config configs.Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(filename)
	viper.SetConfigType("env")

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	return
}

func LoadTestConfig(path string) (config configs.Config, err error) {
	return LoadConfig(path, "app")
}
