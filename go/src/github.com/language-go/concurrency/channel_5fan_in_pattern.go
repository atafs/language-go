package main

import (
    "log"
    "time"
)

func producer(ch chan int, d time.Duration, producer_number int) {
    var i int
    for {
        ch <- i
        i++
        log.Println("Producer", producer_number, ":", "Data : ", i)

        time.Sleep(d)
    }
}

func reader(out chan int) {
    for x := range out {
        log.Println("Reading : ", x)
    }
}

func main() {
    ch := make(chan int)
    out := make(chan int)
    go producer(ch, 100*time.Millisecond, 1)
    go producer(ch, 250*time.Millisecond, 2)
    go reader(out)

    for i := range ch {
        out <- i
    }
}
