package configs

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var config = viper.New()

func Init() *viper.Viper {
	config.AddConfigPath("configs/")
	config.SetConfigName("application")
	config.SetConfigType("yaml")

	setConfigDefaults()

	if err := config.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Debug("Config file not found")
		} else {
			log.Warn("Config file was found but another error was produced")
		}
	}

	return config
}

func setConfigDefaults() *viper.Viper {
	// PostGres setup
	config.SetDefault("postgres.host", "localhost")
	config.SetDefault("postgres.port", "5432")
	config.SetDefault("postgres.database", "postgres")
	config.SetDefault("postgres.user", "postgres")
	config.SetDefault("postgres.password", "postgres")
	config.SetDefault("postgres.timezone", "America/Fortaleza")
	config.SetDefault("postgres.sslmode", "disable")
	// Redis setup
	config.SetDefault("redis.host", "localhost")
	config.SetDefault("redis.port", "6379")
	config.SetDefault("redis.password", "")
	config.SetDefault("redis.database", 0)
	// App setup
	config.SetDefault("app.host", "localhost")
	config.SetDefault("app.port", "8916")
	return config
}

func Config() *viper.Viper {
	return config
}

func GetPostgresDSN() string {
	return "host=" + config.GetString("postgres.host") +
		" user=" + config.GetString("postgres.user") +
		" password=" + config.GetString("postgres.password") +
		" dbname=" + config.GetString("postgres.database") +
		" port=" + config.GetString("postgres.port") +
		" sslmode=" + config.GetString("postgres.sslmode") +
		" TimeZone=" + config.GetString("postgres.timezone")
}

func GetRedisHost() string {
	return config.GetString("redis.host")
}

func GetRedisPort() string {
	return config.GetString("redis.port")
}

func GetRedisPassword() string {
	return config.GetString("redis.password")
}

func GetRedisDatabase() int {
	return config.GetInt("redis.database")
}

func GetAppHost() string {
	return config.GetString("app.host")
}

func GetAppPort() string {
	return config.GetString("app.port")
}

func GetAppUrl() string {
	return GetAppHost() + ":" + GetAppPort()
}
