package main

import "fmt"

func main() {
	resolver := NewResolver(NewFizzBuzzStrategy(), NewFizzStrategy(), NewBuzzStrategy())
	for i := 1; i <= 100; i++ {
		fmt.Println(resolver.Resolve(i))
	}
}
