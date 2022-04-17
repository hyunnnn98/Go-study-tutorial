package main

import fmt "fmt"

// if 구문
// - if 구문의 조건식은 반드시 boolean식으로 표현되어야 한다.
// - if 1 또는 if 0 와 같은 표현은 사용할 수 없다.

func main() {
	var k int = 2

	if k == 1 {
		println("One")
	} else if k == 2 {
		println("Two")
	} else {
		println("Other")
	}

	/*
		if문에서 조건식을 사용하기 이전에 간단한 문장 (Optional Statement)을 함께 실행할 수 있다.
		이러한 optional statement 표현은 switch, for 등의 여러 문법에서 사용 가능하다.
	*/
	var i, max int = 2, 6

	if val := i * 2; val < max {
		println("True", val)
	} else {
		println("False")
	}

	check_fallthrough(1)
}

/* Switch 구문
- 다른언어의 switch 구문과 거의 비슷하지만 Go 만의 특이점 4가지가 있다.

1. switch 뒤에 expression이 없을 수 있음.
2. case문에 expression을 쓸 수 있음.
3. No default fall through
	- 다른 언어의 case 문은 break를 쓰지 않는 한 다음 case로 이동하지만, Go는 다음 case로 가지 않는다.
4. Type switch
	- 다른 언어의 switch는 일반적으로 변수의 값을 기준으로 case로 분기하지만, Go는 그 변수의 Type에 따라 case로 분기할 수 있다.

- Go에서는 컴파일러가 자동으로 break 문을 각 case문 블럭 마지막에 추가해준다. 따라서 코더가 하나하나 입력하지 않아도 괜찮다.
  만약, default break문을 활성화 시키고 싶지 않다면 fallthrough 문을 선언해주면 다음 케이스의 코드 블럭을 계속 실행하게 된다.
*/

func simpleSwitch() {
	var name string
	var category = 1

	switch category {
	case 1:
		name = "Paper Book"
	case 2:
		name = "eBook"
	case 3, 4:
		name = "Blog"
	default:
		name = "Other"
	}

	println(name)
}

func grade(score int) {
	switch {
	case score >= 90:
		println("A")
	case score >= 80:
		println("B")
	case score >= 70:
		println("C")
	case score >= 60:
		println("D")
	default:
		println("No Hope")
	}
}

func typeSwitchTest(i interface{}) {
	switch v := i.(type) {
	case nil:
		fmt.Println("x is nil")
	case int:
		fmt.Println("x is", v)
	case bool, string:
		fmt.Println("x is bool or string")
	default:
		fmt.Printf("type unknown %T\n", v)
	}
}

func Example_TypeSwitch() {
	typeSwitchTest("value")
	typeSwitchTest(23)
	typeSwitchTest(true)
	typeSwitchTest(nil)
	typeSwitchTest([]int{})

	//Output:
	//x is bool or string
	//x is 23
	//x is bool or string
	//x is nil
	//type unknown []int
}

func check_fallthrough(val int) {
	switch val {
	case 1:
		fmt.Println("1 이하")
		fallthrough
	case 2:
		fmt.Println("2 이하")
		fallthrough
	case 3:
		fmt.Println("3 이하")
		fallthrough
	default:
		fmt.Println("default 도달")
	}
}
