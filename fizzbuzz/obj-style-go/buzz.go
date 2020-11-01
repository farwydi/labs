package main

type buzzStrategy struct {
}

func NewBuzzStrategy() Strategy {
	return &buzzStrategy{}
}

func (f buzzStrategy) Apply(n int) string {
	if n%5 == 0 {
		return "Buzz"
	}

	return ""
}
