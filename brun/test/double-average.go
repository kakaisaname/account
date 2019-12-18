package main

import (
	"fmt"
	"github.com/kakaisaname/infra/algo"
)

func main() {

	count, amount := int64(10), int64(10000)
	remain := amount
	sum := int64(0)
	for i := int64(0); i < 100; i++ {
		x := algo.DoubleAverage(count-i, remain)
		remain -= x
		sum += x
		fmt.Print(i+1, "=", float64(x)/float64(100), ", ")
	}
	fmt.Println()
	fmt.Println("总和", sum)
}
