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
	op_INCR   = "INCR"
	op_INCRBY = "INCRBY"
	op_DECR   = "DECR"
	op_DECRBY = "DECRBY"

	// list
	op_RPOP       = "RPOP"
	op_LPUSH      = "LPUSH"
	op_LTRIM      = "LTRIM"
	op_BRPOP      = "BRPOP"
	op_BRPOPLPUSH = "BRPOPLPUSH"
	op_LREM       = "LREM"

	// sets operation
	op_SADD     = "SADD"
	op_SPOP     = "SPOP"
	op_SREM     = "SREM"
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
	op_HGET    = "HGET"
	op_HDEl    = "HDEL"
	op_HDLEN   = "HLEN"

	// transaciton
	op_MULTI = "MULTI"
	op_EXEC  = "EXEC"

	//expire
	op_EXPIRE   = "EXPIRE"
	op_EXPIREAT = "EXPIREAT"
	op_PEXPIRE  = "PEXPIRE"
	op_PERSIST  = "PERSIST"

	op_EXIST = "EXISTS"

	DEFAULT_EXPIRE_DURATION = 6 * time.Hour

	NO_EXPIRED = time.Millisecond
)

var (
	Default_Expire_Seconds      = int(DEFAULT_EXPIRE_DURATION.Seconds())
	Default_Expire_Milliseconds = Default_Expire_Seconds * 1000
)

func init() {
	gob.Register(map[string]interface{}{})
}

func IsExist(key string) (bool, error) {
	conn := Pool().GetConn()
	defer conn.Close()
	var ret int
	var err error
	ret, err = redis.Int(conn.Do(op_EXIST, prekey+key))
	if err != nil {
		return false, err
	}
	return ret == 1, err
}

func GetBytes(key string) ([]uint8, error) {
	con := Pool().GetConn()
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
	con := Pool().GetConn()
	defer con.Close()

	ret, err := con.Do(op_GETSET, prekey+key, data)
	return ret, err
}

func Sadd(key string, data interface{}) error {
	// check
	if len(key) == 0 || data == nil {
		return errors.New("key can't be empty, or data can't be nil")
	}
	con := Pool().GetConn()
	defer con.Close()
	// _, err := redis.Int64(con.Do(op_SADD, prekey+key, data))
	_, err := con.Do(op_SADD, prekey+key, data)
	return err
}

func Spop(key string) (interface{}, error) {
	if len(key) == 0 {
		return nil, errors.New("spop key can't be empty")
	}
	con := Pool().GetConn()
	defer con.Close()

	n, err := con.Do(op_SPOP, prekey+key)
	return n, err
}

func Srem(key string, data interface{}) (interface{}, error) {
	if len(key) == 0 {
		return nil, errors.New("srem key can't be empty")
	}
	con := Pool().GetConn()
	defer con.Close()

	n, err := con.Do(op_SREM, prekey+key, data)
	return n, err
}

func Smembers(key string) ([]interface{}, error) {
	if len(key) == 0 {
		return nil, errors.New("smembers key can't be empty")
	}
	con := Pool().GetConn()
	defer con.Close()

	ms, err := con.Do(op_SMEMBERS, prekey+key)
	return ms.([]interface{}), err
}

func SetBytes(key string, data []uint8, expired time.Duration) error {

	if data == nil {
		return nil
	}
	con := Pool().GetConn()
	defer con.Close()

	var err error
	result := ""

	exp := int(expired.Seconds())
	if expired < time.Second {
		exp = Default_Expire_Seconds
	}

	result, err = redis.String(con.Do(op_SETEX, prekey+key, exp, data))
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
		logger.Error("Set:encode:", err)
		return err
	}

	if err := SetBytes(key, buffer.Bytes(), expired); err != nil {
		logger.Error("Set:SetBytes:", err)
		return err
	}

	return nil
}

