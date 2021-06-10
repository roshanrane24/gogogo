package main

import (
    "fmt"
)

func main() {
    // prefix * to type
    // & address of
    // dereferencing :* infront of pointer 
    a := 27
    b := a

    fmt.Println(a, b)

    b = 42

    fmt.Println(a, b)

    var c *int = &a

    fmt.Println(a, c)
    fmt.Println(&a, c)
    fmt.Println(*&a, *c)

    *c = 42

    fmt.Println(a, c)
}
