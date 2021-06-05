package redis

import (
	"errors"
	"github.com/gomodule/redigo/redis"
)

// append prekey in function
func (rcp *RedisConnPoolWithPrekey) MakeTransaction(paramArr [][]interface{}) ([]interface{}, error) {
	conn := rcp.GetConn()
	defer conn.Close()

	conn.Send(op_MULTI)
	for _, params := range paramArr {
		if len(params) >= 2 {
			if key, ok := params[1].(string); ok {
				params[1] = rcp.prekey + key // params[0]:op, params[1]:<prekey+key>
			}
		}
		conn.Send(params[0].(string), params[1:]...)
	}
	ret, err := redis.Values(conn.Do(op_EXEC))
	if err != nil {
		return nil, errors.New("MakeTransaction:" + err.Error())
	}

	return ret, nil
}

func (rcp *RedisConnPoolWithPrekey) Del(key string) (n int64, err error) {
	conn := rcp.GetConn()
	defer conn.Close()

	ret, err := conn.Do(op_DEL, rcp.prekey+key)
	if err != nil {
		return 0, errors.New("RedisConnPoolWithPrekey.Del:" + err.Error())
	}

	return ret.(int64), nil
}

func (rcp *RedisConnPoolWithPrekey) ExpireAt(key string, expireTs int64) error {
	conn := rcp.GetConn()
	defer conn.Close()

	_, err := conn.Do(op_EXPIREAT, rcp.prekey+key, expireTs)
	if err != nil {
		return err, errors.New("RedisConnPoolWithPrekey.ExpireAt:" + err.Error())

	}

	return nil
}

func (rcp *RedisConnPoolWithPrekey) Set(key string, value interface{}) error {
	conn := rcp.GetConn()
	defer conn.Close()

	_, err := conn.Do(op_SET, rcp.prekey+key, value)
	if err != nil {
		return errors.New("RedisConnPoolWithPrekey.Set:" + err.Error())
	}

	return err
}

func (rcp *RedisConnPoolWithPrekey) Get(key string) (interface{}, error) {
	conn := rcp.GetConn()
	defer conn.Close()

	ret, err := conn.Do(op_GET, rcp.prekey+key)
	if err != nil {
		return nil, errors.New("RedisConnPoolWithPrekey.Get:" + err.Error())
	}

	return ret, nil
}

func (rcp *RedisConnPoolWithPrekey) Hgetall(key string) ([]interface{}, error) {
	conn := rcp.GetConn()
	defer conn.Close()

	ret, err := redis.Values(conn.Do(op_HGETALL, rcp.prekey+key))
	if err != nil {
		return nil, errors.New("RedisConnPoolWithPrekey.Hgetall:" + err.Error())
	}
	return ret, nil
}

func (rcp *RedisConnPoolWithPrekey) Hmget(key string, args ...interface{}) ([]interface{}, error) {
	conn := rcp.GetConn()
	defer conn.Close()

	args2 := redis.Args{}.Add(rcp.prekey + key).Add(args...)
	ret, err := redis.Values(conn.Do(op_HMGET, args2...))
	if err != nil {
		return nil, errors.New("RedisConnPoolWithPrekey.Hmget:" + err.Error())
	}

	return ret, nil
}

func (rcp *RedisConnPoolWithPrekey) Hmset(key string, pairs ...interface{}) error {
	conn := rcp.GetConn()
	defer conn.Close()

	args := append([]interface{}{rcp.prekey + key}, pairs...)
	_, err := conn.Do(op_HMSET, args...)
	if err != nil {
		return errors.New("RedisConnPoolWithPrekey.Hmset:" + err.Error())
	}
	return nil
}

// sorted set
func (rcp *RedisConnPoolWithPrekey) Zadd(key string, score int64, member string) (isAdd bool, err error) {
	conn := rcp.GetConn()
	defer conn.Close()

	ret, err := conn.Do(op_ZADD, rcp.prekey+key, score, member)
	if err != nil {
		return false, errors.New("RedisConnPoolWithPrekey.Zadd:" + err.Error())
	} else {
		if ret.(int64) == 0 {
			isAdd = false
			return
		} else {
			isAdd = true
			return
		}
	}
}

func (rcp *RedisConnPoolWithPrekey) Zlexcount(key, min, max string) (count int64, err error) {
	conn := rcp.GetConn()
	defer conn.Close()

	args := redis.Args{}.Add(rcp.prekey+key, min, max)
	n, err := conn.Do(op_ZLEXCOUNT, args...)
	if err != nil {
		return 0, errors.New("RedisConnPoolWithPrekey.Zlexcount:" + err.Error())
	}

	return n.(int64), nil
}

func (rcp *RedisConnPoolWithPrekey) Zrange(key string, start, stop int, withScores bool) ([]interface{}, error) {
	conn := rcp.GetConn()
	defer conn.Close()

	args := redis.Args{}.Add(rcp.prekey+key, start, stop)
	if withScores == true {
		args = args.Add("withscores")
	}

	// improve: add limit
	ret, err := redis.Values(conn.Do(op_ZRANGE, args...))
	if err != nil {
		return nil, errors.New("RedisConnPoolWithPrekey.Zrange:" + err.Error())
	}
	return ret, nil
}

/*
	min, max: int64, -inf, +inf
	withScores:
		false: return {member}
		true: return {member, score}
*/
func (rcp *RedisConnPoolWithPrekey) Zrangebyscore(key string, min, max interface{}, withScores bool) ([]interface{}, error) {
	conn := rcp.GetConn()
	defer conn.Close()

	args := redis.Args{}.Add(rcp.prekey+key, min, max)
	if withScores == true {
		args = args.Add("withscores")
	}

	// improve: add limit
	ret, err := redis.Values(conn.Do(op_ZRANGEBYSCORE, args...))
	if err != nil {
		return nil, errors.New("RedisConnPoolWithPrekey.Zrangebyscore:" + err.Error())
	}
	return ret, nil
}

func (rcp *RedisConnPoolWithPrekey) Zrem(key string, args ...interface{}) (int64, error) {
	conn := rcp.GetConn()
	defer conn.Close()

	args2 := redis.Args{}.Add(rcp.prekey + key).Add(args...)
	ret, err := conn.Do(op_ZREM, args2...)
	if err != nil {
		return 0, errors.New("RedisConnPoolWithPrekey.Zrem:" + err.Error())
	}

	return ret.(int64), nil
}

func (rcp *RedisConnPoolWithPrekey) Zremrangebyscore(key string, min, max interface{}) (interface{}, error) {
	conn := rcp.GetConn()
	defer conn.Close()

	i, err := conn.Do(op_ZREMRANGEBYSCORE, rcp.prekey+key, min, max)
	if err != nil {
		return nil, errors.New("RedisConnPoolWithPrekey.Zremrangebyscore:" + err.Error())
	}

	return i, nil
}
