package rpcmodel

type ReservationDataAccessResponse struct {
    Success bool        `json:"Success"`
    Data    interface{} `json:"Data"`
    Garbage string      `json:"Garbage"`
}

type ClassInformationByLobbySn struct {
    Recordfrom          int64      `json:"RecordFrom"`
    Recordsn            string   `json:"RecordSn"`
    Brandid             int64      `json:"BrandId"`
    Clientsn            int64      `json:"ClientSn"`
    Isjunior            bool     `json:"IsJunior"`
    Startdatetime       string   `json:"StartDateTime"`
    Sstnumber           int64      `json:"SstNumber"`
    Serviceid           int64      `json:"ServiceId"`
    Contractsn          int64      `json:"ContractSn"`
    Level               int64      `json:"Level"`
    Lobbysn             int64      `json:"LobbySn"`
    Consultantsn        int64      `json:"ConsultantSn"`
    Materialsn          int64      `json:"MaterialSn"`
    Sessionsn           string   `json:"SessionSn"`
    Costpoints          int64      `json:"CostPoints"`
    Costfrom            int64      `json:"CostFrom"`
    Versioncode         string   `json:"VersionCode"`
    Creatorsourcetoken  string   `json:"CreatorSourceToken"`
    Modifiersourcetoken string   `json:"ModifierSourceToken"`
    Iscancel            bool     `json:"IsCancel"`
    Isdelete            bool     `json:"IsDelete"`
    Createdat           string   `json:"CreatedAt"`
    Modifiedat          string   `json:"ModifiedAt"`
    Strategyids         []string `json:"StrategyIds"`
}

type ClassInfoByLobbySnReq struct {
    Data *ClassInfoByLobbySnData `json:"data"`
}

type ClassInfoByLobbySnData struct {
    Lobbysns         []int64 `json:"lobbySns"`
    IsonlyStrategyid bool    `json:"IsOnlyStrategyId"`
    HaveStrategyid   bool    `json:"haveStrategyId"`
}
