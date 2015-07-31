package main
import "fmt"

func average(xs []float64) float64 {
    total := 0.0
    for _, v := range xs {
        total += v
    }
    return total / float64(len(xs))
}

func sum(args ...float64) float64{
    total := 0.0
    for _, v := range args {
        total += v
    }
    return total
}

func half(num int) bool{
    half := int(num / 2)
    if half % 2 == 0 {
        return true
    } else {
        return false
    }
}

func main() {
    xs := []float64{98, 93, 77, 82, 83}

    fmt.Println("Average: ", average(xs))
    fmt.Println("Add: ", sum(xs...))

    fmt.Println("half(1): ", half(1))
    fmt.Println("half(2): ", half(2))

}
