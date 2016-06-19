package fizzbuzz

import "testing"

const Fizz = "Fizz"
const Buzz = "Buzz"

func TestFizzBuzzTo3(t *testing.T)  {
	expected := "12"+Fizz
	actual := FizzBuzz(3)

	if expected != actual {
		t.Error("Error! Expected %s", expected, " but received %s", actual)
	}
}

func TestFizzBuzzTo5(t *testing.T)  {
	expected := "12"+Fizz+"4"+Buzz
	actual := FizzBuzz(5)

	if expected != actual {
		t.Error("Error! Expected %s", expected, " but received %s", actual)
	}
}

func TestFizzBuzzTo15(t *testing.T)  {
	expected := "12"+Fizz+"4"+Buzz+Fizz+"78"+Fizz+Buzz+"11"+Fizz+"1314"+Fizz+Buzz
	actual := FizzBuzz(15)

	if expected != actual {
		t.Error("Error! Expected %s", expected, " but received %s", actual)
	}
}