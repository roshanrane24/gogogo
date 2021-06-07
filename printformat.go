package main

import (
	"fmt"
)

func main() {
	// Binary base2
	fmt.Printf("Binary (10,3): %b,%b\n", 10, 3)

    // Octal base8
	fmt.Printf("Octal (10,3): %o,%o\n", 10, 3)

    // Hex base16
	fmt.Printf("Hexadecimal (10,3): %x,%x\n", 10, 3)

	// String/Character
	name := "Roshan"
	lastName := "Rane"
	fmt.Printf("Name: %s\n", name)

	//Character
	gender := 'M'
	fmt.Printf("Gender: %c\n", gender)
	fmt.Printf("Gender(Quoted): %q\n", gender)
	fmt.Printf("Gender(Unicode): %U\n", gender)
	fmt.Printf("Gender(Unicode + Quoted): %#U\n", gender)

	// Integer/Float
	age := 23
	var height float32 = 1.75
	fmt.Printf("Age: %v,\nHeight: %.2f\n", age, height)


	// Boolean
	var usingGo bool = true
	fmt.Printf("Using Go: %t\n\n", usingGo)

	//Pointer
	fmt.Printf("Pointers:\nName: %p\nAge: %p\nHeight: %p\n", &name, &age, &height)

	fmt.Printf("Pointers:\nName: %v\nAge: %v\nHeight: %v\n", name, age, height)

    // Types
    fmt.Printf("Types:\n Name: %T\nAge: %T\nHeight: %T\nGender: %T\nUsing Go: %T\n", name, age, height, gender, usingGo)

    // Missing Values
    fmt.Printf("Missing Value %v\n")
	fmt.Print("Full Name: ", name + " " + lastName)
}
