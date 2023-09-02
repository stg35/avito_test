package config

import "github.com/spf13/viper"

type DBConfig struct {
	User     string `mapstructure:"user"`
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
}

type ServerConfig struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

type RedisConfig struct {
	Address string `mapstructure:"address"`
}

type Config struct {
	Server   *ServerConfig `mapstructure:"server"`
	Database *DBConfig     `mapstructure:"db"`
	Redis    *RedisConfig  `mapstructure:"redis"`
}

func NewConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath("configs")
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}
	return &config, nil
}
