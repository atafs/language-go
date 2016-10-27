package main

import (
    "log"
    "sync"
    "time"
)

const (
  WORKERS = 36
  TASKS = 75
)

//Workers or Fan-Out Pattern
func worker(tasksCh <-chan int, wg *sync.WaitGroup) {
    defer wg.Done()
    for {
        task, ok := <-tasksCh
        if !ok {
            return
        }
        d := time.Duration(task) * time.Millisecond
        time.Sleep(d)
        log.Println("...processing task", task)
        time.Sleep(500 * time.Millisecond)
    }
}

func pool(wg *sync.WaitGroup, workers, tasks int) {
    tasksCh := make(chan int)

    for i := 0; i < workers; i++ {
        go worker(tasksCh, wg)
        log.Println("Workers: tasksCh", tasksCh, ":", "number : ", i)
        time.Sleep(75 * time.Millisecond)
    }

    for i := 0; i < tasks; i++ {
        tasksCh <- i
        log.Println("Tasks: tasksCh", tasksCh, ":", "number : ", i)
        time.Sleep(10 * time.Millisecond)
    }

    close(tasksCh)
}

func main() {
    var wg sync.WaitGroup
    wg.Add(WORKERS)
    go pool(&wg, WORKERS, TASKS)
    wg.Wait()
}
