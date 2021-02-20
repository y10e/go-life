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

	isFin := false
	isPause := false
	for i := 0; i < maxGen; i++ {

		//Interval & world update
		time.Sleep(time.Millisecond * 100)
		world.Update()

		//pause mode
		if isPause {
			pauseLoop(event)	
			world.Restart()
			isPause = false
		}

		// Complete this game
		if isFin {
			break
		}

		select {
			case keyin := <-event:

				if keyin.Type == "fin" {
					// clicking Ctrl + C
					isFin = true
				}else if keyin.Type == "pause" {
					// clicking Enter Key
					world.Pause()
					isPause = true
					keyin.Type = "restart"
				}
			default:
				continue
		}
	}

}

// Lopping unitl clicking EnterKey for Pause Mode
func pauseLoop(event chan Event){
	
	restart := false
	for {		
		select {
			case keyin := <-event:
				if keyin.Type == "pause" {
					// clicking Enter Key, cancel pause mode
					restart = true
				}
			default:
				continue
		}

		if restart {
			break
		}

		time.Sleep(time.Millisecond * 100)
	}
}