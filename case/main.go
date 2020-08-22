package main

import "fmt"

func main() {
	v := '9'
	switch {
	// case에 조건식 사용
	case '0' <= v && v <= '9':
		fmt.Printf("%c은(는) 숫자입니다", v)
	case 'a' <= v && v <= 'z':
		fmt.Printf("%c은(는) 소문자입니다", v)
	case 'A' <= v && v <= 'Z':
		fmt.Printf("%c은(는) 대문자입니다", v)
	}
}
