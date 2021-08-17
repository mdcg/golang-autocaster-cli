package core

import (
	"log"
	"math/rand"
	"time"

	"github.com/go-vgo/robotgo"
	"github.com/mdcg/golang-autocaster-cli/config"
)

const (
	MAX_SLEEP_TIME = 5
)

func runForever() {
	for {
		time.Sleep(time.Hour)
	}
}

func randomSleep(maxInterval int) {
	rand.Seed(time.Now().UnixNano())
	parsedMaxInterval := rand.Intn(maxInterval)
	config.InfoLogger.Printf("Sleeping %d seconds...\n", parsedMaxInterval)
	time.Sleep(time.Duration(parsedMaxInterval) * time.Second)
}

func execKeyCombination(hotkeys []string) {
	var err string

	if len(hotkeys) > 1 {
		err = robotgo.KeyTap(hotkeys[0], hotkeys[1:])
	} else {
		err = robotgo.KeyTap(hotkeys[0])
	}

	if err != "" {
		log.Fatal(err)
	}
}

func LoadMacros(macros []config.Macro) {
	for _, m := range macros {
		randomSleep(MAX_SLEEP_TIME)
		go runMacro(m)
	}
	runForever()
}

func runMacro(macro config.Macro) {
	for {
		for t := 0; t < macro.ManyTimes; t++ {
			robotgo.Sleep(macro.IntervalBetweenHotkeys)
			config.InfoLogger.Printf("Hotkey: %v\n", macro.Hotkey)
			execKeyCombination(macro.Hotkey)
		}
		config.InfoLogger.Printf("Sleeping for %v seconds.\n", macro.SleepTime)
		robotgo.Sleep(macro.SleepTime)
	}
}
