package dao

import (
    "fmt"
    "testing"
    "github.com/stretchr/testify/assert"
)



func TestService_GetLessonPricesMapByTutorIDs(t *testing.T) {
    t.Run("GetLessonPricesMapByTutorIDs", func(t *testing.T) {

        data, err := GetLessonPricesMapByTutorIDs([]int64{1,2})
        fmt.Println(data)
        assert.Equal(t, err, nil)
    })
}





