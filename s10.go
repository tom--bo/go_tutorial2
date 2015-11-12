package main
import "fmt"

func adder() func(int) int {
    sum := 0
    return func(x int) int{
        sum += x
        return sum
    }
}

func main() {
    pos, neg := adder(), adder()
    fmt.Println(
        pos(1),
        neg(-2),
    )
    fmt.Println(
        pos(1),
        neg(-2),
    )
}

