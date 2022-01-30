package main

import "time"

func isRunning(s State) (time.Time, bool) {
	if s.Running != true || len(s.EndsAt) == 0 {
		return time.Time{}, false
	}

	endt, err := time.Parse(time.RFC3339, s.EndsAt)
	if err != nil {
		return time.Time{}, false
	}

	return endt, true
}

func start() {

}

func stop() {

}

func print() {

}
