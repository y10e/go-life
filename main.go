package main

import (
	"fmt"
	"time"

	"github.com/gdamore/tcell"
)

func main() {

	const maxGen int = 100
	const x, y = 30, 30
	
	NewWorld(x, y)
	defer world.Stop()

	quit := make(chan struct{})

	go func(){
		for {
			ev := screen.PollEvent()
			switch ev := ev.(type) {
				case *tcell.EventKey:
					if ev.Key() == tcell.KeyEnter {
						close(quit)
					}
			}
		}
	}()

	fin := false

	for i := 0; i < maxGen; i++ {
		time.Sleep(time.Millisecond * 100)
		world.Update()

		select {
			case <-quit:
				fmt.Println("done")
				//time.Sleep(time.Millisecond * 5000)
				fin = true
			default:
				//fmt.Println("default")
				//continue
		}

		if fin {
			break
			//world.Stop()
		}

		
		
	}
	//<-quit
	//world.Stop()

}