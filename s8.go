package main

import "fmt"

func main() {
    m := make(map[string]int)

    m["Answer"] = 42
    fmt.Println("The value:", m["Answer"])

    m["Answer"] = 93
    fmt.Println("The value:", m["Answer"])

    v, ok := m["Answer"]
    delete(m, "Answer")
    fmt.Println("The value:", v, "Present?", ok)

    v2, ok2 := m["Answer"]
    fmt.Println("The value:", v2, "Present?", ok2)

}
