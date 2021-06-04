package rpc

import (
    "fmt"
    "goldtalkAPI/pkg/util"
    "testing"
)



func TestService_GenTutorBySlug(t *testing.T) {
    t.Run("A", func(t *testing.T) {

        r , e := GetClassInformationByLobbySn(232020)
        fmt.Println(e)
        fmt.Println("###")
        fmt.Println(util.JsonString(r))
        fmt.Println("###")

    })
}






