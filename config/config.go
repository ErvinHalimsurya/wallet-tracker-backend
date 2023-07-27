package config

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

type Server struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}
type Database struct {
	DBHost            string `mapstructure:"db_host"`
	DBPort            string `mapstructure:"db_port"`
	DBName            string `mapstructure:"db_name"`
	DBUsername        string `mapstructure:"db_username"`
	DBPassword        string `mapstructure:"db_password"`
	DBPostgresSslMode string `mapstructure:"db_postgres_ssl_mode"`
}

type Jwt struct {
	Expired        string `mapstructure:"expired"`
	Secret         string `mapstructure:"secret"`
	RefreshExpired string `mapstructure:"refresh_expired"`
	RefreshSecret  string `mapstructure:"refresh_secret"`
	Issuer         string `mapstructure:"issuer"`
}

type Configuration struct {
	Server   `mapstructure:"server"`
	Database `mapstructure:"database"`
	Jwt      `mapstructure:"jwt"`
}

func Init() {
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Println("[Config][initConfig] Error reading config file: ", err)
		os.Exit(1)
	}

}

func GetDSN() string {
	var config Configuration

	err := viper.Unmarshal(&config)
	if err != nil {
		log.Println("[Config][InitDSNConfig] Unable to decode into struct:", err)
	}

	dsn := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=%v",
		config.DBUsername,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
		config.DBName,
		config.DBPostgresSslMode,
	)

	return dsn
}

func GetJWTConfig() Jwt {
	var config Configuration

	err := viper.Unmarshal(&config)
	if err != nil {
		log.Println("[Config][InitJWTConfig] Unable to decode into struct:", err)
	}

	return config.Jwt
}

func GetServerConfig() Server {
	var config Configuration

	err := viper.Unmarshal(&config)
	if err != nil {
		log.Println("[Config][InitServerConfig] Unable to decode into struct:", err)
	}

	return config.Server
}
