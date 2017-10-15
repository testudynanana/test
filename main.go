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
    jsonStr := `
{
  "name": "日本",
  "prefectures": [
    {
      "name": "東京都",
      "capital": "東京",
      "population": 13482040
    },
    {
      "name": "埼玉県",
      "capital": "さいたま市",
      "population": 7249287
    },
    {
      "name": "神奈川県",
      "capital": "横浜市",
      "population": 1111111
    }
  ]
}
`
    jsonBytes := ([]byte)(jsonStr)
    data := new(Country)

    if err := json.Unmarshal(jsonBytes, data); err != nil {
        fmt.Println("JSON Unmarshal error:", err)
        return
    }

    fmt.Println(data.Name)
    fmt.Println(data.Prefectures[0].Name)
    fmt.Println(data.Prefectures[1].Capital)
    fmt.Println(data.Prefectures[2].Population)
}
