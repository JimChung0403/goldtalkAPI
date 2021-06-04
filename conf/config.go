package conf

import (
    "fmt"
    "github.com/BurntSushi/toml"
    "goldtalkAPI/pkg/thirdparty/go-log"
    "path/filepath"
    "time"
)

var (
    Conf Config
)

type duration struct {
    time.Duration
}

func (d *duration) UnmarshalText(text []byte) error {
    var err error
    d.Duration, err = time.ParseDuration(string(text))
    return err
}

// Config 表示一个 thrift 服务器的配置。
type ServerConfig struct {
    RunMode      string   `toml:"runMode"`
    HttpPort     int      `toml:"httpPort"`
    ReadTimeout  duration `toml:"readTimeout"`
    WriteTimeout duration `toml:"writeTimeout"`
    GinFile string `toml:"gin_file"`
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

// 对应 conf/service.conf 的结构。
type Config struct {
    Server ServerConfig `toml:"server"`
    Log    log.Config   `toml:"log"`

    DB Database `toml:"database"`
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
