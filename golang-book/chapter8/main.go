package main
import "fmt"

func swap(xPtr, yPtr *int) {
    tmp := *xPtr
    *xPtr = *yPtr
    *yPtr = tmp
}

func main() {
    x := 1
    y := 2

    fmt.Println(x, y)
    swap(&x, &y)
    fmt.Println(x, y)

}
