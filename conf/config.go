package conf

import (
    "fmt"
    "github.com/BurntSushi/toml"
    "goldtalkAPI/pkg/thirdparty/go-cache/redis_sentinel"
    "goldtalkAPI/pkg/thirdparty/go-log"
    "path/filepath"
    "time"
)

var (
    Conf Config
)

// Config 表示一个 thrift 服务器的配置。
type ServerConfig struct {
    RunMode      string   `toml:"runMode"`
    HttpPort     int      `toml:"httpPort"`
    ReadTimeout  time.Duration `toml:"readTimeout"`
    WriteTimeout time.Duration `toml:"writeTimeout"`
    GinFile      string   `toml:"gin_file"`
}

type Database struct {
    Type        string
    User        string
    Password    string
    Host        string
    Name        string
    TablePrefix string
    IdleConns   int
    OpenConns   int
}

type SMS struct {
    AppSn string `toml:"appsn"`
}

type Cachecloud struct {
    AppID int64 `toml:"appid"`
}

type APIHost struct {
    RDA string `toml:"rda"`
    Passport string `toml:"passport"`
    SMS string `toml:"sms"`
}

// 对应 conf/service.conf 的结构。
type Config struct {
    Server     ServerConfig `toml:"server"`
    Log        log.Config   `toml:"log"`
    APIHost    APIHost      `toml:"apihost"`
    DB         Database     `toml:"database"`
    SMS        SMS          `toml:"sms"`
    Cachecloud Cachecloud   `toml:"cachecloud"`
    Redis      redis.Config `toml:"redis"`
    //RedisConf            redis.Config         `toml:"redis"`
    //DApi                 client.Config        `toml:"d_api"`
}

func LoadConfigFile(confFilePath string) (err error) {
    if confFilePath == "" {
        confFilePath = "../conf/service.toml"
    }

    absConfFilePath, _ := filepath.Abs(confFilePath)
    fmt.Println("load config from " + absConfFilePath)
    _, err = toml.DecodeFile(confFilePath, &Conf)
    return err
}
