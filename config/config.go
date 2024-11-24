package config

import "github.com/spf13/viper"

type Config struct {
	DBurl       string `mapstructure:"DBURL"`
	DBName      string `mapstructure:"DBNAME"`
	GrpcPort    string `mapstructure:"GRPCPORT"`
	KafkaPort   string `mapstructure:"KAFKA_BROKER"`
	AppEmail    string `mapstructure:"APPEMAIL"`
	AppPassword string `mapstructure:"APPPASSWORD"`
}

func LoadConfig() *Config {
	var config Config
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	viper.Unmarshal(&config)
	return &config
}
