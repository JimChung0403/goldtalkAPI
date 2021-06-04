package client

import "github.com/go-resty/resty/v2"

var (
    PassportAPI *resty.Client
)

func InitPassportAPI(hostUrl string) *resty.Client {
    PassportAPI = resty.New()
    PassportAPI.HostURL = hostUrl
    PassportAPI.SetHeader("Content-Type", "application/json")
    PassportAPI.SetQueryParam("token", "6E86956E80B005D5C65774B24E593F89")
    return PassportAPI
}


