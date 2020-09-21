package main

import (
        "fmt"
        "github.com/rocketlaunchr/dataframe-go"
)

func main () {
  s1 := dataframe.NewSeriesInt64("day", nil, 1, 2, 3, 4, 5, 6, 7, 8)
	s2 := dataframe.NewSeriesFloat64("sales", nil, 50.3, 23.4, 56.2, nil, nil, 84.2, 72, 89)
	df := dataframe.NewDataFrame(s1, s2)

	fmt.Print(df.Table())
  fmt.Printf("\n")
  fmt.Printf("-------------APPEND -------------------\n")
  fmt.Printf("\n")

  df.Append(nil, 9, 123.6)

	df.Append(nil, map[string]interface{}{
		"day":   10,
		"sales": nil,
	})

  fmt.Print(df.Table())
  fmt.Printf("\n")
  fmt.Printf("-------------REMOVE -------------------\n")
  fmt.Printf("\n")

	df.Remove(0)

	fmt.Print(df.Table())
  fmt.Printf("\n")
  fmt.Printf("-------------UPDATE -------------------\n")
  fmt.Printf("\n")

  df.UpdateRow(0, nil, map[string]interface{}{
	"day":   3,
	"sales": 45,
  })
  fmt.Print(df.Table())
}
