package service

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "goldtalkAPI/pkg/app"
    "goldtalkAPI/pkg/model"
    "goldtalkAPI/pkg/thirdparty/go-trace"
    "goldtalkAPI/pkg/util"
    "net/http"
    "time"
)


type TutorListByLang struct {
    LastUpdate time.Time
    Data       []*model.TutorData
}

// @Summary 測試用
// @Id 1
// @Tags Hello
// @version 1.0
// @produce application/json
// @Success 200 string string 成功後返回的值
// @Router /api/tutors/:language [get]
func GetTutorList(c *gin.Context) {
    ctx := util.NewContextFromGin(c)
    appG := app.Gin{C: c}
    //lb := util.NewLogBuild(c, "GetTutorList")
    //defer func() {
    //    log.Info(lb)
    //}()
    //
    fmt.Println(trace.FromContext(ctx))
    ////驗證參數
    //language := c.Param("language")
    //if language == "" {
    //    appG.Response(http.StatusOK, app.INVALID_PARAMS, nil)
    //    return
    //}
    //language = strings.ToLower(language)
    ////todo: 驗證language是否有效
    //
    ////第一個request, 完全沒cache, 重新拿
    //data, err := model.GenTutorListByLang(language)
    //
    //if err != nil && err != dao.ErrNotFoundData {
    //    appG.Response(http.StatusOK, app.ERROR, nil)
    //    return
    //}
    //
    appG.Response(http.StatusOK, app.SUCCESS, 1)
}
