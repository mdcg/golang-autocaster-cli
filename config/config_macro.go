package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

const (
	HOTKET_INDEX                   = 0
	MANY_TIME_INDEX                = 1
	INTERVAL_BETWEEN_HOTKEYS_INDEX = 2
	SLEEP_TIME_INDEX               = 3
)

type Macro struct {
	Hotkey                 []string `json:"hotkey"`
	ManyTimes              int      `json:"many_times"`
	SleepTime              int      `json:"sleep_time"`
	IntervalBetweenHotkeys int      `json:"interval_between_hotkeys"`
}

func LoadMacroConfig(configPath string) []Macro {
	var macros []Macro

	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		fmt.Print(err)
	}

	err = json.Unmarshal(data, &macros)
	if err != nil {
		fmt.Print(err)
	}

	return macros
}
