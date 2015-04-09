// 05_for_gugudan
package main

import (
	"fmt"
)

func main() {
	for i := 2; i < 10; i++ {
		for j := 1; j < 10; j++ {
			fmt.Println(i, " 곱하기 ", j, " 는 ", i*j)
		}
		fmt.Println("=====")
	}
}
