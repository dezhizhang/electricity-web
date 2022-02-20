package zap

import "go.uber.org/zap"


//创建目志

func NewLogger() (*zap.Logger,error)  {
	config := zap.NewProductionConfig()
	config.OutputPaths = []string {
		"../logger/logger.log",

	}
	return config.Build()
}


