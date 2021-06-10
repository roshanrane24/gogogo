package main

import (
    "fmt"
)

func main() {
    // Panic (Exception)
    a := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 21}

    for _, e := range a {
        if e % 2 != 0 { 
            err := fmt.Sprintf("%d is not an even number", e)
            panic(err)
        } else {
            fmt.Println(e)
        }
    }
}
