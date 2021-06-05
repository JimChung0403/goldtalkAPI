package client

import (
    "goldtalkAPI/pkg/thirdparty/go-cache/redis_sentinel"
)



func InitRedis(config redis.Config) error {
    return redis.InitSentinelRedisPool(config)
}

