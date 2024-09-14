package main

import (
	"github.com/ryo-arima/payjp-proxy/pkg/client"
	"github.com/ryo-arima/payjp-proxy/pkg/config"
)

func main() {
	conf := config.NewBaseConfig()
	client.ClientForAppUser(conf)
}