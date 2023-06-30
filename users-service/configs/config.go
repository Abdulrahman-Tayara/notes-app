package configs

import "github.com/spf13/viper"

type Config struct {
	AppEnv   string `mapstructure:"APP_ENV"`
	DbDSN    string `mapstructure:"DB_DSN"`
	Port     string `mapstructure:"PORT"`
	GRPCPort string `mapstructure:"GRPC_PORT"`
	GinMode  string `mapstructure:"GIN_MODE"`

	JWTKey                    string `mapstructure:"JWT_KEY"`
	JWTIssuer                 string `mapstructure:"JWT_ISSUER"`
	JWTAccessTokenAgeMinutes  int    `mapstructure:"JWT_ACCESS_TOKEN_AGE_MINUTES"`
	JWTRefreshTokenAgeMinutes int    `mapstructure:"JWT_REFRESH_TOKEN_AGE_MINUTES"`
}

var AppConfig *Config

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
