package client

import (
	"goldtalkAPI/conf"
	"goldtalkAPI/pkg/thirdparty/go-log"
)

func InitClients(config conf.Config) {
	log.Init(&config.Log)
	InitRedis(config.Redis)
    InitRDAAPI(config.APIHost.RDA)
    InitPassportAPI(config.APIHost.Passport)
	InitSMS(config.APIHost.SMS)
	//Setup(config.DB)
}
