package redis

import (
	"github.com/gomodule/redigo/redis"
	"strings"
	"sync"

	"errors"
	"fmt"
)

type RedisConnPoolWithPrekey struct {
	rcpool *RedisConnPool
	prekey string
}

var (
	_poolMap = make(map[string]*RedisConnPoolWithPrekey) // key: poolName
	initLock sync.Mutex
)

// GetPoolName by redisHost and Host
func GetPoolName(redisHost, prefix string) string {
	return redisHost + "@" + prefix
}

// InitRedisPool to initialize redis pool
func InitRedisMultiPool(redisHost string, size int, prefix string) {
	initLock.Lock()
	defer initLock.Unlock()

	poolName := GetPoolName(redisHost, prefix)
	if _, exists := _poolMap[poolName]; exists == true {
		return // already exists
	}
	pool := newPool(redisHost, size)
	keyPool := &RedisConnPoolWithPrekey{
		rcpool: pool,
	}


	if strings.TrimSpace(prefix) != "" {
		keyPool.prekey = prefix + "." // prekey: domain, without period
	}

	_poolMap[poolName] = keyPool
}

// for fetching rcpool and prekey
func GetPoolByName(name string) (*RedisConnPoolWithPrekey, error) {
	rcp, exists := _poolMap[name]
	if exists == false {
		return nil, errors.New(fmt.Sprintf("GetConnByName: pool [%v] not found.", name))
	}
	return rcp, nil
}

// GetConn returns a idled connection
func (rcp *RedisConnPoolWithPrekey) GetConn() redis.Conn {
	conn := rcp.rcpool.GetConn()
	return conn
}
