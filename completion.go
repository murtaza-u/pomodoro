package main

import (
	"fmt"
	"os"
)

var commands = []string{"start", "add", "stop", "help"}

func comp() {
	for _, c := range commands {
		if c[0:len(os.Args[2])] == os.Args[2] {
			fmt.Println(c)
		}
	}
}
