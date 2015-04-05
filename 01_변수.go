package main

import "fmt"

var p = fmt.Println

func main() {
	//초기값이 없는 경우 zero-value 를 나타내게 됩니다.
	var money int
	fmt.Println("나는 태어날 때부터", money, "원을 가지고 태어났습니다")

	var developer = true // 변수 선언하면서 타입추론이 가능합니다.
	fmt.Println("당신은 개발자입니까?", developer)

	solo := true // 중요한 부분입니다. 변수를 선언과 동시에 값을 입력하면서 var 도 생략가능합니다.
	// 즉 이 문장은 var solo bool = true 가 var solo = true 과정을 거쳐서  solo := true 로 최종 변환된 것이죠.
	fmt.Println("당신은 솔로부대입니까? ", solo)

	p("hello world") //이렇게도 줄여쓸수도 있겠습니다^^;
	p2 := fmt.Println
	p2("줄여써보기도 잘 되죠? ^^;")
}
