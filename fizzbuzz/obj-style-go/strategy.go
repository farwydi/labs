package main

type Strategy interface {
	Apply(n int) string
}
