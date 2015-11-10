package main

import (
    "fmt"
    "math"
)

type Vertex struct {
    X int
    Y int
}

func sqrt(x float64) string {
    if x < 0 {
        return sqrt(-x) + "i"
    }
    return fmt.Sprint(math.Sqrt(x))
}

func main() {
    sum := 0
    for i:= 0; i< 10; i++ {
        sum += i
    }
    fmt.Println(sum)

    fmt.Println(sqrt(2), sqrt(-4))
    p := Vertex{1,2}
    q := &p
    q.X = 1e9
    fmt.Println(p)
}
