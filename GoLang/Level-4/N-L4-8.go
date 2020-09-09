package main

import "fmt"

func main () {
  // EDIT HERE
  m := map[string][]string{
    `PC`: []string{`Hiking`, `Biking`, `Soccer`},
    `JC`: []string{`Soccer`, `XBOX`, `Band`},
    `WC`: []string{`LAX`, `XBOX`, `Youtube`},
  }
    for i, v := range m{
      fmt.Println("Record", i)
        for i2, v2 := range v {
          fmt.Printf("Index Position: %v\t Value: %v\n", i2, v2)
        }
    }
}
