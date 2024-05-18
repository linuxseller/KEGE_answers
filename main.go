package main

import (
    "fmt"
    "net/http"
    "io"
    "log"
    "encoding/json"
)

type Data struct {
	Tasks []Task `json:"tasks"`
}

type Task struct {
    Number int `json:"number"`
    Key string `json:"key"`
}

func main(){
    var data Data
    var variant_id string
    fmt.Print("Enter variant id: ")
    _, err := fmt.Scanf("%s", &variant_id)
    if err != nil {
        log.Println("Error: wrong variant id")
    }
    resp, err := http.Get("https://kompege.ru/api/v1/variant/kim/"+variant_id)
    if err != nil {
        log.Fatal("Error: could not establiish a connection")
    }
    defer resp.Body.Close()
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        fmt.Println(err)
        log.Fatal("Error: could ot ReadAll")
    }
    err = json.Unmarshal(body, &data)
    if err != nil {
        fmt.Println(err)
        log.Fatal("Error: could not parse json")
    }
    for _, task := range data.Tasks {
		fmt.Printf("№%d: %s\n", task.Number, task.Key)
	}
}
