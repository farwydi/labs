package main

type fizzBuzzStrategy struct {
}

func NewFizzBuzzStrategy() Strategy {
	return &fizzBuzzStrategy{}
}

func (f fizzBuzzStrategy) Apply(n int) string {
	if n%3 == 0 && n%5 == 0 {
		return "FizzBuzz"
	}

	return ""
}
