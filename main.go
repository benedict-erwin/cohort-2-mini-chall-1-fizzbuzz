package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) > 1 {
		val := os.Args[1]
		max, err := strconv.Atoi(val)
		if err != nil {
			fmt.Println("Input must be integer!")
		}
		fmt.Println("Using N =", max)
		modOfThreeOrFive(max)
	} else {
		fmt.Println("Using N = 15")
		modOfThreeOrFive(15)
	}
}

func modOfThreeOrFive(max int) {
	fmt.Println("=====================")
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
	fmt.Println("=====================")
}
