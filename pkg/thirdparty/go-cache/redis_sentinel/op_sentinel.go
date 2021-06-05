package redis

import (
	"bytes"
	"encoding/gob"
	"errors"
	"github.com/gomodule/redigo/redis"
	"strings"
	"time"
)

const (
	op_GET    = "GET"
	op_SET    = "SET"
	op_SETEX  = "SETEX"
	op_DEL    = "DEL"
	op_GETSET = "GETSET"

	// list
	op_RPOP  = "RPOP"
	op_LPUSH = "LPUSH"
	op_LTRIM = "LTRIM"

	// sets operation
	op_SADD     = "SADD"
	op_SPOP     = "SPOP"
	op_SMEMBERS = "SMEMBERS"

	// sorted set
	op_ZADD             = "ZADD"
	op_ZCOUNT           = "ZCOUNT"
	op_ZLEXCOUNT        = "ZLEXCOUNT"
	op_ZRANGE           = "ZRANGE"
	op_ZRANGEBYSCORE    = "ZRANGEBYSCORE"
	op_ZRANK            = "ZRANK" // check exists
	op_ZREM             = "ZREM"
	op_ZREMRANGEBYSCORE = "ZREMRANGEBYSCORE"

	// hash
	op_HINCRBY = "HINCRBY"
	op_HMGET   = "HMGET"
	op_HMSET   = "HMSET"
	op_HGETALL = "HGETALL"

	// transaciton
	op_MULTI = "MULTI"
	op_EXEC  = "EXEC"

	//expire
	op_EXPIRE   = "EXPIRE"
	op_EXPIREAT = "EXPIREAT"
	op_PEXPIRE  = "PEXPIRE"

	op_EXIST = "EXISTS"

	NO_EXPIRED = time.Millisecond
)

func GetBytes(key string) ([]uint8, error) {
	con := SentinelPool().Get()
	defer con.Close()

	buf, err := redis.Bytes(con.Do(op_GET, prekey+key))
	if err == redis.ErrNil {
		return make([]uint8, 0), nil
	} else if err != nil {
		return nil, err
	}

	return buf, err
}

func GetString(key string) (string, error) {

	buf, err := GetBytes(key)
	if err != nil {
		return "", err
	}
	result := new(string)
	dec := gob.NewDecoder(bytes.NewBuffer(buf))

	if err = dec.Decode(result); err != nil {
		return "", err
	} else {
		return *result, nil
	}
}

func Get(key string, val interface{}) error {

	buf, err := GetBytes(key)
	if err != nil {
		return err
	}

	dec := gob.NewDecoder(bytes.NewBuffer(buf))

	if err = dec.Decode(val); err != nil {
		return err
	}

	return nil
}

func GetSet(key string, data interface{}) (interface{}, error) {
	if len(key) == 0 || data == nil {
		return nil, errors.New("key can't be empty, or data can't be nil")
	}
	con := SentinelPool().Get()
	defer con.Close()

	ret, err := con.Do(op_GETSET, prekey+key, data)
	return ret, err
}

func Sadd(key string, data interface{}) error {
	// check
	if len(key) == 0 || data == nil {
		return errors.New("key can't be empty, or data can't be nil")
	}
	con := SentinelPool().Get()
	defer con.Close()
	// _, err := redis.Int64(con.Do(op_SADD, prekey+key, data))
	_, err := con.Do(op_SADD, prekey+key, data)
	return err
}

func Spop(key string) (interface{}, error) {
	if len(key) == 0 {
		return nil, errors.New("spop key can't be empty")
	}
	con := SentinelPool().Get()
	defer con.Close()

	n, err := con.Do(op_SPOP, prekey+key)
	return n, err
}

func Smembers(key string) ([]interface{}, error) {
	if len(key) == 0 {
		return nil, errors.New("smembers key can't be empty")
	}
	con := SentinelPool().Get()
	defer con.Close()

	ms, err := con.Do(op_SMEMBERS, prekey+key)
	return ms.([]interface{}), err
}

func SetBytes(key string, data []uint8, expired time.Duration) error {

	if data == nil {
		return nil
	}
	con := SentinelPool().Get()
	defer con.Close()

	var err error
	result := ""

	if expired < time.Second {
		result, err = redis.String(con.Do(op_SET, prekey+key, data))
	} else {
		result, err = redis.String(con.Do(op_SETEX, prekey+key, int64(expired), data))
	}

	if err != nil {
		return err
	}

	if !strings.EqualFold("OK", result) {
		return errors.New(result)
	}
	return nil
}

func Set(key string, val interface{}, expired time.Duration) error {

	if val == nil {
		return nil
	}
	buffer := bytes.NewBuffer(nil)
	enc := gob.NewEncoder(buffer)
	if err := enc.Encode(val); err != nil {
		return errors.New("Set:encode:" + err.Error())
	}

	if err := SetBytes(key, buffer.Bytes(), expired); err != nil {
		return errors.New("Set:SetBytes:" + err.Error())
	}

	return nil
}

func Delete(key string) error {

	con := SentinelPool().Get()
	defer con.Close()

	_, err := redis.Int(con.Do(op_DEL, prekey+key))
	if err != nil {
		return errors.New("Delete:" + err.Error())
	}

	return nil
}

