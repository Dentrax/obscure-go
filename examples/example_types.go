package main

import (
	secure "github.com/Dentrax/obscure-go/types"
)

func main() {
	lhs := secure.NewInt(15)
	lhs.Inc()
	lhs.Inc()
	lhs.Inc()
	lhs.Dec()

	println("LHS: ", lhs.Get())

	rhs := secure.NewInt(99)
	rhs.Set(18)
	rhs.Dec()

	println("RHS: ", rhs.Get())

	println("LHS == RHS: ", lhs.IsEquals(rhs))
}
