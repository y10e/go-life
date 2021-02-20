package main

import (
	"time"
)

func main() {

	const maxGen int = 100
	const x, y = 30, 30
	
	NewWorld(x, y)
	defer world.Stop()

	// key event loop start
	event := make(chan Event)
	go inputLoop(event)
	defer close(event)

	fin := false
	for i := 0; i < maxGen; i++ {
		time.Sleep(time.Millisecond * 100)
		world.Update()

		select {
			// clicking Enter Key 
			case keyin := <-event:
				//fmt.Println(keyin)

				if keyin.Type == "fin" {
					fin = true
				}else if keyin.Type == "pause" {
					
					//fin = true
					pauseLoop(event)
					keyin.Type = "restart"
				}
			default:
				continue
		}

		// Complete this game
		if fin {
			break
		}
	}

}

func pauseLoop(event chan Event){
	world.Pause()
	restart := false
	for {		
		select {
			// clicking Enter Key 
			case keyin := <-event:
				if keyin.Type == "pause" {
					restart = true
					world.Restart()
				}
			default:
				continue
		}

		if restart {
			break
		}
		time.Sleep(time.Millisecond * 250)
	}
}