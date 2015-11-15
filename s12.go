package main
import(
    "fmt"
    "runtime"
)

func main() {
    fmt.Print("Go runs on ")
    switch os := runtime.GOOS; os{
    case "darwin":
        fmt.Println("OS X.")
    case "linux":
        fmt.Println("Linux.")
    default:
        fmt.Println("%s.", os)
    }

    switch i := 0; i{
    case 1:
        fmt.Println("1!")
    case 0:
        fmt.Println("0!")
    }
}


