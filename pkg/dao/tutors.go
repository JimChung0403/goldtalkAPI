package dao

import "goldtalkAPI/pkg/client"

type Tutors struct {
    ID           int64  `gorm:"id"`
    Slug         string `gorm:"slug"`
    Name         string `gorm:"name"`
    Headline     string `gorm:"headline"`
    Introduction string `gorm:"introduction"`
}

func GetTutorsMapByIDs(tutorIDs []int64) (data map[int64]*Tutors, err error) {
    data = make(map[int64]*Tutors)
    tutorList, err := GetTutorsByTutorIDs(tutorIDs)
    if err != nil {
        return
    }
    for _, t := range tutorList {
        data[t.ID] = t
    }
    return
}

func GetTutorsByLanguage(lang string) (data *Tutors, err error) {
    parmas := make(map[string]interface{})
    parmas["slug"] = lang

    tutorList, err := GetTutorsByParams(parmas)
    if err != nil{
        return
    }

    data = tutorList[0]
    return
}

func GetTutorsByParams(params map[string]interface{}) (data []*Tutors, err error) {
    if err := client.DB.Where(params).Find(&data).Error; err != nil {
        return nil, err
    }
    return
}

//gorm應該有支持in的用法, 暫時先這樣寫
func GetTutorsByTutorIDs(tutorIDs []int64) (data []*Tutors, err error) {
    if err := client.DB.Where("id IN (?)", tutorIDs).Find(&data).Error; err != nil {
        return nil, err
    }
    return
}
