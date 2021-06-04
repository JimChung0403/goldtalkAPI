package rpc

import (
    "errors"
    "goldtalkAPI/pkg/client"
    "goldtalkAPI/pkg/rpcmodel"
    "goldtalkAPI/pkg/util"
)

func GetCustomerBaseInfoList(clientSnList []int64) (data map[int64]*rpcmodel.CustomerBaseInfo, err error) {
    data = make(map[int64]*rpcmodel.CustomerBaseInfo)
    apiData := &rpcmodel.CustomerBaseInfoList{}
      sss := `
    {
      "list": [
        {
          "customerId": 1399959067501428736,
          "clientSn": 510040334,
          "leadSn": 55956858,
          "createdAt": "2021-06-02 13:20:26",
          "email": "0958003068@goldtalkAPI.com",
          "website": "T",
          "brandId": 1,
          "countryArea": "886",
          "mobile": "0958003068",
          "fname": "",
          "lname": "",
          "cname": "微職人註冊會員",
          "englishName": ""
        },
        {
          "customerId": 879178118768361472,
          "clientSn": 259465,
          "leadSn": 534559,
          "createdAt": "2011-03-03 01:04:00",
          "email": "nicoletestno4@tutorabc.com",
          "website": "T",
          "brandId": 1,
          "countryArea": "886",
          "mobile": "0987488478",
          "fname": "Tstest",
          "lname": "Cheng",
          "cname": "監控專用",
          "englishName": "Tstest Cheng"
        }
      ],
      "totalCount": 2
    }
      `
      err = util.JsonUnmarshalFromString(sss, &apiData)
      for _, d := range apiData.List{
          data[d.Clientsn] = &d
      }

      return

    req := &rpcmodel.CustomerBaseInfoReq{
        ClientSnList: clientSnList,
    }

    resp, err := client.PassportAPI.R().SetBody(util.JsonString(req)).Post("/customerBaseInfo/list")
    apiResp := &rpcmodel.PossportResponse{}
    err = util.JsonUnmarshalFromString(resp.String(), apiResp)
    if err != nil {
        return
    }
    if !apiResp.Success {
        err = errors.New("Api error" + apiResp.Message)
        return
    }

    err = util.JsonUnmarshalFromString(util.JsonString(apiResp.Data), &apiData)
    if err != nil {
        return
    }
    for _, d := range apiData.List {
        data[d.Clientsn] = &d
    }
    return
}
