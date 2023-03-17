package initializers

import "github.com/spf13/viper"

type Config struct {
	DbDSN   string `mapstructure:"DB_DSN"`
	Port    string `mapstructure:"PORT"`
	GinMode string `mapstructure:"GIN_MODE"`
}

func LoadConfig(path string, filename string) (config Config, err error) {
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

func LoadTestConfig(path string) (config Config, err error) {
	return LoadConfig(path, "app")
}
