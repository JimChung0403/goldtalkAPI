package client

import (
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"goldtalkAPI/conf"
	"goldtalkAPI/pkg/thirdparty/go-cache/redis_sentinel"
	"goldtalkAPI/pkg/util"
	"net/http"
	"strings"
)

func InitRedis(redisConf redis.Config) error {
	if conf.Conf.Cachecloud.AppID > 0 {
        d, err := GetCacheCloud(conf.Conf.Cachecloud.AppID)
        if err != nil{
        	return err
		}
		redisConf.Addrs = strings.Split(d.Sentinels, " ")
		redisConf.MasterName = d.Mastername
	}

	return redis.InitSentinelRedisPool(redisConf)
}


var urlFmt = "http://cachecloud.tutorabc.com/cache/client/redis/sentinel/%d.json?clientVersion=1.0"

type CacheCloud struct {
	Sentinels  string `json:"sentinels"`
	Message    string `json:"message"`
	Status     int    `json:"status"`
	Mastername string `json:"masterName"`
	Appid      int    `json:"appId"`
}


func GetCacheCloud(appID int64) (data *CacheCloud, err error) {
	url := fmt.Sprintf(urlFmt, appID)
	c := resty.New()
	resp, err := c.R().Get(url)
	if err != nil{
		return nil, err
	}
	if resp.StatusCode() != http.StatusOK{
		return nil, errors.New("cache cloud error")
	}
	err = util.JsonUnmarshalFromString(resp.String(), &data)
	if err != nil {
		return
	}
	return
}
