package main

import "fmt"

type vehicle struct {
  Doors   string
  Color  string
}

type truck struct {
  vehicle
  FourWheel bool
}

type sedan struct {
  vehicle
  Luxury bool
}

func main () {
  // EDIT HERE
  v1 := truck {
      vehicle: vehicle {
        Doors:  "4",
        Color:  "Black",
      },
        FourWheel:  true,
  }
  v2 := sedan {
      vehicle: vehicle {
        Doors: "4",
        Color:  "Red",
      },
      Luxury:    true,
  }
  fmt.Printf("Doors:\t %v\nColor:\t %v\n4WD:\t %v\n", v1.Doors, v1.Color, v1.FourWheel)
  fmt.Println(v2.Doors, v2.Color, v2.Luxury)
}
