package main

import "fmt"


func main() {
	x := 1
	for {
		if x > 100 {
			break
		}

		if x%2 != 0 {
			continue
		}
    x++
		fmt.Println(x)

	}
	fmt.Println("done.")
}
