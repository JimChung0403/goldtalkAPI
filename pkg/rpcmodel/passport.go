package rpcmodel

type PossportResponse struct {
    Success bool        `json:"success"`
    Code    string      `json:"code"`
    Message string      `json:"message"`
    Data    interface{} `json:"data"`
    Trace   struct {
        Requestid        string `json:"requestId"`
        Requeststarttime string `json:"requestStartTime"`
        Requestendtime   string `json:"requestEndTime"`
        Provider         string `json:"provider"`
        Consumer         string `json:"consumer"`
    } `json:"trace"`
}

type CustomerBaseInfoReq struct {
    ClientSnList []int64 `json:"clientSnList"`
}

type CustomerBaseInfoList struct {
    List []CustomerBaseInfo `json:"list"`
    Totalcount int `json:"totalCount"`
}

type CustomerBaseInfo struct {
    Customerid  int64  `json:"customerId"`
    Clientsn    int64  `json:"clientSn"`
    Leadsn      int64  `json:"leadSn"`
    Createdat   string `json:"createdAt"`
    Email       string `json:"email"`
    Website     string `json:"website"`
    Brandid     int    `json:"brandId"`
    Countryarea string `json:"countryArea"`
    Mobile      string `json:"mobile"`
    Fname       string `json:"fname"`
    Lname       string `json:"lname"`
    Cname       string `json:"cname"`
    Englishname string `json:"englishName"`
    Isjr        bool   `json:"isJr"`
}
