package main

import (
	"fmt"
	"os"
)

func run() {
	if len(os.Args) == 1 {
		print()
		return
	}

	switch os.Args[1] {
	case "start":
		start()
	case "stop":
		stop()
	default:
		fmt.Println("Invalid argument")
	}
}
