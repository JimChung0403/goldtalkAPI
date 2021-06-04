package service

import (
    "fmt"
    "goldtalkAPI/pkg/model"
    "sync"
    "time"
)

var (
    TutorCacheMap  sync.Map //用sync.map解決race conditions
    UpdTutorChan   = make(chan string)
    TutorCacheTime = time.Duration(10 * time.Minute)
    TutorBySlugData = &TutorBySlugD{CacheTime: time.Duration(time.Minute * 10)}
)


type TutorBySlug struct {
    LastUpdate time.Time
    Data       *model.TutorData
}


type SyncType struct {
    UpdType int64
}


type TutorBySlugD struct {
    CacheTime time.Duration
}

func(c *TutorBySlug) GetData() (r *model.TutorData){


    return
}
//
//
//func GetTutorBySlug(c *gin.Context) {
//    appG := app.Gin{C: c}
//    slug := c.Param("slug")
//
//    //if value, ok := TutorCacheMap.Load(slug); ok {
//    //    tutor := value.(*TutorBySlug)
//    //    fmt.Println(time.Since(tutor.LastUpdate))
//    //    if time.Since(tutor.LastUpdate) >= TutorCacheTime {
//    //        putUpdTutorQueue(slug)
//    //    }
//    //
//    //    if tutor.Data != nil {
//    //        appG.Response(http.StatusOK, app.SUCCESS, tutor.Data)
//    //        return
//    //    }
//    //}
//    fmt.Println("No Cache: ", slug)
//    putUpdTutorQueue(slug)
//
//    //第一個request, 完全沒cache
//    data, err := model.GenTutorBySlug(slug)
//    if err != nil && err != dao.ErrNotFoundData {
//        appG.Response(http.StatusOK, app.ERROR, nil)
//        return
//    }
//
//    appG.Response(http.StatusOK, app.SUCCESS, data)
//}

func putUpdTutorQueue(slug string) {
    go func(slug string) {
        fmt.Println("putQueue", slug)
        UpdTutorChan <- slug
    }(slug)
}
