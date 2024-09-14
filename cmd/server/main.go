package main

import (
	"github.com/ryo-arima/payjp-proxy/pkg/config"
	"github.com/ryo-arima/payjp-proxy/pkg/server"
)

func main() {
	conf := config.NewBaseConfig()
	server.Main(conf)
}