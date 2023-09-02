package main
import (
  "os"
  "path/filepath"
  "fmt"
  "strings"
)

func visit(path string, fileInfo os.FileInfo, err error) error {
  if err != nil {
    fmt.Println(err)
    return nil
  }
  
  searchQuery := os.Args[1]
  if !fileInfo.IsDir() && strings.Contains(path, searchQuery){
    fmt.Println(path)
  }

  return nil
}
  
func main(){
  err := filepath.Walk(".", visit)
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}
