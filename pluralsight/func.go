package main

import (
    "fmt"
    "strings"
)

func main () {
    
    module := "function basis"
    author := "chris powell"

    fmt.Println(converter(module, author))
}

func converter(author, module string) (string, string) {
    s1 := strings.Title(module)
    s2 := strings.ToUpper(author)

    return s1, s2
}
