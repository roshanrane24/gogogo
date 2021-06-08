package main

import(
    "fmt"
)

func main() {
    // Array
    arr := [5]int{65, 21, 21}
    arr2 := [3]int{65, 21, 21}
    arr3 := [...]int{12, 54, 32, 12, 23}
    var students [3]string
    
    fmt.Printf("Array: %v\n", arr)
    fmt.Printf("Array: %v\n", arr2)
    fmt.Printf("Array: %v\n", arr3)
    fmt.Printf("Student: %v\n", students)
    
    students[0] = "Roshan"

    fmt.Println(students)
    fmt.Println("Number of student ", len(students))
    students[1] = "Izuku"
    students[2] = "Eren"

    fmt.Println(students)
    fmt.Println("Number of student ", len(students))

    // Matrix
    var matrix [4][3]int
    fmt.Println(matrix)

    matrix[0] = [3]int{1, 2, 3}
    matrix[1] = [3]int{4, 5, 6}
    matrix[2] = [3]int{7, 8, 9}
    matrix[3] = [3]int{10, 11, 12}

    fmt.Println(matrix)

    matrix1 := matrix
    matrix1[0] = [3]int{13, 14, 15}
    fmt.Println("Matrix:\n", matrix, "\nMatrix1:\n", matrix1)

    matrix2 := &matrix1
    matrix2[3] = [3]int{16, 17, 18}
    fmt.Println("Matrix1:\n", matrix1, "\nMatrix2:\n", matrix2)
    // Slices
    matrix3 := matrix[:]
    matrix3[3] = [3]int{19, 20, 21}
    fmt.Println("Matrix:\n", matrix, "\nMatrix3:\n", matrix3)

    //Make
    a := make([]int, 3, 10)
    a = []int{1,2,3}
    fmt.Println(a)
    fmt.Println(len(a))
    fmt.Println(cap(a))
    a = append(a, 4)
    fmt.Println(a)
    fmt.Println(len(a))
    fmt.Println(cap(a))
    a = append(a, []int{5,6,7,8,9}...) //Spread *(,) in python
    fmt.Println(a)
    fmt.Println(len(a))
    fmt.Println(cap(a))

    //pop
    //left
    a = a[1:]
    fmt.Println(a)
    //right
    a = a[:len(a)-1]
    fmt.Println(a)
    //mid
    i := 3
    a = append(a[:i], a[i+1:]...)
    fmt.Println(a)
}
