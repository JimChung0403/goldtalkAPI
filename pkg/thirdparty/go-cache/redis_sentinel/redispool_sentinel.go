package redis

import (
    "errors"
    "github.com/FZambia/sentinel"
    "github.com/gomodule/redigo/redis"
    "strings"
    "time"
)

var (
    sentinelPool *redis.Pool
    prekey       string
)

// Pool returns singleton instance
func SentinelPool() *redis.Pool {
    if sentinelPool == nil {
        //todo: 应该要返回 nil, err
    }
    return sentinelPool
}

func InitSentinelRedisPool(config Config) (error) {

    if config.MaxIdle <= 0 {
        config.MaxIdle = DefaultMaxIdle
    }

    if config.IdleTimeout <= 0 {
        config.IdleTimeout = DefaultIdleTimeout
    }

    if config.ConnTimeout <= 0 {
        config.ConnTimeout = DefaultConnTimeout
    }

    if config.SendTimeout <= 0 {
        config.SendTimeout = DefaultSendTimeout
    }

    if config.ReadTimeout <= 0 {
        config.ReadTimeout = DefaultReadTimeout
    }

    if config.MasterName == "" {
        config.MasterName = DefaultMasterName
    }

    sentinelPool = newSentinelPool(config)
    if sentinelPool == nil {
        return errors.New("InitSentinelRedisPool fail")
    }
    if strings.TrimSpace(config.Prefix) != "" {
        prekey = config.Prefix + "."
    } else {
        prekey = ""
    }
    return nil
}

func newSentinelPool(config Config) *redis.Pool {
    if config.PoolSize < minPoolSize {
        config.PoolSize = minPoolSize
    } else if config.PoolSize > maxPoolSize {
        config.PoolSize = maxPoolSize
    }
    sntnl := &sentinel.Sentinel{
        Addrs:      config.Addrs,
        MasterName: config.MasterName,
        Dial: func(addr string) (redis.Conn, error) {
            c, err := redis.DialTimeout("tcp", addr, config.ConnTimeout, config.ReadTimeout, config.SendTimeout)
            if err != nil {
                return nil, err
            }
            return c, nil
        },
    }
    return &redis.Pool{
        MaxIdle:     config.MaxIdle,
        MaxActive:   config.PoolSize,
        Wait:        true,
        IdleTimeout: config.IdleTimeout,
        Dial: func() (redis.Conn, error) {
            masterAddr, err := sntnl.MasterAddr()
            //masterAddr = "127.0.0.1:6379" //todo for dev
            if err != nil {
                return nil, err
            }
            c, err := redis.Dial("tcp", masterAddr)
            if err != nil {
                return nil, err
            }
            return c, nil
        },
        TestOnBorrow: func(c redis.Conn, t time.Time) error {
            if !sentinel.TestRole(c, "master") {
                return errors.New("Role check failed")
            } else {
                return nil
            }
        },
    }
}
