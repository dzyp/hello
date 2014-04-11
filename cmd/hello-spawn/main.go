package main

import (
	_ "circuit/load/cmd"  // Link the circuit into this executable
	"circuit/use/circuit" // Import the circuit language API
	"hello/x"             // Import the package implementing the worker function
	"log"
)

func main() {
	println("Starting...")
	println(`TEST1`)

	retrn, _, err := circuit.Spawn(
		"23.251.153.186",   // Host where to spawn the new worker
		[]string{"/hello"}, // List of anchors for the new worker
		x.App{},            // User type that encloses the worker function
		"agreed!",          // Argument to pass to worker function
	)

	println(`ABOUT TO SLEEP`)
	if err != nil {
		println("Oh well", err.Error())
		return
	}

	tbl := retrn[0].(circuit.XPerm)

	results := tbl.Call(`Bounce`)

	log.Println(results[0])

	results = tbl.Call(`Bounce`)

	log.Println(results[0])

	println(`we are done`)
}
