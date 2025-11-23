package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
	Redis    RedisConfig    `mapstructure:"redis"`
	JWT      JWTConfig      `mapstructure:"jwt"`
}

type ServerConfig struct {
	Port string `mapstructure:"port"`
	Host string `mapstructure:"host"`
}

type DatabaseConfig struct {
	URL string `mapstructure:"url"`
}

type RedisConfig struct {
	URL string `mapstructure:"url"`
}

type JWTConfig struct {
	Secret string `mapstructure:"secret"`
}

var AppConfig Config

func LoadConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")

	// Set defaults
	viper.SetDefault("server.port", "8080")
	viper.SetDefault("server.host", "0.0.0.0")
	viper.SetDefault("database.url", "postgres://postgres:postgres@localhost:5432/golang-boilerplate?sslmode=disable")
	viper.SetDefault("redis.url", "localhost:6379")
	viper.SetDefault("jwt.secret", "your-secret-key")

	// Enable reading from environment variables
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Warning: config file not found, using defaults and env vars: %v", err)
	}

	if err := viper.Unmarshal(&AppConfig); err != nil {
		return err
	}

	return nil
}