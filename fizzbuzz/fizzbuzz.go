package fizzbuzz

import (
	"bytes"
	"strconv"
)

func FizzBuzz(max int) string {
	var result bytes.Buffer

	for num := 1; num <= max; num++ {
		if num%3 == 0 {
			result.WriteString("Fizz")
		} else {
			result.WriteString(strconv.Itoa(num))
		}
	}
	return result.String()
}