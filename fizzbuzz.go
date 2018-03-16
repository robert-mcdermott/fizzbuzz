package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Fprintf(os.Stderr, "\nA number is required\n")
		os.Exit(1)
	}
	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "\nPlease provide an integer only!\n")
		os.Exit(1)
	}
	fmt.Printf("FizzBuzz for %d:\n", num)
	for _, n := range fizzBuzz(num) {
		fmt.Println(n)
	}
}

func fizzBuzz(num int) []string {
	res := []string{}
	for i := 1; i <= num; i++ {
		t := ""
		if i%3 == 0 {
			t += "Fizz"
		}
		if i%5 == 0 {
			t += "Buzz"
		}
		if len(t) == 0 {
			t = strconv.Itoa(i)
		}
		res = append(res, t)
	}
	return res
}
