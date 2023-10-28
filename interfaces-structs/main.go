package main

import "fmt"

type Blogger interface {
  GetAuthor() string
}

type Author struct {
  id string
  name string
  posts []Post
}

type Post struct {
  id string
  author string
  content string
}

func (p Post) GetAuthor() string {
  return p.author
}

func main(){

  p := Post{
    id: "1",
    content: "some post",
    author: "Michael Jackson",
  }

  result := p.GetAuthor()

  fmt.Println(result)
}
