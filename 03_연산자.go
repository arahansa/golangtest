package main

import "fmt"

func main() {
	fmt.Println("3 더하기 4 =", 3+4) //7

	fmt.Println("AND, 둘다 참일때 참을 반환 & : ", 0&2)  //0
	fmt.Println("OR, 둘 중의 하나만 참이면 참을 반환 ", 0|3) // 3
	fmt.Println("NOT, 두 값이 서로다르면 참을 반환 ", 0^1)  // 1
	fmt.Println("AND NOT", 3&^5)

	fmt.Println(">> 연산자는 비트를 오른쪽으로 뒤의 숫자만큼 이동 ", 16>>2) //4
	fmt.Println("<<연산자는 비트를 왼쪽으로 뒤의 숫자만큼 이동 ", 2<<3)    //1

	fmt.Println(`AND`)
	fmt.Printf("%b & %b results in %03b\n", 4, 5, 4&5)
	fmt.Printf("%b & %b results in %03b\n", 5, 4, 5&4)
	fmt.Println(`OR`)
	fmt.Printf("%b | %b results in %03b\n", 4, 5, 4|5)
	fmt.Printf("%b | %b results in %03b\n", 5, 4, 5|4)
	fmt.Println(`XOR`)
	fmt.Printf("%b ^ %b results in %03b\n", 4, 5, 4^5)
	fmt.Printf("%b ^ %b results in %03b\n", 5, 4, 5^4)
	fmt.Println(`AND NOT`)
	fmt.Printf("%b &^ %b results in %03b\n", 7, 5, 7&^5)
	fmt.Printf("%b &^ %b results in %03b\n", 5, 7, 5&^7)
}
