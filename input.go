package main

import (
	"github.com/gdamore/tcell"
)

func inputLoop(event chan Event) {
	for {
		ev := screen.PollEvent()
		switch ev := ev.(type) {
			case *tcell.EventKey:
				if ev.Key() == tcell.KeyEnter {
					//close(event)
					//TODO:Game Start
					event <- Event{Type:"start"}
					
				}else if ev.Key() == tcell.KeyCtrlC {
					//Stop Game
					event <- Event{Type:"fin"}
				}else if ev.Rune() == ' ' {
					//Pause, Restart
					event <- Event{Type:"pause"}
				}

			default :
				continue
		}
	}
}