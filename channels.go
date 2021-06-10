package main

import (
    "fmt"
    "sync"
)

var wg = sync.WaitGroup{}

func main() {
    ch := make(chan int)
    wg.Add(2)
    go func(){
        i := <- ch
        fmt.Println("INIT",i)
        wg.Done()
    }()
    go func(i int) {
        ch <- i
        i = 10
        fmt.Println("Changed", i)
        wg.Done()
    }(42)

    counter := 0
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func(ch chan<- int) {
            ch <- counter * 15
            counter++
            wg.Done()
        }(ch)
        if counter == 9{
            close(ch)
        }
    }

    go func(ch <-chan int)  {
        for {
            if i, ok := <- ch; ok{
                fmt.Println(i)
            }
        }
        wg.Done()
    }(ch)

    wg.Wait()
}
