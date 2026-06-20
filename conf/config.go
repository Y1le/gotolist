package config

import (
	"os"

	"github.com/spf13/viper"
)

var Conf *Config

type Config struct {
	Server *Server `mapstructure:"service" yaml:"service"`
	MySQL  *MySQL  `mapstructure:"mysql" yaml:"mysql"`
	Redis  *Redis  `mapstructure:"redis" yaml:"redis"`
}

type Server struct {
	Port      string `mapstructure:"port" yaml:"port"`
	Version   string `mapstructure:"version" yaml:"version"`
	JwtSecret string `mapstructure:"jwtSecret" yaml:"jwtSecret"`
	Metrics   string `mapstructure:"metrics" yaml:"metrics"`
}

type MySQL struct {
	DriverName string `mapstructure:"driverName" yaml:"driverName"`
	Host       string `mapstructure:"host" yaml:"host"`
	Port       string `mapstructure:"port" yaml:"port"`
	Database   string `mapstructure:"database" yaml:"database"`
	UserName   string `mapstructure:"username" yaml:"username"`
	Password   string `mapstructure:"password" yaml:"password"`
	Charset    string `mapstructure:"charset" yaml:"charset"`
}

type Redis struct {
	RedisHost     string `mapstructure:"redisHost" yaml:"redisHost"`
	RedisPort     string `mapstructure:"redisPort" yaml:"redisPort"`
	RedisUsername string `mapstructure:"redisUsername" yaml:"redisUsername"`
	RedisPassword string `mapstructure:"redisPassword" yaml:"redisPassword"`
	RedisDbName   int    `mapstructure:"redisDbName" yaml:"redisDbName"`
}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(workDir + "/conf")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&Conf)
	if err != nil {
		panic(err)
	}
}
