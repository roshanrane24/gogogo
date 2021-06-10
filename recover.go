package main

import (
    "fmt"
    "log"
)

func main() {
    fmt.Println("Start")
    panicNow()
    fmt.Println("End")
}

func panicNow () {
    fmt.Println("Paaaniiiicc")
    defer func ()  {
        if err := recover(); err != nil {
            log.Println("Error", err)
        }
    }()
    panic("This is Bad")
    fmt.Println("Calm")
}
