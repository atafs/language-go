package main

import (
  "time"
  "log"
)

func main() {
  var Ball int
  table := make(chan int)
  const PLAYERS = 100

  for i := 0; i < PLAYERS; i++ {
    go player(table, i)
  }

  table <- Ball

  //time for the match
  time.Sleep(1000 * time.Millisecond)
  <-table
}

func player(table chan int, player_number int) {
    for {
        ball := <-table
        ball++

        //time for each play
        time.Sleep(50 * time.Millisecond)

        table <- ball
        log.Println("table:", table)

        log.Println("Player", player_number, "did his shot!!")
    }
}
