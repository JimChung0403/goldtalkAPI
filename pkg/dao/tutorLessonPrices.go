package dao

import "goldtalkAPI/pkg/client"

func (TutorLessonPrices) TableName() string {
    return "tutor_lesson_prices"
}

type TutorLessonPrices struct {
    ID          int64   `gorm:"id"`
    TutorID     int64   `gorm:"tutor_id"`
    TrialPrice  float32 `gorm:"trial_price"`
    NormalPrice float32 `gorm:"normal_price"`
}

func GetLessonPricesMapByTutorIDs(tutorIDs []int64) (data map[int64]*TutorLessonPrices, err error) {
    data = make(map[int64]*TutorLessonPrices)
    tutorPriceList, err := GetTutorLessonPricesByTutorIDs(tutorIDs)
    if err != nil{
        return
    }
    for _, tp := range tutorPriceList{
        data[tp.ID] = tp
    }
    return
}

func GetTutorLessonPricesByParams(params map[string]interface{}) (data []*TutorLessonPrices, err error) {
    if err := client.DB.Where(params).Find(&data).Error; err != nil {
        return nil, err
    }
    return
}

//gorm應該有支持in的用法, 暫時先這樣寫
func GetTutorLessonPricesByTutorIDs(tutorIDs []int64) (data []*TutorLessonPrices, err error) {
    if err := client.DB.Where("tutor_id IN (?)", tutorIDs).Find(&data).Error; err != nil {
        return nil, err
    }
    return
}
