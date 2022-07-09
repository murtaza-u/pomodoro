package main

import (
	"fmt"
	"os"
)

var cmds = []string{"start", "add", "stop", "help"}

func comp() {
	if os.Args[1] != os.Args[3] {
		return
	}

	for _, c := range cmds {
		if c[0:len(os.Args[2])] == os.Args[2] {
			fmt.Println(c)
		}
	}
}
