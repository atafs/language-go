package main

import (
    "fmt"
    "encoding/json"
)

type Message struct {
    Name string
    Body string
    Time int64
}

func main() {

  //Encoding
  message := Message{"Americo", "Hello", 1294706395881547000}
  encode, err := json.Marshal(message)
  if err != nil {
      fmt.Println("Error when encoding JSON data")
      return
  }
  fmt.Println(encode, "\n")

  //Decoding
  var message_decoded Message
  err1 := json.Unmarshal(encode, &message_decoded)
  if err1 != nil {
      fmt.Println("Error when encoding JSON arbitrary data")
      return
  }
  fmt.Println(message_decoded, "\n")
  fmt.Println(message_decoded.Name, "\n")
  fmt.Println(err1, "\n")

  //Decoding arbitrary data
  b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)

  var f interface{}
  err2 := json.Unmarshal(b, &f)
  if err2 != nil {
      fmt.Println("Error when encoding JSON arbitrary data")
      return
  }

  message_decoded_arbitrary := f.(map[string]interface{})
  for k, v := range message_decoded_arbitrary {
      switch vv := v.(type) {
      case string:
          fmt.Println(k, "is string", vv, "\n")
      case int:
          fmt.Println(k, "is int", vv, "\n")
      case []interface{}:
          fmt.Println(k, "is an array:")
          for i, u := range vv {
              fmt.Println(i, u)
          }
          fmt.Println("")

      default:
          fmt.Println(k, "is of a type I don't know how to handle", "\n")
      }
  }

  //Generic JSON with interface{}
  var i interface{}
  i = "a string"
  i = 2011
  i = 2.777

  switch v := i.(type) {
  case int:
      fmt.Println("twice i is", v*2, "\n")
  case float64:
      fmt.Println("the reciprocal of i is", 1/v, "\n")
  case string:
      h := len(v) / 2
      fmt.Println("i swapped by halves is", v[h:]+v[:h], "\n")
  default:
      // i isn't one of the types above
  }
  





}
