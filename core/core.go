package core

import (
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
	if len(hotkeys) > 1 {
		robotgo.KeyTap(hotkeys[0], hotkeys[1:])
	} else {
		robotgo.KeyTap(hotkeys[0])
	}
}

func runForever() {
	for {
		time.Sleep(time.Hour)
	}
}