func Delete(key string) error {

	con := Pool().GetConn()
	defer con.Close()

	_, err := redis.Int(con.Do(op_DEL, prekey+key))
	if err != nil {
		logger.Error("Delete:", err)
		return err
	}

	return nil
}

//region sorted set
/*
	sorted set
	return:
		isAdd: true/false
		err
*/
func Zadd(key string, score int64, member string) (isAdd bool, err error) {
	conn := Pool().GetConn()
	defer conn.Close()

	ret, err := conn.Do(op_ZADD, prekey+key, score, member)
	if err != nil {
		logger.Error("Zadd:", err)
		isAdd = false
		return
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
	conn := Pool().GetConn()
	defer conn.Close()

	n, err := conn.Do(op_ZCOUNT, prekey+key, min, max)
	if err != nil {
		logger.Error("Zcount:", err)
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
	conn := Pool().GetConn()
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
	conn := Pool().GetConn()
	defer conn.Close()

	rank, err := conn.Do(op_ZRANK, prekey+key, member)
	if err != nil {
		logger.Error("Zrank:", err)
		return -1, err
	}

	if rank == nil {
		return -1, nil
	} else {
		return rank.(int64), nil
	}
}

// sorted set
func Zrem(key string, members ...string) error {
	con := Pool().GetConn()
	defer con.Close()

	var err, receiveErr error
	for _, member := range members {
		err = con.Send(op_ZREM, prekey+key, member)
		if err != nil {
			logger.Error("Zrem send:", err)
			return err
		}
	}
	err = con.Flush()
	if err != nil {
		logger.Error("Zrem flush:", err)
		return err
	}
	return nil

	for i := 0; i < len(members); i++ {
		_, err = con.Receive()
		if err != nil {
			logger.Error("Zrem receive:", err)
			if receiveErr == nil {
				receiveErr = err
			}
		}
	}
	return receiveErr
}

func ZRangeByScore(key string, min int64, max int64) ([]string, error) {
	conn := Pool().GetConn()
	defer conn.Close()

	reply, err := redis.Strings(conn.Do(op_ZRANGEBYSCORE, prekey+key, min, max))
	if err != nil {
		logger.Error(op_ZRANGEBYSCORE, err)
		return nil, err
	}
	return reply, nil
}

func ZAddBatch(key string, members interface{}) (err error) {
	conn := Pool().GetConn()
	defer conn.Close()

	val := redis.Args{}.Add(prekey + key).AddFlat(members)

	_, err = conn.Do(op_ZADD, val...)
	if err != nil {
		logger.Error(err)
	}
	return
}

//endregion

//region hash op
func Hincrby(key, field string, num interface{}) (interface{}, error) {
	conn := Pool().GetConn()
	defer conn.Close()

	args := redis.Args{}.Add(prekey+key, field, num)
	ret, err := conn.Do(op_HINCRBY, args...)
	if err != nil {
		logger.Error("Hincrby:", err)
		return nil, err
	}

	return ret, nil
}

func Hmget(key string, fields ...interface{}) ([]interface{}, error) {
	conn := Pool().GetConn()
	defer conn.Close()

	args := append([]interface{}{prekey + key}, fields...)
	ret, err := conn.Do(op_HMGET, args...)
	if err != nil {
		logger.Error("Hmget:", err)
		return nil, err
	}

	return ret.([]interface{}), nil
}

/*
	hash: hmset
	pairs: field, value pair
*/
func Hmset(key string, pairs ...interface{}) error {
	conn := Pool().GetConn()
	defer conn.Close()

	args := append([]interface{}{prekey + key}, pairs...)
	_, err := conn.Do(op_HMSET, args...)
	if err != nil {
		logger.Error("Hmset:", err)
		return err
	}

	_, err = conn.Do(op_PEXPIRE, prekey+key, Default_Expire_Milliseconds)
	if err != nil {
		logger.Error("Hmset:expire:", err)
	}

	return nil
}

/*
	hash: hmset2
	pairs: field, value pair
*/
func Hmset2(key string, pair interface{}) error {
	conn := Pool().GetConn()
	defer conn.Close()

	_, err := conn.Do(op_HMSET, redis.Args{}.Add(prekey+key).AddFlat(pair)...)
	if err != nil {
		logger.Error("Hmset:", err)
		return err
	}

	_, err = conn.Do(op_PEXPIRE, prekey+key, Default_Expire_Milliseconds)
	if err != nil {
		logger.Error("Hmset:expire:", err)
	}

	return nil
}

/*
	hash
	objPtr: object pointer
*/
func Hgetall(key string, objPtr interface{}) error {
	conn := Pool().GetConn()
	defer conn.Close()

	// if not found, err still equals nil
	ret, err := redis.Values(conn.Do(op_HGETALL, prekey+key))
	if err != nil {
		logger.Error("Hgetall:", err)
		return err
	}

	err = redis.ScanStruct(ret, objPtr)
	if err != nil {
		logger.Error("Hgetall:", err)
		return err
	}

	return nil
}

// set expire
func PexpireSecond(key string, seconds int64) error {
	var errMsg string
	if key == "" {
		errMsg = "PexpireSecond: key can't be empty"
		logger.Error("PexpireSecond:", errMsg)
		return errors.New(errMsg)
	}

	if seconds <= 0 {
		errMsg = "PexpireSecond: seconds can't be set <= 0"
		logger.Error("PexpireSecond:", errMsg)
		return errors.New(errMsg)
	}

	conn := Pool().GetConn()
	defer conn.Close()

	_, err := conn.Do(op_PEXPIRE, prekey+key, 1000*seconds)
	if err != nil {
		logger.Error("PexpireSecond:", err)
		return err
	}

	return nil
}

//endregion

//region lists
func LPush(key string, value interface{}) (data int, err error) {
	conn := Pool().GetConn()
	defer conn.Close()

	ret, err := redis.Int(conn.Do(op_LPUSH, prekey+key, value))

	if err != nil {
		logger.Error("LPush:", err)
		return ret, err
	}
	return ret, nil
}

func LPushBatch(key string, values interface{}) (data int, err error) {
	conn := Pool().GetConn()
	defer conn.Close()
	args := redis.Args{}.Add(prekey + key).AddFlat(values)
	reply, err := redis.Int(conn.Do("LPUSH", args...))

	if err != nil {
		logger.Error("LPushBatch:", err)
		return
	}
	data = reply
	return
}

func LTrim(key string, start, end int64) (err error) {
	conn := Pool().GetConn()
	defer conn.Close()

	_, err = redis.String(conn.Do(op_LTRIM, prekey+key, start, end))
	if err != nil {
		logger.Error("LTrim:", err)
		return
	}
	return
}

func BRPop(key string, timeout int) ([]interface{}, error) {
	conn := Pool().GetConn()
	defer conn.Close()

	ret, err := redis.Values(conn.Do(op_BRPOP, prekey+key, timeout))

	if err != nil {
		logger.Error(op_BRPOP, err)
		return ret, err
	}

	return ret, nil
}
func BRPopBatchKey(keys []string, timeout int) ([]interface{}, error) {
	conn := Pool().GetConn()
	defer conn.Close()

	for i, v := range keys {
		keys[i] = prekey + v
	}

	ret, err := redis.Values(conn.Do(op_BRPOP, redis.Args{}.Add().AddFlat(keys).Add(timeout)...))

	if err != nil {
		logger.Error(op_BRPOP, err)
		return ret, err
	}

	return ret, nil
}

func BRPopLPush(key1, key2 string, timeout int) (interface{}, error) {
	conn := Pool().GetConn()
	defer conn.Close()

	ret, err := conn.Do(op_BRPOPLPUSH, prekey+key1, prekey+key2, timeout)

	if err != nil {
		//logger.Error(op_BRPOPLPUSH, err)
		return ret, err
	}

	return ret, nil
}

func LREM(key string, count int, value interface{}) (data int, err error) {
	conn := Pool().GetConn()
	defer conn.Close()
	reply, err := redis.Int(conn.Do(op_LREM, prekey+key, count, value))

	if err != nil {
		logger.Error(op_LREM, err)
		return
	}
	data = reply
	return
}

//endregion

func INCR(key string) (data int, err error) {
	conn := Pool().GetConn()
	defer conn.Close()
	reply, err := redis.Int(conn.Do(op_INCR, prekey+key))

	if err != nil {
		logger.Error(op_INCR, err)
		return
	}
	data = reply
	return
}
func INCRBY(key string, increment int) (data int, err error) {
	conn := Pool().GetConn()
	defer conn.Close()
	reply, err := redis.Int(conn.Do(op_INCRBY, prekey+key, increment))

	if err != nil {
		logger.Error(op_INCR, err)
		return
	}
	data = reply
	return
}

func DECR(key string) (data int, err error) {
	conn := Pool().GetConn()
	defer conn.Close()
	reply, err := redis.Int(conn.Do(op_DECR, prekey+key))

	if err != nil {
		logger.Error(op_DECR, err)
		return
	}
	data = reply
	return
}

func DECRBY(key string, val int) (data int, err error) {
	conn := Pool().GetConn()
	defer conn.Close()
	reply, err := redis.Int(conn.Do(op_DECRBY, prekey+key, val))

	if err != nil {
		logger.Error(op_DECRBY, err)
		return
	}
	data = reply
	return
}

func SetExpireAt(key string, expireAt int64) error {
	conn := Pool().GetConn()
	defer conn.Close()

	_, err := conn.Do(op_EXPIREAT, prekey+key, expireAt)

	return err
}

func SetExpire(key string, expire int64) error {
	conn := Pool().GetConn()
	defer conn.Close()

	_, err := conn.Do(op_EXPIRE, prekey+key, expire)
	return err
}

func SetPersist(key string) error {
	conn := Pool().GetConn()
	defer conn.Close()

	_, err := conn.Do(op_PERSIST, prekey+key)
	return err
}

func GetStr(key string) (data string, err error) {
	conn := Pool().GetConn()
	defer conn.Close()
	reply, err := redis.String(conn.Do(op_GET, prekey+key))

	if err != nil {
		logger.Error(op_GET, err)
		return
	}
	data = reply
	return
}

func SetStr(key string, value string, expire int64) (err error) {
	conn := Pool().GetConn()
	defer conn.Close()

	if 0 == expire {
		_, err = conn.Do(op_SET, prekey+key, value)
	} else {
		_, err = conn.Do(op_SET, prekey+key, value, "EX", expire)
	}

	if err != nil {
		logger.Error(op_SET, err)
		return
	}

	return
}
func HGetAllMap(key string) (data map[string]string, err error) {
	conn := Pool().GetConn()
	defer conn.Close()

	reply, err := redis.StringMap(conn.Do(op_HGETALL, prekey+key))
	if err != nil {
		logger.Error("HGetall map:", key, err)
		return
	}
	data = reply
	return
}

func HGet(key string, field interface{}) (data string, err error) {
	conn := Pool().GetConn()
	defer conn.Close()

	data, err = redis.String(conn.Do(op_HGET, prekey+key, field))
	if err != nil {
		logger.Error("Hget:", err)
		return
	}

	return
}

func HDel(key string, field interface{}) (err error) {
	conn := Pool().GetConn()
	defer conn.Close()

	_, err = conn.Do(op_HDEl, prekey+key, field)
	return
}

func HLen(key string) (data int, err error) {
	conn := Pool().GetConn()
	defer conn.Close()

	return redis.Int(conn.Do(op_HDLEN, prekey+key))
}
