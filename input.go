package main

import "github.com/gdamore/tcell"

func inputLoop(quit chan struct{}) {
	for {
		ev := screen.PollEvent()
		switch ev := ev.(type) {
			case *tcell.EventKey:
				if ev.Key() == tcell.KeyEnter {
					close(quit)
				}
			default :
				continue
		}
	}
}