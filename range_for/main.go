package main

import "fmt"

func main() {
	fruits := [10]string{"Apple", "Banana", "Coconut", "Durian", "Grape", "Pineapple", "Watermelon", "Lemon", "Mango", "Peach"}
	for i, s := range fruits {
		// i 는 문자열 배열의 인덱스를 뜻한다
		// s 는 ㄴ인덱스에 대응하는 문자열의 값을 뜻한다
		fmt.Printf("fruits[%d]=%s\n", i, s)
	}
}
