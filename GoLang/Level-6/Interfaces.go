package main

import "fmt"

type person struct {
	first string
	last  string
}

type secret struct {
	person
	ltk bool
}

type human interface {
	foo()
}

func main () {
  // EDIT HERE
  sa1 := secret{
		person: person{
			"James",
			"Bond",
		},
		ltk: true,
	}

	sa2 := secret{
		person: person{
			"Miss",
			"Moneypenny",
		},
		ltk: true,
	}
  sa1.foo ()
  sa2.foo ()
  bar(sa1)
}

func (s secret) foo () {
  b := "Hi"
  fmt.Println(b, s.first, s.last)
}
func (p person) foo () {
  b := "Bye"
  fmt.Println(b)
}

func bar (h human) {
  fmt.Println("Goodbye", h.(secret).first)
}
