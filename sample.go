package main

import (
    "encoding/json"
    "fmt"
)

type Country struct {
    Name string              `json:"name"`
    Prefectures []Prefecture `json:"prefectures"`
}

type Prefecture struct {
    Name string    `json:"name"`
    Capital string `json:"capital"`
    Population int `json:"population"`
}

func main() {
    tokyo := Prefecture{Name:"東京都", Capital:"東京", Population:13482040}
    saitama := Prefecture{Name:"埼玉県", Capital:"さいたま市", Population:7249287}
    kanagawa := Prefecture{Name:"神奈川県", Capital:"横浜市", Population:9116252}
    japan := Country{
        Name:"日本",
        Prefectures:[]Prefecture{tokyo, saitama, kanagawa},
    }

    jsonBytes, err := json.Marshal(japan)
    if err != nil {
        fmt.Println("JSON Marshal error:", err)
        return
    }

    fmt.Println(string(jsonBytes))
}
