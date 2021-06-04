package middleware

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "goldtalkAPI/pkg/app"
    "goldtalkAPI/pkg/thirdparty/go-log"
    "goldtalkAPI/pkg/util"
    "net/http"
)

func PanicHandle() gin.HandlerFunc {
    return func(c *gin.Context) {
    defer func() {
        err := recover()
        if nil != err {
            stackInfo := util.GetStackInfo()
            msg := fmt.Sprintf("```%s``` \n\n `%v`||%s", stackInfo, err, util.FromGinContext(c))
            log.Error(msg)
            appG := app.Gin{C: c}
            appG.Response(http.StatusInternalServerError, app.ERROR, err)
        }
    }()
        c.Next()
    }
}