/*
	sorted set
	return:
		isAdd: true/false
		err
*/
func Zadd(key string, score int64, member string) (isAdd bool, err error) {
	conn := SentinelPool().Get()
	defer conn.Close()

	ret, err := conn.Do(op_ZADD, prekey+key, score, member)
	if err != nil {
		isAdd = false
		return isAdd, errors.New("Zadd:" + err.Error())
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

/*
min/max: number, -inf, +inf
*/
func Zcount(key string, min, max interface{}) (count int64, err error) {
	conn := SentinelPool().Get()
	defer conn.Close()

	n, err := conn.Do(op_ZCOUNT, prekey+key, min, max)
	if err != nil {
		err = errors.New("Zcount:" + err.Error())
		return
	}

	count = n.(int64)
	return
}

/*
	start: 0
	stop: -1
*/
func Zrange(key string, start, stop int, isWithScores bool) ([]interface{}, error) {
	conn := SentinelPool().Get()
	defer conn.Close()

	args := redis.Args{}.Add(prekey + key).Add(start).Add(stop)
	if isWithScores == true {
		args = args.Add("withscores")
	}

	ret, err := redis.Values(conn.Do(op_ZRANGE, args...))
	return ret, err
}

// sorted set: can check exist
func Zrank(key, member string) (int64, error) {
	conn := SentinelPool().Get()
	defer conn.Close()

	rank, err := conn.Do(op_ZRANK, prekey+key, member)
	if err != nil {
		return -1, errors.New("Zrank:" + err.Error())
	}

	if rank == nil {
		return -1, nil
	} else {
		return rank.(int64), nil
	}
}

// sorted set
func Zrem(key string, members ...string) error {
	con := SentinelPool().Get()
	defer con.Close()

	var err, receiveErr error
	for _, member := range members {
		err = con.Send(op_ZREM, prekey+key, member)
		if err != nil {
			return errors.New("Zrem send:" + err.Error())
		}
	}
	err = con.Flush()
	if err != nil {
		return errors.New("Zrem flush:" + err.Error())
	}
	return nil
}

// hash: hincrby
func Hincrby(key, field string, num interface{}) (interface{}, error) {
	conn := SentinelPool().Get()
	defer conn.Close()

	args := redis.Args{}.Add(prekey+key, field, num)
	ret, err := conn.Do(op_HINCRBY, args...)
	if err != nil {
		return nil, errors.New("Hincrby:" + err.Error())
	}

	return ret, nil
}

// hash: hmget
func Hmget(key string, fields ...interface{}) ([]interface{}, error) {
	conn := SentinelPool().Get()
	defer conn.Close()

	args := append([]interface{}{prekey + key}, fields...)
	ret, err := conn.Do(op_HMGET, args...)
	if err != nil {
		return nil, errors.New("Hmget:" + err.Error())
	}

	return ret.([]interface{}), nil
}

/*
	hash: hmset
	pairs: field, value pair
*/
func Hmset(key string, pairs ...interface{}) error {
	conn := SentinelPool().Get()
	defer conn.Close()

	args := append([]interface{}{prekey + key}, pairs...)
	_, err := conn.Do(op_HMSET, args...)
	if err != nil {
		return errors.New("Hmset:" + err.Error())
	}
	return nil
}

/*
	hash: hmset2
	pairs: field, value pair
*/
func Hmset2(key string, pair interface{}) error {
	conn := SentinelPool().Get()
	defer conn.Close()

	_, err := conn.Do(op_HMSET, redis.Args{}.Add(prekey+key).AddFlat(pair)...)
	if err != nil {
		return errors.New("Hmset:" + err.Error())
	}
	return nil
}

/*
	hash
	objPtr: object pointer
*/
func Hgetall(key string, objPtr interface{}) error {
	conn := SentinelPool().Get()
	defer conn.Close()

	// if not found, err still equals nil
	ret, err := redis.Values(conn.Do(op_HGETALL, prekey+key))
	if err != nil {
		return errors.New("Hgetall:" + err.Error())
	}

	err = redis.ScanStruct(ret, objPtr)
	if err != nil {
		return errors.New("Hgetall:" + err.Error())
	}

	return nil
}

// set expire
func PexpireSecond(key string, seconds int64) error {
	var errMsg string
	if key == "" {
		errMsg = "PexpireSecond: key can't be empty"
		return errors.New(errMsg)
	}

	if seconds <= 0 {
		errMsg = "PexpireSecond: seconds can't be set <= 0"
		return errors.New(errMsg)
	}

	conn := SentinelPool().Get()
	defer conn.Close()

	_, err := conn.Do(op_PEXPIRE, prekey+key, 1000*seconds)
	if err != nil {
		return errors.New("PexpireSecond:" + err.Error())
	}

	return nil
}

func init() {
	gob.Register(map[string]interface{}{})
}

func LPush(key string, value interface{}) (data int, err error) {
	conn := SentinelPool().Get()
	defer conn.Close()

	ret, err := redis.Int(conn.Do(op_LPUSH, prekey+key, value))

	if err != nil {
		return ret, errors.New("Hgetall:" + err.Error())
	}
	return ret, nil
}

func LPushBatch(key string, values interface{}) (data int, err error) {
	conn := SentinelPool().Get()
	defer conn.Close()
	args := redis.Args{}.Add(prekey + key).AddFlat(values)
	reply, err := redis.Int(conn.Do("LPUSH", args...))

	if err != nil {
		err = errors.New("LPushBatch:" + err.Error())
		return
	}
	data = reply
	return
}

func LTrim(key string, start, end int64) (err error) {
	conn := SentinelPool().Get()
	defer conn.Close()

	_, err = redis.String(conn.Do(op_LTRIM, prekey+key, start, end))
	if err != nil {
		err = errors.New("LTrim:" + err.Error())
		return
	}
	return
}

func IsExist(key string) (bool, error) {
	conn := SentinelPool().Get()
	defer conn.Close()
	var ret int
	var err error
	ret, err = redis.Int(conn.Do(op_EXIST, prekey+key))
	if err != nil {
		err = errors.New("Exist:" + err.Error())
		return false, err
	}
	return ret == 1, err
}
