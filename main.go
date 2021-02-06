package main

import (
	"time"
)

func main() {

	const maxGen int = 1000
	const x, y = 50, 50
	
	NewWorld(x, y)
	for i := 0; i < maxGen; i++ {
		time.Sleep(time.Millisecond * 100)
		world.Update()
	}
	world.Stop()

}