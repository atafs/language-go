package main

import "fmt"

func main() {
  var x [5]int
  x[4] = 100
  fmt.Println(x)

  slice1 := []int{1,2,3}
  slice2 := append(slice1, 4, 5)
  fmt.Println(slice1, slice2)

  y := make(map[int]int)
  y[1] = 10
  fmt.Println(y[1])

  z := make(map[int]string)
  z[1] = "Hello"
  fmt.Println(z[1])
}
