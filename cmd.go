package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func run() {
	if len(os.Args) == 1 {
		print()
		return
	}

	switch os.Args[1] {
	case "start":
		dur := 25
		if len(os.Args) >= 3 {
			var err error
			dur, err = strconv.Atoi(os.Args[2])
			if err != nil {
				fmt.Println(err)
				return
			}
		}
		start(time.Duration(dur * 60 * 1e9))

	case "stop":
		stop()

	default:
		fmt.Println("Invalid argument")
	}
}
