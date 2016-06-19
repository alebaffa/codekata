package fizzbuzz

import "testing"

func TestFizzBuzzTo3(t *testing.T)  {
	expected := "12Fizz"
	actual := FizzBuzz(3)

	if expected != actual {
		t.Error("Error! Expected %s", expected, " but received %s", actual)
	}
}

func TestFizzBuzzTo5(t *testing.T)  {
	expected := "12Fizz4Buzz"
	actual := FizzBuzz(5)

	if expected != actual {
		t.Error("Error! Expected %s", expected, " but received %s", actual)
	}
}