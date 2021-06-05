package redis

import "time"

const (
    DefaultMaxIdle     = 10
    DefaultMasterName  = "mymaster" //一定要mymaster why??
    DefaultIdleTimeout = 5 * time.Minute
    DefaultConnTimeout = 50 * time.Millisecond
    DefaultSendTimeout = 500 * time.Millisecond
    DefaultReadTimeout = 3 * time.Second

    maxPoolSize = 128
    minPoolSize = 8

)

type Config struct {
    Addrs       []string      `toml:"addrs"`
    PoolSize    int           `toml:"pool_size"`
    Prefix      string        `toml:"prefix"`
    MasterName  string        `toml:"master_name"`
    MaxIdle     int           `toml:"max_idle"`     // 连接超时时间，默认是 DefaultConnTimeout。
    IdleTimeout time.Duration `toml:"idle_timeout"` // 如果连接空闲，自动断开的时间，默认是 DefaultIdleTimeout。
    ConnTimeout time.Duration `toml:"conn_timeout"` // 连接超时时间，默认是 DefaultConnTimeout。
    SendTimeout time.Duration `toml:"send_timeout"` // 发送数据超时时间，默认是 DefaultSendTimeout。
    ReadTimeout time.Duration `toml:"read_timeout"` // 读取数据超时时间，默认是 DefaultReadTimeout。
}
