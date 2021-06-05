package dao

import (
    "goldtalkAPI/pkg/client"
    "time"
)

func (SessionInfo) TableName() string {
    return "session_info"
}

type SessionInfo struct {
    ID               int64     `gorm:"id"`
    RefNo1               int64     `gorm:"ref_no1"`
    SessionStartTime time.Time `gorm:"session_start_time"`
    SessionEndTime   time.Time `gorm:"session_end_time"`
    Topic            string    `gorm:"topic"`
}

func GetSessionInfoAtTimeRange(start string, end string) (data []*SessionInfo, err error) {
    if err := client.DB.Where(
        "session_start_time >= ? and session_end_time < ? and channel_id = 1 and valid = 1", start, end,
    ).Find(&data).Error; err != nil {
        return nil, err
    }
    return
}

func getSessionInfoByParams(params map[string]interface{}) (data []*SessionInfo, err error) {
    if err := client.DB.Where(params).Find(&data).Error; err != nil {
        return nil, err
    }
    return
}
