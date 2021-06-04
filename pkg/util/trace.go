package util

import (
    "context"
    "fmt"
    "github.com/gin-gonic/gin"
    "goldtalkAPI/pkg/thirdparty/go-trace"
)

func NewContextFromGin(c *gin.Context) (ctx context.Context) {
    traceInfo, exist := c.Get("_tr")
    if exist {
        tr := traceInfo.(map[string]string)
        ctx = trace.NewContext(c.Request.Context(), trace.Trace(tr))
        return
    }
    return c.Request.Context()
}

func FromGinContext(c *gin.Context) (s string) {
    traceInfo, exist := c.Get("_tr")
    if exist {
        tr := traceInfo.(map[string]string)
        ctx := trace.NewContext(c.Request.Context(), trace.Trace(tr))
        return  fmt.Sprintf("%v", trace.FromContext(ctx))
    }
    return "trace_id=unknow"
}

