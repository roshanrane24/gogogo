package main

import (
    "fmt"
)

func main() {
    someValue := "Some"
    fmt.Println("Start")
    defer fmt.Println(someValue)
    defer fmt.Println("Middle")
    fmt.Println("Somewhere in Middle")
    defer fmt.Println("End")
    someValue = "New Some"
}
