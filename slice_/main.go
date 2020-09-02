package main

import "fmt"

func main() {
	/* 슬라이스는 Go에서 가장 많이 사용되는 데이터 구조이다*/
	/* int타입의 슬라이스 */
	var a []int
	b := make([]int, 10)
	var c [10]int

	fmt.Println(a)
	/* 출력결과*/
	/* [] */
	fmt.Println(b)
	/* 출력결과*/
	/* [0 0 0 0 0 0 0 0 0 0] */
	fmt.Println(c)
	/* 출력결과*/
	/* [0 0 0 0 0 0 0 0 0 0] */

	/* float타입의 슬라이스 */
	d := make([]float64, 3)
	fmt.Println(d)
	/* 출력결과*/
	/* [0 0 0] */
	d[0] = 0.25
	fmt.Println(d)
	/* 출력결과*/
	/* [0.25 0 0] */
	d[1] = 0.75
	fmt.Println(d)
	/* 출력결과*/
	/* [0.25 0.75 0] */
	d[2] = 0.125
	fmt.Println(d)
	/* 출력결과*/
	/* [0.25 0.75 0.125] */

}
