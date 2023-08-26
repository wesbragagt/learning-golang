package main

import (
	"fmt"
	"io"
	"net/http"
)

// example of a go routine that fetches todos from jsonplaceholder and returns the body
func getTodos() []byte {
  resultChannel := make(chan []byte)

  go func(){
    res, err := http.Get("https://jsonplaceholder.typicode.com/todos") 
    if err != nil {
      panic(err)
    }

    body, err := io.ReadAll(res.Body)
    if err != nil {
      panic(err)
    }

    resultChannel <- body
  }()

  return <-resultChannel
}



func main(){
  fmt.Println(string(getTodos()))
}
