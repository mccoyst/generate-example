package main

//go:generate stringer -type=Frob
type Frob int

const (
	Meow Frob = iota
	Woof
	Ribbit
)
