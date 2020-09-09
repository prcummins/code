package main

import "fmt"

type games struct {
  Game string
  Category string
  Ranking int
}

type console struct {
  games
  Console string
}

func main () {
  // EDIT HERE
  cg1 := console {
      games: games {
        Game: `Fortnite`,
        Category: `FPS`,
        Ranking: 1,
      },
      Console: `XBOX`,
  }

  g2 := games {
    Game: `Madden`,
    Category: `Sports`,
    Ranking: 2,
  }
  fmt.Println(cg1.Ranking, cg1.Console, g2.Ranking)
}
