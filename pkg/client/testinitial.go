package client

import (
	"goldtalkAPI/conf"
	"goldtalkAPI/pkg/thirdparty/go-log"
)

func InitClients(config conf.Config) {
	log.Init(&config.Log)
	log.Init(&config.Log)

    InitRDAAPI("http://httpbin.org")
    InitPassportAPI("http://apitw.passport.tutorabc.com/web")
	InitSMS("http://sms.tutorabc.com/twapi")
	Setup(config.DB)
}
