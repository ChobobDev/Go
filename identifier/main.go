package main

import "fmt"

const (
	아침인사 = " 굿 모닝"
	낮인사  = " 굿 애프터눈"
	저녁인사 = "굿 으브닝"
)

func 인삿말(m string) {
	fmt.Println(m)
}

func main() {
	인삿말(저녁인사)
}
