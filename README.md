# golang微服务
### viper的使用
```go
type MysqlConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type ServerConfig struct {
	ServerName string `mapstructure:"name"`
	MysqlConfig MysqlConfig `mapstructure:"mysql"`
}

func main() {
	v := viper.New()
	v.SetConfigFile("config/config.yaml")
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

	fmt.Println(v.Get("name"))
	fmt.Println(v.Get("mysql"))
}
```
