package middleware

import (
    "github.com/gin-gonic/gin"
    "goldtalkAPI/pkg/thirdparty/go-trace"
    "time"
)

func SetTrace() gin.HandlerFunc {
    return func(c *gin.Context) {
        traceInfo := make(map[string]string)
        traceID := trace.MakeTraceid(time.Now().UnixNano())
        traceInfo["traceid"] = traceID.String()
        //fmt.Println("setTrace: ", traceID.String())
        c.Set("_tr", traceInfo)
    }
}
