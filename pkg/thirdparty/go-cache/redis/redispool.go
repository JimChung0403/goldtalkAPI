package redis

import (
	"strings"
	"sync"
	"time"

	"github.com/gomodule/redigo/redis"
)

type RedisConnPool struct {
	pool   *redis.Pool
	locker *sync.Mutex
	host   string
}

const (
	maxPoolSize = 128
)

var (
	defaultPool    *RedisConnPool
	prekey         string
	defaultTimeout = 5000
	// poolMap = make(map[string]*RedisConnPool)
)

// Pool returns singleton instance
func Pool() *RedisConnPool {
	if defaultPool == nil {
		logger.Error("defaultPool is empty")
	}
	return defaultPool
}

// InitRedisPool to initialize redis pool
func InitRedisPool(redisHost string, size int, prefix string) {
	defaultPool = newPool(redisHost, size)
	if strings.TrimSpace(prefix) != "" {
		prekey = prefix + "."
	} else {
		prekey = ""
	}
}

func newPool(redisHost string, size int) *RedisConnPool {

	if size < 10 {
		size = 16
	} else if size > maxPoolSize {
		size = maxPoolSize
	}

	return &RedisConnPool{
		locker: new(sync.Mutex),
		host:   redisHost,
		pool:   initRedisPool(redisHost, size),
	}
}

func initRedisPool(host string, max int) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     8,
		MaxActive:   max,
		IdleTimeout: 30 * time.Second,
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")

			return err
		},
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", host,
				redis.DialConnectTimeout(time.Duration(defaultTimeout)*time.Millisecond),
				redis.DialWriteTimeout(time.Duration(defaultTimeout)*time.Millisecond),
				redis.DialReadTimeout(time.Duration(defaultTimeout)*time.Millisecond))
			if err != nil {
				return nil, err
			}

			return c, err
		},
	}
}

// GetConn returns a idled connection
func (rcp *RedisConnPool) GetConn() redis.Conn {

	var con redis.Conn
	rcp.locker.Lock()
	defer rcp.locker.Unlock()
	con = rcp.pool.Get()

	return con
}

func SetTimeOut(timeout int) {
	defaultTimeout = timeout
}
