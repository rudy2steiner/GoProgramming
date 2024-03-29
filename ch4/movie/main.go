package main

import (
    "encoding/json"
    "fmt"
    "log"
)

type Movie struct {
    Title string
    Year int `json:"released"`
    Color bool `json:"color,omitempty"`
    Actors []string
}

var movies = []Movie {
   {Title:"Casablanca",Year:1942,Color:false,Actors:[]string{"Humphrey Bogart", "Ingrid Bergman"}},
   {Title:"Cool Hand luke", Year:1967, Color:true, Actors:[]string{"Paul Newman"}},
   {Title:"Bullitt",Year:1968,Color:true,Actors:[]string{"Steve McQueen"}},
}

func main() {
    data, err := json.Marshal(movies)
    if err != nil {
        log.Fatalf("Json marshaling failed:%s", err)
    }
    fmt.Printf("%s", data)
}