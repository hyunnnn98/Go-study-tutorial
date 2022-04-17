package main

// mainパッケージのmain関数は
// いくつかの初期化の後に最初に呼ばれるプログラムの入り口になります。

import fmt "fmt" // フォーマット付き入出力を実装したパッケージ

func main() {
	print()
}

/*
***** Go 데이터 타입 *****

01. 불린 타입
- bool

02. 문자열 타입
- string: string은 한번 생성되면 수정될 수 없는 Immutable 타입

03. 정수형 타입
- int int8 int16 int32 int64
- uint uint8 uint16 uint32 uint64 uintptr

04. Float 및 복소수 타입
- float32 float64 complex64 complex128

05. 기타 타입
- byte : unit8과 동일하며 바이트 코드에 사용
- rune : int32과 동일하며 유니코드 코드포인트에 사용
*/

/*
***** 문자열 *****

문자열 리터럴은 Back Quote(``) 혹은 이중인용부호("")를 사용하여 표현 가능.

01. 백틱 (``) = バッククォート 로 둘러싸인 문자열은 Raw String Literal 라고 불린다.
	즉, Raw String 값 그대로를 가진다. (\n 같은거 전부 다 그대로 출력)

02. 이중인용부호 ("") 로 둘러싸인 문자열은 Interpreted String Literal 라고 불린다.
	복수 라인에 걸쳐 쓸 수 없으며, 인용부호 안의 Escape 문자열들은 특별한 의미로 해석된다.
	다음 라인으로 내려서 출력하고 싶으면 \n 를 이용하여 New Line 로 해석하게 할 것.
*/

func print() {
	// Raw String Literal.
	rawLiteral := `아리랑\n
아리랑\n
아라리요`

	// Interpreted String Literal
	interLiteral := "아리랑아리랑\n아라리요"
	// 아리와 같이 + 를 사용하여 두 라인에 걸쳐 사용할 수도 있다.
	// interLiteral := "아리랑아리랑\n" +
	//				   "아라리요"

	fmt.Println(rawLiteral)
	fmt.Println()
	fmt.Println(interLiteral)
}

/*
***** 출력구문 *****

아리랑\n
아리랑\n
  아리리요

아리랑아리랑
아리리요

*/
