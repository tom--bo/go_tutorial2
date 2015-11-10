package main

import (
    "fmt"
    "net"
    "time"
    "math"
)

func main() {
    fmt.Println("Hello, 世界")
    fmt.Println("Time is", time.Now())
    fmt.Println("access network:")
    fmt.Println(net.Dial("tcp", "google.com"))
    fmt.Println(math.Pi)
}
