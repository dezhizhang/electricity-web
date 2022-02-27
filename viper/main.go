package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)



type MysqlConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type ServerConfig struct {
	ServerName string `mapstructure:"name"`
	MysqlConfig MysqlConfig `mapstructure:"mysql"`
}

func GetEnvInfo(env string) bool  {
	viper.AutomaticEnv()
	return viper.GetBool(env)
}

func main() {
	fmt.Println(GetEnvInfo("GOVERSION"))
	v := viper.New()

	err := v.ReadInConfig()
	if err != nil {
		log.Fatalln("读取文件失败")
		return
	}

	var serverConfig ServerConfig
	err = v.Unmarshal(&serverConfig)
	if err != nil {
		log.Fatalln("解析失败")
		return
	}
	
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		v.ReadInConfig()
		err = v.Unmarshal(&serverConfig)
		if err != nil {
			log.Fatalln("err")
			return
		}

		fmt.Println(serverConfig)

	})

	fmt.Println(v.Get("name"))
	fmt.Println(v.Get("mysql"))
}
