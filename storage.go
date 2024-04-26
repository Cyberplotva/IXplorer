package main

var storage = struct {
	cursorPosition map[string]int // Map dir path to cursor position on its entries
}{
	cursorPosition: make(map[string]int),
}
