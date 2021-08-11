package config

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Macro struct {
	Hotkey    []string
	SleepTime int
	ManyTimes int
}

func LoadConfig(configPath string) []Macro {
	macros := make([]Macro, 0)
	f, err := os.Open(configPath)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		macros = append(macros, ParseConfig(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return macros
}

func ParseConfig(config string) Macro {
	splittedConfig := strings.Split(config, ";")

	hotkeys := strings.Split(splittedConfig[0], " ")
	sleepTime, _ := strconv.ParseInt(splittedConfig[1], 10, 32)
	manyTimes, _ := strconv.ParseInt(splittedConfig[2], 10, 32)

	return Macro{
		Hotkey:    hotkeys,
		SleepTime: int(sleepTime),
		ManyTimes: int(manyTimes),
	}
}
