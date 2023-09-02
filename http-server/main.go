package main

import (
	"fmt"
	"io"
	"net/http"
)

func golangWebsite(res http.ResponseWriter, req *http.Request){
  io.WriteString(res, "What's up dude!?")
}

func main(){
  http.HandleFunc("/", golangWebsite)

  fmt.Println("Server is running on port 3000")
  err := http.ListenAndServe(":3000", nil)

  if err != nil {
    panic(err)
  }
}
