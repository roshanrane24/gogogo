package main

import (
    "fmt"
)

func main() {
    // For Loop
    for i := 0; i < 10; i++ {
        fmt.Println(i)
    }
    label: 
        for i := 1; i < 25; i++ {
            for j := 1; j < 55; j++ {
                if i == j { 
                    fmt.Println("Break")
                    break
                } else if i * 2 == j * 15 {
                    fmt.Println("Breaking label")
                    break label
                }
                fmt.Println(i * j)
            }
        }

    fmt.Println("")
    // For as while
    i := 0
    for {
        i++
        if i % 2 == 0 { continue }
        fmt.Println(i)
        if i > 10 { break }
    }
    
    fmt.Println("")
    arr := []int{1,5,2,5,65,21}
    for k, v := range arr {
        fmt.Println(k, v)
    }
}
