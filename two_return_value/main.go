package main

import "fmt"

//매개변수가 여러 개일 때에는 `.`로 구분하여 작성한다.
// 같은 타입의 매개변수일 경우 마지막에만 타입을 표기한다
/*
func thisfunc(b bool, s stinrg, i,j,k int) (int, int {
	. . .
})
*/
/*
func 함수명(매개변수) (리턴타입 & 리턴값){
	. . .
}

*/

func multiply(w, h int) (int, int) {
	return w * 2, h * 2
}

func main() {
	w, h := multiply(3, 4)
	fmt.Println(w, h)
}
