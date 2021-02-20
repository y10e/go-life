package main

import (
	"fmt"

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

func pauseInputLoop(event chan bool) {
	fmt.Println("pauseInputLoop Start")
	for {
		ev := screen.PollEvent()
		switch ev := ev.(type) {
			case *tcell.EventKey:
				fmt.Println("pauseInputLoop Finish")
				fmt.Println(ev.Key())
				event <- true
			default :
				continue
		}
	}
}