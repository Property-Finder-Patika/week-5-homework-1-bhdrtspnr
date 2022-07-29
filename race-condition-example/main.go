package main

import (
	"fmt"
	"time"
)

func main() {
	numbers := []int{}
	go appendOne(&numbers)
	go appendTwo(&numbers)
	time.Sleep(30 * time.Millisecond)
}

func appendTwo(numbers *[]int) {
	for i := 0; i <= 10; i++ {
		*numbers = append(*numbers, 2)
		fmt.Printf("numbers %v\n", numbers)
	}
}

func appendOne(numbers *[]int) {
	for i := 0; i <= 10; i++ {
		*numbers = append(*numbers, 1)
		fmt.Printf("numbers %v\n", numbers)
	}
}
