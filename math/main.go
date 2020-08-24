package main

import (
	"fmt"
	"math"
	"math/bits"
)

func main() {
	var mi32 int64 = math.MinInt32
	var mi64 int64 = math.MinInt64

	var i32 uint64 = math.MaxInt32
	var ui32 uint64 = math.MaxUint32
	var i64 uint64 = math.MaxInt64
	var ui64 uint64 = math.MaxUint64
	var ui uint64 = (1 << bits.UintSize) - 1
	var i uint64 = (1<<bits.UintSize)/2 - 1
	var mi int64 = (1 << bits.UintSize) / -2

	fmt.Printf(" MinInt32: %d\n", mi32)
	fmt.Printf(" MaxInt32:  %d\n", i32)
	fmt.Printf("MaxUint32:  %d\n", ui32)
	fmt.Printf(" MinInt64: %d\n", mi64)
	fmt.Printf(" MaxInt64:  %d\n", i64)
	fmt.Printf("MaxUint64:  %d\n", ui64)
	fmt.Printf("  MaxUint:  %d\n", ui)
	fmt.Printf("   MinInt: %d\n", mi)
	fmt.Printf("   MaxInt:  %d\n", i)
}
