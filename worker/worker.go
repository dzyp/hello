package main

import (
	// Package worker ensures that this executable will act as a circuit worker
	_ "circuit/load/worker"

	// Importing hello/x ensures that the hello tutorial's logic is linked into the worker executable
	_ "hello/x"
)

// Main will never be executed.
func main() {}
