package config

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	HOTKET_INDEX                   = 0
	MANY_TIME_INDEX                = 1
	INTERVAL_BETWEEN_HOTKEYS_INDEX = 2
	SLEEP_TIME_INDEX               = 3
)

type Macro struct {
	Hotkey                 []string
	ManyTimes              int
	SleepTime              int
	IntervalBetweenHotkeys int
}

func LoadMacroConfig(configPath string) []Macro {
	macros := make([]Macro, 0)
	f, err := os.Open(configPath)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		macros = append(macros, ParseMacroConfig(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return macros
}

func ParseMacroConfig(config string) Macro {
	splittedConfig := strings.Split(config, ";")

	hotkeys := strings.Split(splittedConfig[HOTKET_INDEX], " ")
	manyTimes, _ := strconv.Atoi(splittedConfig[MANY_TIME_INDEX])
	sleepTime, _ := strconv.Atoi(splittedConfig[SLEEP_TIME_INDEX])
	intervalBetweenHotkeys, _ := strconv.Atoi(splittedConfig[INTERVAL_BETWEEN_HOTKEYS_INDEX])

	return Macro{
		IntervalBetweenHotkeys: int(intervalBetweenHotkeys),
		SleepTime:              int(sleepTime),
		ManyTimes:              int(manyTimes),
		Hotkey:                 hotkeys,
	}
}
