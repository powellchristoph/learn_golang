package main

import (
    "fmt"
)

func main() {

    mySlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    fmt.Println(mySlice[4])

    //mySlice[1] = 0
    //fmt.Println(mySlice)

    //sliceOfSlice := mySlice[2:5]
    //fmt.Println(sliceOfSlice)

    mySlice = append(mySlice, 11)
    fmt.Println(mySlice)

    newSlice := []int{}
    fmt.Println(newSlice)

    newSlice = append(newSlice, 1, 2, 3, 4, 5)
    fmt.Println(newSlice)





}
