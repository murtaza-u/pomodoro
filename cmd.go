package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"
)

func help() {
	fmt.Printf(
		`Usage: %s [start|stop|add|help]

Run %s without any options to get status of the current session

start
  - Starts pomodoro session(default duration: 25m)
  - Takes duration(in minutes) of pomodoro session as an optional argument

stop
  - Stop pomodoro session

add [duration(in min)]
  - Adds additional time to the pomodoro session
  - Takes duration(in minutes) as a compulsory argument

help
  - Displays this help information

Add 'complete -C pom pom' in your .bashrc to enable completion%s`,
		os.Args[0],
		os.Args[0],
		"\n",
	)
}

func getDuration(optional bool) (int, error) {
	dur := 25
	if len(os.Args) >= 3 {
		var err error
		dur, err = strconv.Atoi(os.Args[2])
		if err != nil {
			return 0, err
		}
	} else if !optional {
		return 0, errors.New("Duration not provided")
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
		dur, err := getDuration(true)
		if err != nil {
			fmt.Println(err)
			return
		}
		start(time.Duration(dur * 60 * 1e9))

	case "add":
		dur, err := getDuration(false)
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
