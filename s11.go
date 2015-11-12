package main

import "fmt"

func fibonacci() func() int {
    pre0 := 0
    pre1 := 0
    now := 1
    return func() int {
        pre0 = pre1
        pre1 = now
        now = pre0 + pre1
        return now
    }
}

func main() {
    f := fibonacci()
    for i:=0; i<10; i++ {
        fmt.Println(f())
    }
    f2 := fibonacci()
    for i:=0; i<10; i++ {
        fmt.Println(f2())
    }
}

