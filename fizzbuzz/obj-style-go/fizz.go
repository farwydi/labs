package main

type fizzStrategy struct {
}

func NewFizzStrategy() Strategy {
	return &fizzStrategy{}
}

func (f fizzStrategy) Apply(n int) string {
	if n%3 == 0 {
		return "Fizz"
	}

	return ""
}
