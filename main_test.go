package main

import "testing"

func TestMainFunc(t *testing.T) {
	go main() // run in goroutine so it doesn’t block
}
