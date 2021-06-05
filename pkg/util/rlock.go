package util

import (
   "context"
   "goldtalkAPI/pkg/thirdparty/go-log"
   "time"
   "goldtalkAPI/pkg/thirdparty/go-cache/redis_sentinel"
)

const (
   lockKeyPrefix = "lock_"
)
//todo 定时任务实例间获取锁，防止多个实例重复跑任务
//DLock for delivery lock
type DLock struct {
   key    string
   val    string
   expire time.Duration
}

// NewDLock used for new DLock instance
func NewDLock(key string, expire time.Duration) *DLock {
   return &DLock{
       key:    lockKeyPrefix + key,
       val:    "1",
       expire: expire,
   }
}

func (d *DLock) MyKey() (string) {
    return d.key
}


// Lock lock the key
func (d *DLock) SetLock(ctx context.Context) (locked bool, err error) {
   lb := NewLogBuild(ctx, "SetLock")

   defer func() {
       if lb.IsError() {
           log.Error(lb)
           return
       }

       log.Info(lb)
   }()

    res, err := redis.SetNXExpired(d.key, d.val, d.expire)
   if err != nil {
       lb.SetError(err).SetDescp("redis_error_happens")
       return
   }
   if res == false {
       lb.SetDescp("key_already_exists")
       return
   }

   lb.SetDescp("lock_ok")
   return true, nil
}


// UnLock unlock the key
func (d *DLock) UnLock(ctx context.Context) (err error) {
   lb := NewLogBuild(ctx, "UnLock")

   defer func() {
       if lb.IsError() {
           log.Error(lb)
           return
       }
       log.Info(lb)
   }()

   _, err = redis.Delete(d.key)
   if err != nil {
       lb.SetError(err).SetDescp("unlock_failed")

       return
   }
   lb.SetDescp("unlock_succ")
   return
}
