package main

import (
    "fmt"
)

func main() {
    yourName("Roshan")

    // Pointer
    age := 23
    increaseByFive(&age)
    fmt.Println(age)

    // Variatic parameters
    sumIt(23, 42, 69, 83, 29)

    // Return
    fn := giveFullName("Roshan", "Rane")
    fmt.Println(fn)
    d, err := divide(45.0, 6.0)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(d)

    //methods
    me := person {
        name: "Roshan",
        last: "Rane",
        age: 23,
    }

    me.info()
}

func yourName(name string) {
    fmt.Println("Name:", name)
}

func increaseByFive(value *int)  {
    fmt.Println(value, *value)
    *value += 5
}

// Only works for last arg
func sumIt(values ...int)  {
    fmt.Println(values)
    result := 0
    for _, v := range values {
        result += v
    }
    fmt.Println(result)
}

// Return
func giveFullName(first string, last string) string  {
    return first + " " + last
}

func divide(a, b float64) (float64, error) {
    if b == 0.0 {
        return 0.0, fmt.Errorf("Cannot Divide by zero")
    }
    return a/b, nil
}

// methods
type person struct {
    name string
    last string
    age int
}
func (p person) info() {
    fmt.Println("\nInfo")
    fmt.Println("Name:", p.name, p.last)
    fmt.Println("Age:", p.age)
}
