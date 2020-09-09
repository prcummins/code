package main

import "fmt"

type games struct {
  Game string
  Category string
  Ranking int
}

func main () {
  // EDIT HERE
  g1 := games {
    Game: `Fortnite`,
    Category: `FPS`,
    Ranking: 1,
  }

  g2 := games {
    Game: `Madden`,
    Category: `Sports`,
    Ranking: 2,
  }
  fmt.Println(g1.Ranking, g2.Ranking)
}
