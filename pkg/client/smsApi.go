package client

import "github.com/go-resty/resty/v2"

var (
    SMSAPI *resty.Client
)

func InitSMS(hostUrl string) *resty.Client {
    SMSAPI = resty.New()
    SMSAPI.HostURL = hostUrl
    SMSAPI.SetHeader("Content-Type", "application/json")
    return SMSAPI
}

