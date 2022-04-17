package main

/*
    주의점..!

	Go에서 타입간 변환은 명시적으로 지정해 주어야 한다는 점
	예를들어, 정수형 int 에서 unit 로 변환할 때, 암묵적(implicit) 변환이 일어나지 않도록
	unit(i)처럼 반드시 변환을 지정해줘야 한다.

	만약, 명시적 지정이 없이 변환이 일어난다면 런타임 에러가 발생함.
*/

func main() {
	var i int = 100
	var u uint = uint(i)
	var f float32 = float32(i)

	println(f, u)

	str := "ABC"
	bytes := []byte(str)
	str2 := string(bytes)

	println(bytes, str2)
}
