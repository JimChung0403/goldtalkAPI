package app

import (
    "github.com/gin-gonic/gin"
)

type Gin struct {
    C *gin.Context
}

type Response struct {
    Code int         `json:"code"`
    Msg  string      `json:"msg"`
    Data interface{} `json:"data"`
    Trace string `json:"_trace"`
}

// Response setting gin.JSON
func (g *Gin) Response(httpCode, errCode int, data interface{}) {
    g.C.JSON(httpCode, Response{
        Code: errCode,
        Msg:  GetMsg(errCode),
        Data: data,
        //Trace: trace.FromGinTrace(g.C),
    })
    return
}

const (
    SUCCESS        = 200
    ERROR          = 500
    INVALID_PARAMS = 400
)

var MsgFlags = map[int]string{
    SUCCESS:        "ok",
    ERROR:          "server error",
    INVALID_PARAMS: "param invalid",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
    msg, ok := MsgFlags[code]
    if ok {
        return msg
    }

    return MsgFlags[ERROR]
}
