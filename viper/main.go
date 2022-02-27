package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

func main()  {
	v := viper.New()
	v.SetConfigFile("config/config.yaml")
	err := v.ReadInConfig()
	if err != nil {
		log.Fatalln("读取文件失败")
		return
	}

	fmt.Println(v.Get("languages"))
}
