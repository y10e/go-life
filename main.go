package main

import (
	"fmt"
	"time"

	"github.com/gdamore/tcell"
)

func main() {

	const maxGen int = 10
	const x, y = 30, 30
	
	NewWorld(x, y)
	defer world.Stop()

	quit := make(chan struct{})
	go func(){
		ev := screen.PollEvent()
		switch ev := ev.(type) {
			case *tcell.EventKey:
				if ev.Key() == tcell.KeyEnter {
					fmt.Println("enter")
					close(quit)
					break
				}
		}
	}()

	<-quit

	for i := 0; i < maxGen; i++ {
		time.Sleep(time.Millisecond * 100)
		world.Update()
		
	}
	//<-quit
	//world.Stop()

}