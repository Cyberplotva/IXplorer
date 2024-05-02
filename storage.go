package main

import "sync"

var storage = struct {
	cursorPosition sync.Map // Map dir path to cursor position on its entries, string -> int
}{}
