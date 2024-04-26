package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Name string `mapstructure:"Name"`
		Mode string `mapstructure:"Mode"`
		Port int    `mapstructure:"Port"`
		//MachineId int64  `mapstructure:"MachineId"`
	}

	Auth struct {
		AccessSecret  string `mapstructure:"AccessSecret"`
		AccessExpire  int64  `mapstructure:"AccessExpire"`
		RefreshExpire int64  `mapstructure:"RefreshExpire"`
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
		Address []string `mapstructure:"Address"`
	}

	UserRpc struct {
		Key string `mapstructure:"Key"`
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
