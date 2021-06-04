package dao

import "goldtalkAPI/pkg/client"

func (TutorLanguages) TableName() string {
    return "tutor_languages"
}

type TutorLanguages struct {
    ID         int64 `gorm:"id"`
    TutorID    int64 `gorm:"tutor_id"`
    LanguageID int64 `gorm:"language_id"`
}

func GetTutorIDByLanguageID(languageID int64) (data []int64, err error) {
    parmas := make(map[string]interface{})
    parmas["language_id"] = languageID

    tutorLangList, err := GetTutorLanguagesByParams(parmas)
    if err != nil {
        return
    }

    for _, d := range tutorLangList {
        data = append(data, d.TutorID)
    }
    return
}

func GetLangIDMapByTutorIDs(tutorIDs []int64) (data map[int64][]int64, err error) {
    data = make(map[int64][]int64)
    tutorLangList, err := GetTutorLanguagesByTutorIDs(tutorIDs)
    if err != nil {
        return
    }

    for _, tutorLang := range tutorLangList {
        if _, ok := data[tutorLang.TutorID]; !ok {
            data[tutorLang.TutorID] = []int64{}
        }
        data[tutorLang.TutorID] = append(data[tutorLang.TutorID], tutorLang.LanguageID)
    }

    return
}

func GetTutorLanguagesByParams(params map[string]interface{}) (data []*TutorLanguages, err error) {
    if err := client.DB.Where(params).Find(&data).Error; err != nil {
        return nil, err
    }
    return
}

//gorm應該有支持in的用法, 暫時先這樣寫
func GetTutorLanguagesByTutorIDs(tutorIDs []int64) (data []*TutorLanguages, err error) {
    if err := client.DB.Where("tutor_id IN (?)", tutorIDs).Find(&data).Error; err != nil {
       return nil, err
    }
    return
}
