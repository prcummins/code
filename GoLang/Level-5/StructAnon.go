package main

import "fmt"

func main () {
  // EDIT HERE
  cg1 := struct {
          Game      string
          Category  string
          Ranking   int
    }{
          Game:       "Fortnite",
          Category:   "FPS",
          Ranking:    1,
    }

  fmt.Println(cg1.Game, cg1.Category, cg1.Ranking)
}
