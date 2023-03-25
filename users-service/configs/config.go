package configs

type Config struct {
	DbDSN   string `mapstructure:"DB_DSN"`
	Port    string `mapstructure:"PORT"`
	GinMode string `mapstructure:"GIN_MODE"`

	JWTKey                    string `mapstructure:"JWT_KEY"`
	JWTIssuer                 string `mapstructure:"JWT_ISSUER"`
	JWTAccessTokenAgeMinutes  int    `mapstructure:"JWT_ACCESS_TOKEN_AGE_MINUTES"`
	JWTRefreshTokenAgeMinutes int    `mapstructure:"JWT_REFRESH_TOKEN_AGE_MINUTES"`
}

var AppConfig *Config
