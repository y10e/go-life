package main

import (
	"time"
)

func main() {

	maxGen := 5
	x, y := 50, 50
	
	NewWorld(x, y)
	for i := 0; i < maxGen; i++ {
		time.Sleep(time.Second * 2)
		world.Update()
	}
	world.Stop()

}