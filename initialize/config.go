package initialize

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)


type MysqlConfig struct {
	Host string `mapstructure:"host" json:"host,omitempty"`
	Port int    `mapstructure:"port" json:"port,omitempty"`
}


type Config struct {
	Name string `mapstructure:"name" json:"name,omitempty"`
	MysqlConfig MysqlConfig `mapstructure:"mysql_config" json:"mysql_config,omitempty"`
}

// 初始化全局变量

var GlobalConfig *Config = &Config{}

func GetEnvInfo(env string) bool  {
	viper.AutomaticEnv()
	return viper.GetBool(env)
}

// 初始化配置文件

func InitConfig()  {
	debug := GetEnvInfo("DEBUG")
	configPrefix := "config"
	configName := fmt.Sprintf("config/%s-prod.yaml",configPrefix)
	if debug {
		configName = fmt.Sprintf("config/%s-debug.yaml",configPrefix)
	}

	v := viper.New()
	v.SetConfigFile(configName)
	err := v.ReadInConfig()
	if err != nil {
		log.Fatalln(err)
		return
	}

	//var config config2.Config
	err = v.Unmarshal(&GlobalConfig)
	if err != nil {
		log.Fatalln(err)
		return
	}

	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		v.ReadInConfig()
		v.Unmarshal(&GlobalConfig)
	})
}