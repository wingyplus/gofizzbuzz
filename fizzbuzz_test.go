package gofizzbuzz

import "testing"

func TestFizzBuzz(t *testing.T) {
	var tests = []struct {
		number int
		actual string
	}{
		{number: 1, actual: "1"},
		{number: 3, actual: "Fizz"},
		{number: 5, actual: "Buzz"},
		{number: 15, actual: "FizzBuzz"},
	}

	for _, test := range tests {
		var word string = FizzBuzz(test.number).Say()
		if word != test.actual {
			t.Errorf("expect %d but was %s", test.number, word)
		}
	}
}
