package main

import(
    "fmt"
    "math"
)

func main() {
    var a int8 = 12
    var b int = 545
    var c int = 456
    var d float32 = 0.52
    var e float64 = 6.99
    var f float64 = 35.69

    // int operations
    fmt.Printf("Addition (int + int8)%d\n", a + int8(b))
    fmt.Printf("Addition (int + int)%d\n", b + c)

    fmt.Printf("Subtraction (int + int8)%d\n", a - int8(b))
    fmt.Printf("Subtraction (int + int)%d\n", b - c)

    fmt.Printf("Mutiplication (int + int8)%d\n", a * int8(b))
    fmt.Printf("Mutiplication (int + int)%d\n", b * c)

    fmt.Printf("Division (int + int8)%d\n", a / int8(b))
    fmt.Printf("Division (int + int)%d\n", b / c)

    fmt.Printf("Remiander %v\n", b % c)

    //Float Ofrations
    fmt.Printf("Addition (float64 + float32)%f\n", e + float64(d))
    fmt.Printf("Addition (float64 + float64)%f\n", e + f)

    fmt.Printf("Subtraction (float64 + float32)%f\n", f - float64(d))
    fmt.Printf("Subtraction (float64 + float64)%f\n", f - e)

    fmt.Printf("Mutiplication (float64 + float32)%f\n", float64(d) * f)
    fmt.Printf("Mutiplication (float64 + float64)%f\n", e * f)

    fmt.Printf("Division (float64 + float32)%f\n", e / float64(d))
    fmt.Printf("Division (float64 + float64)%f\n", f / e)

    fmt.Printf("Remiander %v\n", math.Remainder(math.Pi, 3))
}
