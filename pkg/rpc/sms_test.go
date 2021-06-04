package rpc

import (
    "context"
    "fmt"
    "goldtalkAPI/pkg/rpcmodel"
    "testing"
)



func TestService_SendSMSBatch(t *testing.T) {
    t.Run("SendSMSBatch", func(t *testing.T) {
        ctx := context.Background()

        req := []*rpcmodel.SendSmsBatch{{
            Phone: "0930049641",
            Countrycode: "886",
            Message: "RD++++Jim",
            Signaturecode:  "TutorABC",
        }}
        e := SendSMSBatch(ctx, req)
        fmt.Println(e)
    })
}






