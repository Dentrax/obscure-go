package main

import (
	obs "github.com/Dentrax/obscure-go/observer"
	secure "github.com/Dentrax/obscure-go/types"
)

func main() {
	ExampleInt()
	ExampleString()
}

func ExampleInt()  {
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

func ExampleString()  {
	w := obs.CreateWatcher("watcher")
	lhs := secure.NewString("foo")
	lhs.AddWatcher(w)

	lhs.Set("foo 2")

	println("LHS: ", lhs.Get())

	rhs := secure.NewString("bar")
	rhs.Set("foo 2")

	println("RHS: ", rhs.Get())

	println("LHS == RHS: ", lhs.IsEquals(rhs))
}
