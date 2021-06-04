package dao

import (
    "fmt"
    "testing"
    "github.com/stretchr/testify/assert"
)




func TestService_GetTutorIDByLanguageID(t *testing.T) {
    t.Run("GetTutorIDByLanguageID", func(t *testing.T) {

        data, err := GetTutorIDByLanguageID(1)
        fmt.Println(data)
        assert.Equal(t, err, nil)
    })
}





