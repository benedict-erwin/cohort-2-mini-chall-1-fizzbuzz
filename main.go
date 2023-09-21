package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var num int
	if len(os.Args) > 1 {
		val := os.Args[1]
		max, err := strconv.Atoi(val)
		if err != nil {
			fmt.Println("Input must be integer!")
		}
		num = max
	} else {
		num = 15
	}
	modOfThreeOrFive(num)
}

func modOfThreeOrFive(max int) {
	fmt.Println("Using N =", max)
	fmt.Println("==========")
	for i := 1; i <= max; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
	fmt.Println("==========")
}
