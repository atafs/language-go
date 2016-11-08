package main

 import (
     "os"
     "fmt"
 )

 func main () {
     // open file
     file, err := os.OpenFile("output.json", os.O_APPEND|os.O_WRONLY,0644)
     if err != nil {
         panic(err)
     }
     defer file.Close()

     if _, err = file.WriteString("\n this is the APPENDED text"); err != nil {
      panic(err)
     }

    fmt.Printf("Appended into file\n")
 }
