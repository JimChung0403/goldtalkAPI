package rpcmodel

type SMSResponse struct {
    Success    bool        `json:"success"`
    Message    string      `json:"message"`
    Code       int         `json:"code"`
    Data       interface{} `json:"data"`
    Reminder   interface{} `json:"reminder"`
    Errorlevel int         `json:"errorLevel"`
    Trace      struct {
        Traceid        string `json:"traceId"`
        Tracestarttime string `json:"traceStartTime"`
        Traceendtime   string `json:"traceEndTime"`
        Provider       string `json:"provider"`
        Consumer       string `json:"consumer"`
        Remoteip       string `json:"remoteIp"`
    } `json:"trace"`
}

type SendSmsBatchReq struct {
    Sendsmsmessagereqlist []*SendSmsBatch `json:"sendSMSMessageReqList"`
    Appsn                 string         `json:"appSn"`
    Smstype               string         `json:"smsType"`
}

type SendSmsBatch struct {
    Phone         string `json:"phone"`
    Countrycode   string `json:"countryCode"`
    Message       string `json:"message"`
    Signaturecode string `json:"signatureCode"`
}
