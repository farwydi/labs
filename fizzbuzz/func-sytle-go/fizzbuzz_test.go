package main

import (
	"strconv"
	"testing"
)

func Test_fizzbuzz(t *testing.T) {
	tests := []struct {
		n    int
		want string
	}{
		{
			n:    1,
			want: "1",
		},
		{
			n:    3,
			want: "Fizz",
		},
		{
			n:    5,
			want: "Buzz",
		},
		{
			n:    7,
			want: "7",
		},
		{
			n:    15,
			want: "FizzBuzz",
		},
		{
			n:    30,
			want: "FizzBuzz",
		},
		{
			n:    33,
			want: "Fizz",
		},
	}
	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.n), func(t *testing.T) {
			if got := fizzbuzz(tt.n); got != tt.want {
				t.Errorf("fizzbuzz() = %v, want %v", got, tt.want)
			}
		})
	}
}
