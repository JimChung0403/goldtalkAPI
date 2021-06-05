package redis

import (
	"errors"
	logger "scm.tutorabc.com/Research/goutil/log"
	"scm.tutorabc.com/Research/goutil/util"
	"github.com/FZambia/sentinel"
	"github.com/gomodule/redigo/redis"
	"time"
)

const (
	maxPoolSize = 128
)

var (
	sentinelPool *redis.Pool
	prekey       string
)

// Pool returns singleton instance
func SentinelPool() *redis.Pool {
	if sentinelPool == nil {
		logger.Error("sentinelPool is empty")
	}
	return sentinelPool
}

func InitSentinelRedisPool(addrs []string, masterName string, size int, prefix string) {
	sentinelPool = newSentinelPool(addrs, masterName, size)
	if !util.IsEmptyStr(prefix) {
		prekey = prefix + "."
	} else {
		prekey = ""
	}
}

func newSentinelPool(addrs []string, masterName string, size int) *redis.Pool {
	if size < 10 {
		size = 16
	} else if size > maxPoolSize {
		size = maxPoolSize
	}

	sntnl := &sentinel.Sentinel{
		Addrs:      addrs,
		MasterName: masterName,
		Dial: func(addr string) (redis.Conn, error) {
			timeout := 500 * time.Millisecond
			c, err := redis.DialTimeout("tcp", addr, timeout, timeout, timeout)
			if err != nil {
				return nil, err
			}
			return c, nil
		},
	}
	return &redis.Pool{
		MaxIdle:     8,
		MaxActive:   size,
		Wait:        true,
		IdleTimeout: 30 * time.Second,
		Dial: func() (redis.Conn, error) {
			masterAddr, err := sntnl.MasterAddr()
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
