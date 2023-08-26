package main

import (
  "os"
  "fmt"
  "path/filepath"
  "regexp"
)

func findContentMatch(content string, match string) bool {
  regex, err := regexp.Compile(match)
  if err != nil {
    fmt.Println("Error compiling regex", match)
    os.Exit(1)
  }

  return regex.MatchString(content)
}

func visit(path string, info os.FileInfo, err error) error {
	if err != nil {
		fmt.Printf("Error accessing path %q: %v\n", path, err)
		return err
	}

  if(!info.IsDir()){
    content, err := os.ReadFile(path)
    if err != nil {
      fmt.Println("Error reading file", path)
      os.Exit(1)
    }
     
    arg := os.Args[1]
    found := findContentMatch(string(content), arg)

    if(found){
      fmt.Println(path)
      os.Exit(0)
    }
  }

	return nil
}

func main(){
  root := "."
  err := filepath.Walk(root, visit)
  if err != nil {
    fmt.Printf("Error walking the path %q: %v\n", root, err)
    os.Exit(1)
  }

  os.Exit(0)
}
