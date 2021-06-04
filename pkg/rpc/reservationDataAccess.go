package rpc

import (
    "errors"
    "goldtalkAPI/pkg/client"
    "goldtalkAPI/pkg/rpcmodel"
    "goldtalkAPI/pkg/util"
)



//func Get() error {
//    fmt.Println(client.TTAPI)
//    fmt.Println("=====")
//    resp, _ := client.TTAPI.R().Execute(resty.MethodGet, "/get")
//    fmt.Println(resp)
//    fmt.Println("=====")
//    return nil
//}

func GetClassInformationByLobbySn(lobbySn int64) (data []*rpcmodel.ClassInformationByLobbySn, err error) {
    sss := `
[
{
  "RecordFrom": 4,
  "RecordSn": "202106022030_510012483_4060871",
  "BrandId": 1,
  "ClientSn": 510012483,
  "IsJunior": false,
  "StartDateTime": "2021-06-02 20:30:00",
  "SstNumber": 99,
  "ServiceId": 1,
  "ContractSn": 1222116,
  "Level": 0,
  "LobbySn": 232020,
  "ConsultantSn": 4327,
  "MaterialSn": 134930,
  "SessionSn": "2021060220430721",
  "CostPoints": 1,
  "CostFrom": 1,
  "VersionCode": "ff53a1ac-e48e-47c1-95a6-843945ec9712",
  "CreatorSourceToken": "2101121",
  "ModifierSourceToken": "2007091",
  "IsCancel": false,
  "IsDelete": false,
  "CreatedAt": "2021-06-02 17:17:21",
  "ModifiedAt": "2021-06-02 20:05:02",
  "StrategyIds": [
    "goldtalkAPI"
  ]
}
]
    `
    //sss := `[{"RecordFrom":4,"RecordSn":"202106022030_510012483_4060871","BrandId":1}]`
    err = util.JsonUnmarshalFromString(sss, &data)
    return

    req := &rpcmodel.ClassInfoByLobbySnReq{
        Data: &rpcmodel.ClassInfoByLobbySnData{
            Lobbysns:         []int64{lobbySn},
            HaveStrategyid:   true,
            IsonlyStrategyid: true,
        },
    }

    resp, err := client.RDAAPI.R().SetBody(util.JsonString(req)).Post("/Class/GetClassInformationByLobbySn")
    apiResp := &rpcmodel.ReservationDataAccessResponse{}
    err = util.JsonUnmarshalFromString(resp.String(), apiResp)
    if err != nil{
        return
    }
    if !apiResp.Success{
        err = errors.New("Api error" + apiResp.Garbage)
        return
    }

    err = util.JsonUnmarshalFromString(util.JsonString(apiResp.Data), &data)
    if err != nil{
        return
    }
    return
}
