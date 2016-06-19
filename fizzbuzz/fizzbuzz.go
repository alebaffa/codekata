package fizzbuzz

import (
	"bytes"
	"strconv"
)

func FizzBuzz(max int) string {
	var result bytes.Buffer

	for num := 1; num <= max; num++ {
		switch {
		case isDivisibleBy3And5(num):
			result.WriteString("FizzBuzz")
		case isDivisibleBy3(num):
			result.WriteString("Fizz")
		case isDivisibleBy5(num):
			result.WriteString("Buzz")
		default:
			result.WriteString(strconv.Itoa(num))
		}
	}
	return result.String()
}

func isDivisibleBy3And5(num int) bool {
	return num%3 == 0 && num%5 == 0
}

func isDivisibleBy3(num int) bool {
	return num%3 == 0
}

func isDivisibleBy5(num int) bool {
	return num%5 == 0
}