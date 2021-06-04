package rpc

import (
    "context"
    "errors"
    "goldtalkAPI/conf"
    "goldtalkAPI/pkg/client"
    "goldtalkAPI/pkg/rpcmodel"
    "goldtalkAPI/pkg/thirdparty/go-log"
    "goldtalkAPI/pkg/thirdparty/go-trace"
    "goldtalkAPI/pkg/util"
)

const (
    Notify = "NOTIFY"
)

func SendSMSBatch(ctx context.Context, data []*rpcmodel.SendSmsBatch) (err error) {
    return
    req := &rpcmodel.SendSmsBatchReq{
        Sendsmsmessagereqlist: data,
        Appsn:                 conf.Conf.SMS.AppSn,
        Smstype:               Notify,
    }
    resp, err := client.SMSAPI.R().SetBody(util.JsonString(req)).Post("/batchAllTypeSendSMS")
    apiResp := &rpcmodel.SMSResponse{}
    err = util.JsonUnmarshalFromString(resp.String(), apiResp)
    if err != nil {
        return
    }
    log.Infof("%v||SendSMSBatch: %s", trace.ContextString(ctx), resp.String())
    if !apiResp.Success {
        err = errors.New("Api error" + apiResp.Message)
        return
    }
    return
}
