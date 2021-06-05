package schedule

import (
    "context"
    "fmt"
    "goldtalkAPI/pkg/dao"
    "goldtalkAPI/pkg/rpc"
    "goldtalkAPI/pkg/rpcmodel"
    "goldtalkAPI/pkg/thirdparty/go-log"
    "goldtalkAPI/pkg/thirdparty/go-trace"
    "goldtalkAPI/pkg/util"
    "time"
)

const (
    querySize  = 100
    contentFmt = "%s %s，快到上課時間囉！提醒您，開課前3分鐘，點擊以下連結進入教室。%s"
    courseUrl  = "https://www.gogoldtalk.com/course/%s"
    signature  = "TutorABC"
)

func SchSendSmsAtSessionTime() {

    var timer *time.Ticker = time.NewTicker(time.Minute * 10)
    defer timer.Stop()
    for {
        select {
        case <-timer.C:
            ctx := trace.NewContext(context.Background(), nil)
            log.Infof("SchSendSmsAtSessionTime||%v||start at %s", trace.FromContext(ctx), util.NowDateTimeStr())
            now := util.NowDateTime()
            dataList, err := dao.GetSessionInfoAtStartTime(util.TimeMin2Str(now.Add(1 * time.Hour)))
            if err != nil {
                log.Errorf("_com_sch_error||%v||SchSendSmsAtSessionTime error||err=%v", trace.FromContext(ctx), err)
                break
            }
            for i, _ := range dataList {
                go SendSmsBySessionInfo(ctx, dataList[i])
            }


        }
    }
    return
}

func SendSmsBySessionInfo(ctx context.Context, info *dao.SessionInfo) {

    smsConent := fmt.Sprint(contentFmt,
        util.TimeMin2Str(info.SessionStartTime),
        info.Topic,
        fmt.Sprint(courseUrl, info.ID),
    )
    classInfoList, err := rpc.GetClassInformationByLobbySn(info.RefNo1)
    if err != nil {
        log.Errorf("_com_sch_error||%v||GetClassInformationByLobbySn error||err=%v", trace.FromContext(ctx), err)
        return
    }

    log.Infof("%v||SendSmsBySessionInfo: %s||bch: %s", trace.ContextString(ctx), util.JsonString(info), util.JsonString(classInfoList))

    return
    cList := []int64{}
    cMap := make(map[int64]struct{})
    for _, info := range classInfoList {
        cList = append(cList, info.Clientsn)
        cMap[info.Clientsn] = struct{}{}
    }
    count := len(cList)

    for no := 0; no <= (count / querySize); no++ {
        pagingIDs := paginate(cList, no*querySize, querySize)
        if len(pagingIDs) == 0 {
            continue
        }
        cusBaseInfoMap, err := rpc.GetCustomerBaseInfoList(pagingIDs)
        if err != nil {
            log.Errorf("_com_sch_error||%v||GetCustomerBaseInfoList error||err=%v", trace.FromContext(ctx), err)
            return
        }

        smsReq := []*rpcmodel.SendSmsBatch{}
        for sn, cInfo := range cusBaseInfoMap {
            if _, ok := cMap[sn]; ok {
                if util.InStringSlice(cInfo.Mobile, []string{"0930049641", "0955011176", "0975498244", "0911835036"}) == false {
                    continue
                }
                bch := &rpcmodel.SendSmsBatch{
                    Phone:         cInfo.Mobile,
                    Countrycode:   cInfo.Countryarea,
                    Message:       smsConent,
                    Signaturecode: signature,
                }
                smsReq = append(smsReq, bch)
                log.Infof("%v||cusBaseInfoMap: %s||bch: %v", trace.ContextString(ctx), err, util.JsonString(bch))
            }
        }

        err = rpc.SendSMSBatch(ctx, smsReq)
        log.Infof("%v||SendSMSBatch: %s||session: %v", trace.ContextString(ctx), err, util.JsonString(info))

    }
    return
}

func paginate(x []int64, skip int, size int) []int64 {
    if skip > len(x) {
        skip = len(x)
    }

    end := skip + size
    if end > len(x) {
        end = len(x)
    }

    return x[skip:end]
}
