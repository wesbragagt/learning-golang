package main

import (
  "testing"
)

func hello() string {
  return "Hello dude!"
}

func TestHello(t *testing.T){
  want := "Hello dude!"
  result := hello()
  if result != want {
    t.Fatalf("Expected Hello dude!, but got %s instead", result)
  }
}
