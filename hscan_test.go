package structs

import (
    "fmt"
	"testing"
    "github.com/dreamsxin/structs"
)

// go test -v -count=1 -run TestScan hscan_test.go
func TestScan(t *testing.T) {

    type Model struct {
        Str1    string
        Str2    string
        Int     int
        Bool    bool
        Ignored struct{} `redis:"-"`
    }

    model1 := &Model{
        Str1: "hello",
        Str2: "world",
        Int: 3,
        Bool: true,
    }

    structs.DefaultTagName = "redis"

    m := structs.Map(model1)
    fmt.Printf("%+v\n", m)

    var model2 Model
    // Scan all fields into the model.
    map1 := map[string]interface{}{
        "Str1": "hello",
        "Str2": "world",
        "Int": "3",
        "Bool": "true",
    }

    structs.Scan(&model2, map1)
    fmt.Printf("%+v\n", model2)

    var model3 Model
    // Scan all fields into the model.
    map2 := map[string]interface{}{
        "Str1": "hello",
        "Str2": "world",
        "Int": 3,
        "Bool": true,
    }

    structs.Scan(&model3, map2)
    fmt.Printf("%+v\n", model3)
}
