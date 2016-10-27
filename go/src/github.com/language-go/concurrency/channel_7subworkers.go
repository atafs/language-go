package main

import (
    "log"
    "sync"
    "time"
)

const (
  WORKERS    = 5
  SUBWORKERS = 3
  TASKS      = 20
  SUBTASKS   = 10
)

func subworker(subtasks chan int) {
    for {
        task, ok := <-subtasks
        if !ok {
            return
        }
        time.Sleep(time.Duration(task) * time.Millisecond)
        log.Println("...processing task", task)
    }
}

func worker(tasks <-chan int, wg *sync.WaitGroup) {
    defer wg.Done()
    for {
        task, ok := <-tasks
        if !ok {
            return
        }

        subtasks := make(chan int)
        for i := 0; i < SUBWORKERS; i++ {
            go subworker(subtasks)
            log.Println("SUB Workers: subtasks", subtasks, ":", "number : ", i)
        }
        for i := 0; i < SUBTASKS; i++ {
            task1 := task * i
            subtasks <- task1
            log.Println("SUB Tasks: subtasks", subtasks, ":", "number : ", i)
        }
        close(subtasks)
    }
}

func main() {
    var wg sync.WaitGroup
    wg.Add(WORKERS)
    tasks := make(chan int)

    for i := 0; i < WORKERS; i++ {
        go worker(tasks, &wg)
        log.Println("Workers: tasks", tasks, ":", "number : ", i)
    }

    for i := 0; i < TASKS; i++ {
        tasks <- i
        log.Println("Tasks: tasks", tasks, ":", "number : ", i)
    }

    close(tasks)
    wg.Wait()
}
