package schedule

import (
    "context"
    "fmt"
    "goldtalkAPI/pkg/dao"
    "goldtalkAPI/pkg/rpc"
    "goldtalkAPI/pkg/thirdparty/go-log"
    "goldtalkAPI/pkg/thirdparty/go-trace"
    "goldtalkAPI/pkg/util"
    "time"
)

const (
    querySize = 100
)

func SchSendSmsAtSessionTime() {

    var timer *time.Ticker = time.NewTicker(time.Second * 10)
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
            for i, d := range dataList {
                fmt.Sprint(d)
                go SendSmsBySessionInfo(ctx, dataList[i])
            }


        }
    }
    return
}

type SMSContent struct {
    Msg   string
    Phone string
}

func SendSmsBySessionInfo(ctx context.Context, info *dao.SessionInfo) {

    classInfoList, err := rpc.GetClassInformationByLobbySn(info.RefNo1)
    if err != nil {
        log.Errorf("_com_sch_error||%v||GetClassInformationByLobbySn error||err=%v", trace.FromContext(ctx), err)
        return
    }

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

        aa := []*SMSContent{}
        for clientSn, Info := range cusBaseInfoMap {
            if _, ok := cMap[clientSn]; ok {
                aa = append(aa, &SMSContent{
                    Msg:   "",
                    Phone: Info.Mobile,
                })
            }
        }

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
