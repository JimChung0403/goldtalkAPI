package model

type PriceInfo struct {
    Trial  float32 `json:"trial"`
    Normal float32 `json:"normal"`
}

type TutorData struct {
    ID                int64      `json:"id"`
    Slug              string     `json:"slug"`
    Name              string     `json:"name"`
    Headline          string     `json:"headline"`
    Introduction      string     `json:"introduction"`
    PriceInfo         *PriceInfo `json:"price_info"`
    TeachingLanguages []int64    `json:"teaching_languages"`
}

