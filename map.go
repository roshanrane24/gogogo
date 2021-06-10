package main

import (
    "fmt"
)

func main() {
    // Map
    // map[keyType]valueType{key: value}
    level := map[string]int{
        "High": 3,
        "Med":  2,
        "Low":  1,
    }

    fmt.Println(level)
    fmt.Println(level["High"])

    level["Empty"] = 0

    fmt.Println(level)

    delete(level, "Empty")
    fmt.Println(level)

    fmt.Println(level["Empty"])
    emp, ok := level["Empty"]
    fmt.Println(ok, emp)
    h, ok := level["High"]
    fmt.Println(ok, h)
}
