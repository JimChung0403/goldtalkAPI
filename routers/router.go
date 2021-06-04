package routers

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "goldtalkAPI/conf"
    "goldtalkAPI/middleware"
    "goldtalkAPI/pkg/service"
    "os"
)

func setLogger(r *gin.Engine){
    if gin.Mode() == gin.ReleaseMode{
        file, fileErr := os.Create(conf.Conf.Server.GinFile)
        if fileErr != nil {
            fmt.Printf("fail to log file||err=%+v", fileErr)
            os.Exit(1)
        }
        r.Use(gin.LoggerWithWriter(file,""))
        return
    }

    r.Use(gin.Logger())
}


// InitRouter
func InitRouter() *gin.Engine {
    r := gin.New()
    setLogger(r)
    r.Use(middleware.SetTrace())
    r.Use(middleware.PanicHandle())
    apiv1 := r.Group("/api")
    apiv1.Use()
    {
        //取得教師列表
        apiv1.GET("/tutors/:language", service.GetTutorList)


    }

    return r
}
