package main

import "fmt"

type Vertex struct {
    Lat, Long float64
}

var m = map[string]Vertex {
    "Bell Labs": {40.68433, -74.399},
    "Google": {37.4, -122.08},
}

func main (){
    fmt.Println(m)
}
