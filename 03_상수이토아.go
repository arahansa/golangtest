package main

import "fmt"

const (
	One = 1 * iota
	Two
	Three
)

func main() {
	fmt.Println("안녕하세요^^")
	fmt.Println(One)
	fmt.Println(Two)
	fmt.Println(Three)
}
