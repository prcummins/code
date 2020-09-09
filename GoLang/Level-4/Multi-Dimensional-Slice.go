package main

import "fmt"

func main () {
  // EDIT HERE
  pc := []string{"Paul", "Cummins", "Chocolate", "Mule"}
  fmt.Println(pc)
  jc := []string{"John", "Cummins", "Vanilla", "Sunkist"}
  fmt.Println(jc)

  pjc := [][]string{pc, jc}
  fmt.Println(pjc)
}
