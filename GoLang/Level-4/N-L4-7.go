package main

import "fmt"

func main () {
  // EDIT HERE
  jb := []string{"James", "Bond", "Shaken, not stirred"}
  mp := []string{"Miss", "Monneypenny", "Helloooo, James"}
  jbmp := [][]string{jb, mp}
  fmt.Println(jbmp)
  for i, v := range jbmp{
    fmt.Println("Record:", i)
      for i2, v2 := range v {
        fmt.Printf("\t Index Position: %v \t Value: %v \n", i2, v2)
      }
  }
}
