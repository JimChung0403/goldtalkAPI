package app

import (
    "github.com/gin-gonic/gin"
    "goldtalkAPI/pkg/thirdparty/go-trace"
)

type Gin struct {
    C *gin.Context
}

type Response struct {
    Code  int         `json:"code"`
    Msg   string      `json:"msg"`
    Data  interface{} `json:"data"`
    Trace TraceResp   `json:"trace"`
}

type TraceResp struct {
    TraceID string `json:"trace_id"`
}

// Response setting gin.JSON
func (g *Gin) Response(httpCode, errCode int, data interface{}) {
    g.C.JSON(httpCode, Response{
        Code:  errCode,
        Msg:   GetMsg(errCode),
        Data:  data,
        Trace: traceResp(g.C),
    })
    return
}

func traceResp(c *gin.Context) (t TraceResp) {
    traceInfo, exist := c.Get("_tr")
    if exist {
        tr := traceInfo.(map[string]string)
        t.TraceID = trace.Trace(tr).Traceid().String()
        return
    }
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
