package main

const c int = 10
const s string = "Hi"

// 뒤에 타입명을 적지 않아도 타입을 추론해준다. 타입스크립트와 비슷하네.

const (
	Visa   = "Visa"
	Master = "MasterCard"
	Amex   = "Aamerican Express"
)

// 유용한 팁. 상수값을 0 부터 순사척으로 부여하기 위해서는
// iota 라는 identifier를 사용할 수 있다.
const (
	Apple  = iota // 0
	Grape         // 1
	Orange        // 2
)
