package main

import (
    "fmt"
)

type Vertex struct {
    X, Y int
}

func main() {
    v := new(Vertex)
    fmt.Println(v)
    v.X, v.Y = 11, 9
    fmt.Println(v)

    p := []int{2,3,5,7,9}
    fmt.Println("p == ", p)

    for i:=0; i< len(p); i++ {
        fmt.Printf("p[%d] == %d\n", i, p[i])
    }
    fmt.Println("p[:3] ==", p[:3])
    fmt.Println("p[1:4] ==", p[1:4])
    fmt.Println("p[3:] ==", p[3:])

    a := make([]int, 5)
    printSlice("a", a)
    b := make([]int, 0, 5)
    printSlice("b", b)
    b = b[:cap(b)]
    b[1] = 1
    b[3] = 3
    printSlice("b", b)
    c := b[:2]
    printSlice("c", c)
    d := c[2:5]
    printSlice("d", d)
}

func printSlice(s string, x []int) {
    fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

