package main

import (
	"fmt"
	"time"
)

func main() {

	const maxGen int = 100
	const x, y = 30, 30
	
	NewWorld(x, y)
	defer world.Stop()

	// key event loop start
	quit := make(chan struct{})
	go inputLoop(quit)

	fin := false

	for i := 0; i < maxGen; i++ {
		time.Sleep(time.Millisecond * 100)
		world.Update()

		select {
			// clicking Enter Key 
			case <-quit:
				fmt.Println("done")
				fin = true
			default:
				continue
		}

		if fin {
			break
		}
	}

}