package main
import (
    "fmt"
)

func add(x,y int) int {
    return x+y
}


func swap(x,y string) (string, string) {
    return y,x
}

func split(sum int) (x,y int) {
    x = sum*4/9
    y = sum-x
    return
}

const (
    Big   = 1 << 100
    Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
    return x * 0.1
}

func main() {
    fmt.Println(add(43, 13))
    a, b := swap("hello", "world")
    fmt.Println(a,b)
    fmt.Println(split(17))

    var i,j int = 1,2
    k := 3
    c, python, java := true, false, "no"
    fmt.Println(i, j, k, c, python , java)

    fmt.Println(needInt(Small))
    fmt.Println(needFloat(Small))
    fmt.Println(needFloat(Big))

}


