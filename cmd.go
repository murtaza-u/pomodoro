package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func help() {
	fmt.Printf(
		`Usage: %s [start|stop]

Run %s without any options to get status of the current session

start
  - Starts pomodoro session(default duration: 25m)
  - Takes duration of pomodoro session as an optional argument

stop
  - Stop pomodoro session%s`,
		os.Args[0],
		os.Args[0],
		"\n",
	)
}

func getDuration() (int, error) {
	dur := 25
	if len(os.Args) >= 3 {
		var err error
		dur, err = strconv.Atoi(os.Args[2])
		if err != nil {
			return 0, nil
		}
	}
	return dur, nil
}

func run() {
	if len(os.Args) == 1 {
		print()
		return
	}

	switch os.Args[1] {
	case "start":
		dur, err := getDuration()
		if err != nil {
			fmt.Println(err)
			return
		}
		start(time.Duration(dur * 60 * 1e9))

	case "add":
		dur, err := getDuration()
		if err != nil {
			fmt.Println(err)
			return
		}
		add(time.Duration(dur * 60 * 1e9))

	case "stop":
		stop()

	case "help":
		help()

	default:
		fmt.Println("Invalid argument")
	}
}
