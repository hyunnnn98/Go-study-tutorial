package main

// 변수 선언은 var 를 통해 선언.
// 변수 선언 후 변수 타입을 명시하자.
// 선언만하고 사용하지 않으면 garbage collector에 의해 자동으로 삭제된다.
var a int

var f float32 = 11.

// 동일한 타입의 변수가 복수 개 있을 경우.
var i, j, k int
var z, x, c int = 1, 2, 3

// 변수를 선언하는 방식으로 Short Assignment Statement (:=) 를 사용할 수 있다.
// 하지만 short 방식은 func 에서만 사용 가능.
var q int = 1

func test() {
	p := 1

	println(p)
}
