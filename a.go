package main

import (
	f "fmt"
)

//RSP is structur
type RSP struct {
	code int
	msg  string
	pld  string
}

var r RSP

func init() {
	W("a inited.")
}

// W rite something
func W(a ...string) {
	f.Println(a)
}

// ShowRSP show some thin
func ShowRSP() {
	f.Println(r)
}

//A is a fucntion
func A() {
	f.Println("A")
}
