package dao

import (
    "fmt"
    "goldtalkAPI/pkg/util"
    "testing"
)



func TestDao_GetSessionInfoByParams(t *testing.T) {
    t.Run("GetSessionInfoByParams", func(t *testing.T) {

        parmas := make(map[string]interface{})
        parmas["id"] = 16126

        dataList, err := getSessionInfoByParams(parmas)
        fmt.Println(err)
        fmt.Println(util.JsonString(dataList[0].SessionEndTime))

        //assert.Equal(t, err, nil)
        //assert.Equal(t, len(dataMap), 2)
    })
}

func TestDao_GetSessionInfoBySessionTime(t *testing.T) {
    t.Run("GetSessionInfoBySessionTime", func(t *testing.T) {

        dataList, err := GetSessionInfoAtStartTime("2021-06-03 20:30")
        fmt.Println(err)
        for _, d := range dataList{
            fmt.Println(util.JsonString(d))
        }

        //assert.Equal(t, err, nil)
        //assert.Equal(t, len(dataMap), 2)
    })
}








