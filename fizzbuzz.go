package gofizzbuzz

import "strconv"

type FizzBuzz int

func (fb FizzBuzz) Say() string {
	switch {
	case isFizzBuzz(fb):
		return "FizzBuzz"
	case isFizz(fb):
		return "Fizz"
	case isBuzz(fb):
		return "Buzz"
	default:
		return strconv.Itoa(int(fb))
	}
}

func isFizz(fb FizzBuzz) bool {
	return fb%3 == 0
}

func isBuzz(fb FizzBuzz) bool {
	return fb%5 == 0
}

func isFizzBuzz(fb FizzBuzz) bool {
	return isFizz(fb) && isBuzz(fb)
}
