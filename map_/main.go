package main

import "fmt"

func main() {
	/*int타입의 키와 string 타입의 값을 가지는 맵*/
	// var m map[int]string

	/*맵도 참조형이므로 make함수를 이용하여 생성할 수 있다. */
	a := make(map[int]string)
	a[82] = "Korea"
	a[62] = "Indonesia"
	a[1] = "United States"
	a[86] = "China"
	a[81] = "Japan"
	a[7] = "Russian Federation"

	fmt.Println(a)
	// 출력값 : map[1:United States 7:Russian Federation 62:Indonesia 81:Japan 82:Korea 86:China]
}
