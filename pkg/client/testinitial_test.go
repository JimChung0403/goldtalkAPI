package client

import (
	"context"
	"fmt"
	"goldtalkAPI/conf"
	"goldtalkAPI/pkg/util"
	"testing"
	"time"
)

func Test_Initial(t *testing.T) {
	t.Run("Test_Initial", func(t *testing.T) {
		conf.LoadConfigFile("../../conf/service.conf")
		InitClients(conf.Conf)

		//fmt.Println("###")
		//fmt.Println(redis.SetNXExpired("lock_key", "1", time.Second * 10))
		//fmt.Println("###")
		//fmt.Println(redis.Set("aaaa", "111", time.Second))
		//fmt.Println(redis.GetString("aaaa"))

		dLock := util.NewDLock("aaa", time.Second*100)

		ctx := context.Background()
		l , err := dLock.SetLock(ctx)
		fmt.Println("dLock:", l, err)
		fmt.Println(dLock.UnLock(ctx))

		fmt.Println(dLock.SetLock(ctx))

	})
}