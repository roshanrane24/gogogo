package main

import (
    "fmt"
    "sync"
    "runtime"
)

var wg = sync.WaitGroup{}
var counter = 0
// Lock Unlock
var m = sync.RWMutex{}

func main() {
    fmt.Println("Threads:", runtime.GOMAXPROCS(-1))
    runtime.GOMAXPROCS(64)
    fmt.Println("Threads:", runtime.GOMAXPROCS(-1))
    for i := 1; i < 20; i++ {
        wg.Add(2)
        m.Lock()
        go hello()
        m.Lock()
        go inc()
    }
    wg.Wait()
}

func hello()  {
    fmt.Println("Hello There!", counter)
    m.Unlock()
    wg.Done()
}

func inc() {
    counter++
    m.Unlock()
    wg.Done()
}
