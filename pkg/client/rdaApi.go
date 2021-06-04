package client

import "github.com/go-resty/resty/v2"

var (
    RDAAPI *resty.Client
)

func InitRDAAPI(hostUrl string) *resty.Client {
    RDAAPI = resty.New()
    RDAAPI.HostURL = "http://tutorgroupapi.tutorabc.com/ReservationDataAccess"
    RDAAPI.SetHeader("Token", "88P%2fQvn3KFE%3d")
    RDAAPI.SetHeader("Content-Type", "application/json")
    return RDAAPI
}

