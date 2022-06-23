package main

import (
	"fmt"
)

//import "github.com/Jeffail/tunny"

type People struct {
	name string
	tele string
	age  int
}

type alive interface {
	walk()
}

//type Man struct {
//	People
//}

func (receiver People) walk() {
	fmt.Println("呵呵")
}
func main() {
	//tunny.Pool{}

	//m := Man{}
	//m.walk()
	//m.People.walk()
	fmt.Println("hello word")
}
