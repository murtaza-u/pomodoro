package main

import "os"

func main() {
	if len(os.Getenv("COMP_LINE")) != 0 {
		comp()
		return
	}

	run()
}
