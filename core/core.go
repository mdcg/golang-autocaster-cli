package core

import (
	"log"
	"time"

	"github.com/go-vgo/robotgo"
	"github.com/mdcg/golang-autocaster-cli/config"
)

func LoadMacros(macros []config.Macro) {
	for _, m := range macros {
		go runMacro(m)
	}

	runForever()
}

func runMacro(macro config.Macro) {
	for {
		robotgo.Sleep(macro.SleepTime)
		for t := 0; t < macro.ManyTimes; t++ {
			execKeyCombination(macro.Hotkey)
			robotgo.Sleep(1)
		}
	}
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

func runForever() {
	for {
		time.Sleep(time.Hour)
	}
}
