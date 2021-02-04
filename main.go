package main

import (
	"time"
)

func main() {

	const maxGen int = 3
	const x, y = 50, 50
	
	NewWorld(x, y)
	for i := 0; i < maxGen; i++ {
		time.Sleep(time.Second * 4)
		world.Update()
	}
	world.Stop()

}