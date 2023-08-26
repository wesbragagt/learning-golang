package main

import (
	"fmt"
	"log"
	"os"
)

func ListFiles(dir string){ 
  files, err := os.ReadDir(dir)

  if err != nil {
    log.Fatal(err.Error())
  }

  for _, file := range files {
    fmt.Println(file.Name())
  }
}

func main(){
  dir := os.Args[1]
  ListFiles(dir)
}
