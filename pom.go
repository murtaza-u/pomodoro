package main

import (
	"fmt"
	"log"
	"time"
)

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

func start(dur time.Duration) {
	write := func(s State) {
		s.Running = true
		s.EndsAt = time.Now().Add(dur).Format(time.RFC3339)
		if err := s.write(); err != nil {
			log.Panic(err)
		}
	}

	var s State

	if err := s.read(); err != nil {
		write(s)
		return
	}

	if _, running := isRunning(s); running {
		fmt.Println("Pomodoro already running")
		return
	}

	write(s)
}

func stop() {
	var s State
	if err := s.read(); err != nil {
		fmt.Println("Pomodoro not running")
		return
	}

	if _, running := isRunning(s); !running {
		fmt.Println("Pomodoro not running")
		return
	}

	s = State{
		Running: false,
		EndsAt:  "",
	}

	if err := s.write(); err != nil {
		log.Panic(err)
	}
}

func print() {
	var s State
	if err := s.read(); err != nil {
		return
	}

	sTime, running := isRunning(s)
	if !running {
		return
	}

	dur := sTime.Sub(time.Now()).Round(time.Second)

	if dur < 0 {
		fmt.Printf("💀%v\n", dur)
		return
	}

	fmt.Printf("🍅%v\n", dur)
}
