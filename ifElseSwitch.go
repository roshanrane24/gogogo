package main

import (
    "fmt"
)

func main() {
    // If/Else or || and &&
    number := 10
    guess :=111

    if guess < 1 || guess > 100 {
        fmt.Println("Range 1 ~ 100")
    } else {
        if guess < number {
            fmt.Println("Low")
        } else if guess > number {
            fmt.Println("High")
        } else {
            fmt.Println("Right!!!!")
        }
    }


    // Switch
    switch i := 2 + 3;i {
    case 1:
        fmt.Println("One")
    case 2:
        fmt.Println("Two")
    case 3, 4, 5:
        fmt.Println("3 || 4 || 5")
    default:
        fmt.Println("Not 1 2 3 4 5")
    }

    j := 10
    switch {
    case j <= 10:
        fmt.Println("Less than 10")
    case j <= 25:
        fmt.Println("Less than 25")
    case j <= 50:
        fmt.Println("Less than 50")
    case j <= 100:
        fmt.Println("Less than 100")
    default:
        fmt.Println("More than 100")
    }

    switch {
    case j <= 10:
        fmt.Println("Less than 10")
        fallthrough
    case j <= 25:
        fmt.Println("Less than 25")
        fallthrough
    case j <= 50:
        fmt.Println("Less than 50")
        fallthrough
    case j <= 100:
        fmt.Println("Less than 100")
        fallthrough
    default:
        fmt.Println("More than 100")
    }

}
