package main

import (
  "time"
  "log"
)

func main() {
    var Ball int
    table := make(chan int)
    go player(table, 2)
    go player(table, 1)
    go player(table, 3)
    go player(table, 4)

    table <- Ball

    //time for the hole match
    time.Sleep(5000 * time.Millisecond)
    <-table
}

func player(table chan int, player_number int) {
    for {
        ball := <-table
        ball++

        //time for each play
        time.Sleep(200 * time.Millisecond)

        table <- ball
        log.Println("table:", table)

        log.Println("Player", player_number, "did his shot!!")
    }
}
