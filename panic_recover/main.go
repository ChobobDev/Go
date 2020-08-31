package main

import "fmt"

func f() {
	defer func() { // recover 함수는 지연 호출로 사용해야 함
		s := recover() // 패닉이 발생해도 프로그램을 종료하지 않음, panic 함수에서 설정한 에러 메시지를 받아옴
		fmt.Println(s)
	}()

	panic("Error") // panic 함수로 에러 메시지 설정, 패닉 발생
}

// panic 은 다음가 같이 정의되어 있음 "func panic(v interface{})""

func main() {
	f()
	fmt.Println("Hello 世界!") //이 부분은 출력되지 않음

}
