package util

import "time"

const (
    TaipeiTimeZone = "Asia/Taipei"
)

func NowDateTimeStr() string {
    loc, _ := time.LoadLocation(TaipeiTimeZone)
    return time.Now().In(loc).Format("2006-01-02 15:04:05")
}

func NowDateTime() time.Time {
    loc, _ := time.LoadLocation(TaipeiTimeZone)
    return time.Now().In(loc)
}

func TimeMin2Str(t time.Time) string {
    loc, _ := time.LoadLocation(TaipeiTimeZone)
    return t.In(loc).Format("2006-01-02 15:04")
}

