package main

import (
    "flag"
    "fmt"
    "github.com/BurntSushi/toml"
    "github.com/gin-gonic/gin"
    "goldtalkAPI/conf"
    "goldtalkAPI/pkg/client"
    "goldtalkAPI/pkg/schedule"
    "goldtalkAPI/pkg/thirdparty/go-log"
    "goldtalkAPI/routers"
    "net/http"
    "os"
    "runtime"
    "time"
)



func initHttpSvr(config conf.Config) *http.Server{

    //init gin framwork
    gin.SetMode(config.Server.RunMode)
    //InitRouter
    routersInit := routers.InitRouter()
    readTimeout := config.Server.ReadTimeout
    writeTimeout := config.Server.WriteTimeout
    endPoint := fmt.Sprintf(":%d", config.Server.HttpPort)
    maxHeaderBytes := 1 << 20

    server := &http.Server{
        Addr:           endPoint,
        Handler:        routersInit,
        ReadTimeout:    readTimeout.Duration,
        WriteTimeout:   writeTimeout.Duration,
        MaxHeaderBytes: maxHeaderBytes,
    }

    fmt.Printf("start http server listening %s\n", endPoint)
    return server
}

func main() {
    // 命令行参数。
    var configPath string
    flag.StringVar(&configPath, "config", "conf/service.conf", "server config.")
    flag.Parse()

    // 解析配置。
    if _, err := toml.DecodeFile(configPath, &conf.Conf); err != nil {
        fmt.Printf("fail to read config.||err=%+v||config=%+v", err, configPath)
        os.Exit(1)
        return
    }
    config := conf.Conf
    log.Init(&config.Log)
    defer log.Close()

    client.InitRDAAPI("http://tutorgroupapi.tutorabc.com/ReservationDataAccess")
    client.InitPassportAPI("http://apitw.passport.tutorabc.com/web")
    client.InitSMS("http://sms.tutorabc.com/twapi")

    client.Setup(config.DB)
    defer client.CloseDB()
    //統計NumGoroutine
    go monitorInfo()
    go schedule.SchSendSmsAtSessionTime()



    svr := initHttpSvr(config)
    svr.ListenAndServe()

}


func monitorInfo() {
    var timer *time.Ticker = time.NewTicker(10 * time.Second)
    defer timer.Stop()
    for {
        select {
        case <-timer.C:
            log.Infof("mainNumGoroutine=%d", runtime.NumGoroutine())
        }
    }
    return
}
