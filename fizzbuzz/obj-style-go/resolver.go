package main

import "strconv"

type Resolver struct {
	strategys []Strategy
}

func NewResolver(strategys ...Strategy) *Resolver {
	return &Resolver{
		strategys: strategys,
	}
}

func (r *Resolver) Resolve(n int) string {
	for _, strategy := range r.strategys {
		if resolve := strategy.Apply(n); resolve != "" {
			return resolve
		}
	}

	return r.NoMatch(n)
}

func (r Resolver) NoMatch(n int) string {
	return strconv.Itoa(n)
}
