package main

import (
    "io/ioutil"
    "net/http"
		"encoding/json"
    "log"
)



func main() {

    type Message struct {
      agent string
      types string
      pages map[string]interface{}
      conversations map[string]interface{}
    }

    message := new(Message)
    count := 0

    agent := "americo@moonfruit.com"
    client := &http.Client{}

		base_url := "https://api.intercom.io"
		path := "/conversations?email=" + agent
		full_url  := base_url + path

    req, _ := http.NewRequest("GET", full_url, nil)
		req.SetBasicAuth("v39v2i4k", "f61f0c9e41dbc8292ce74d6262d3af4d30ef0202")
    req.Header.Add("Accept", "application/json")
		req.Header.Add("Content-Type", "application/json")

    resp, err := client.Do(req)
    if err != nil {
        log.Println("Error when sending request to the server")
        return
    }

    defer resp.Body.Close()
    body, _ := ioutil.ReadAll(resp.Body)

    log.Println(resp.Status, "\n")
		log.Println(resp.Header, "\n")

		//Unmarshal
		var f interface{}
	  err2 := json.Unmarshal(body, &f)
	  if err2 != nil {
	      log.Println("Error when Decoding JSON arbitrary data")
	      return
	  }

		body_decoded := f.(map[string]interface{})
		for k, v := range body_decoded {
        message = new(Message)

	      switch vv := v.(type) {
	      case string:
            message.types = vv;
            log.Println(message.types, " = ", vv)
	          log.Println(k, " is string ", vv, "\n")
	      case int:
	          log.Println(k, " is int ", vv, "\n")
	      case []interface{}:
	          log.Println(k, " is an array: ", "\n")
	          for i, u := range vv {
                log.Println("\n-----------------------")
	              log.Println(i, " = ", u, "\n")
                for k1, u1 := range u.(map[string]interface{}) {
                    log.Println(k1, " = ", u1)
                    if k1 == "conversation_message" {
                        message.conversations = u1.(map[string]interface{})
                    } else if k1 == "pages" {
                        message.pages = u1.(map[string]interface{})
                    }
                }

                log.Println("\n-----------------------")
                for k2, u2 := range message.conversations {
                    log.Println(k2, " = ", u2)
                }
                log.Println("-----------------------\n")
	          }
	      default:
	          log.Println(k, "is of a type I don't know how to handle")
	      }
        //data
        message.agent = agent;

        messages := make(map[int]string)
        messages[count] = message.agent

        count += 1
        log.Println("messages[count]", messages[count], "\n")


	  }

}
