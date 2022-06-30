package main

import (
	"encoding/json"
	"io/ioutil"
)

type State struct {
	Running bool   `json:"running"`
	EndsAt  string `json:"endsAt"`
}

const path = "/tmp/pomodoro.json"

func (s *State) read() error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, s)
}

func (s *State) write() error {
	data, err := json.Marshal(s)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(path, data, 0600)
}
