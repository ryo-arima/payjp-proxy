package server

import "github.com/ryo-arima/payjp-proxy/pkg/config"

func Main(conf config.BaseConfig) {
	router := InitRouter(conf)
	router.Run(":8000")
}