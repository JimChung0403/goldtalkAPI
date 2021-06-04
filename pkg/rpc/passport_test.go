package rpc

import (
    "fmt"
    "goldtalkAPI/pkg/util"
    "testing"
)



func TestService_GetCustomerBaseInfoList(t *testing.T) {
    t.Run("A", func(t *testing.T) {
        r , e := GetCustomerBaseInfoList([]int64{1,2,3})
        fmt.Println(e)
        fmt.Println("###")
        fmt.Println(util.JsonString(r))
        fmt.Println("###")
    })
}






