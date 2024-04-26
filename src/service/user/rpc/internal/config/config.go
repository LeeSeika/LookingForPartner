package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Name string `mapstructure:"Name"`
		Host string `mapstructure:"Host"`
		Port int    `mapstructure:"Port"`
	}
	Log struct {
		Level      string `mapstructure:"Level"`
		FileName   string `mapstructure:"Filename"`
		Mode       string `mapstructure:"Mode"`
		MaxAge     int    `mapstructure:"MaxAge"`
		MaxBackups int    `mapstructure:"MaxBackups"`
		MaxSize    int    `mapstructure:"MaxSize"`
	}
	Etcd struct {
		Host        []string `mapstructure:"Host"`
		DialTimeout int      `mapstructure:"DialTimeout"`
		TTL         int64    `mapstructure:"TTL"`
		Key         string   `mapstructure:"Key"`
	}
	Mysql struct {
		Database    string
		User        string
		Password    string
		Host        string
		Port        string
		MaxIdleConn int
		MaxOpenConn int
	}
	Snowflake struct {
		MachineID int64
	}
	CacheRedis struct {
		Host     []string
		PoolSize int
		DB       int
	}
}

func MustLoad(path string) Config {
	c := Config{}

	viper.SetConfigFile(path)
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error:%s\n", err))
	}
	err = viper.Unmarshal(&c)
	if err != nil {
		panic(fmt.Errorf("Fatal error:%s\n", err))
	}

	return c
}
