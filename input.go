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
					event <- Event{Type:"pause"}
					
				}else if ev.Key() == tcell.KeyCtrlC {
					event <- Event{Type:"fin"}
				}
			default :
				continue
		}
	}
}