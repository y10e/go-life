package main

import (
	"math/rand"
	"time"

	"github.com/gdamore/tcell"
)

//NewWorld :initilize the world
func NewWorld(x,y int, c string) {
	//fmt.Println("Call NewWorld")

	w := make([][]bool, x)

	for i:=0; i<x; i++{
		w[i] = make([]bool, y)
	}

	//fmt.Printf("x:%d,",len(w))
	//fmt.Printf("y:%d¥n",len(w[0]))

	rand.Seed(time.Now().UnixNano())
	for i:=0; i < len(w); i++ {
		for j := 0; j < len(w[i]); j++ {
			r := rand.Float32()
			if r < 0.1 {
				w[i][j] = true
			}
		}
	} 
	world = &World{0,x,y,w,running, tcell.GetColor(c)}
	NewView()

}

//Stop : Stop the world
func (world *World) Stop() {
	world.status = stop
	view.Stop()
}

//Pause : Pause the world
func (world *World) Pause() {
	world.status = pause
}

//Restart : Restart the world
func (world *World) Restart() {
	world.status = running
}

//Update : update the world
func (world *World) Update() {
	//fmt.Println(world.w)
	nw := world.w
	for i:=0; i < len(world.w); i++ {
		for j := 0; j < len(world.w[i]); j++ {
			nw[i][j] = judge(i,j)
			//nw[i][j] = !(world.w[i][j])
		}
	} 
	world.w = nw
	world.generaition++
	view.Update()
}

func judge(i int, j int) bool{

	result := false
	c := 0

	for xOffset := -1; xOffset < 2; xOffset++ {
		for yOffset := -1; yOffset < 2; yOffset++ {

			x := i + xOffset
			y := j + yOffset

			if x < 0 {
				x = x + world.x
			}else if x >= world.x {
				x = x - world.x
			}

			if y < 0 {
				y = y + world.y
			}else if y >= world.y {
				y = y - world.y
			}

			//fmt.Printf("(x:%d,y:%d),%t\n",x,y,world.w[x][y])

			if !(xOffset == 0 && yOffset == 0) {
				if world.w[x][y] {
					c++
					//fmt.Printf("c:%d\n",c)
				}
			}
			
		}
	} 
	//fmt.Printf("(i:%d,j:%d),%d)\n",i,j,c)
	
	if world.w[i][j] {
		//live
		if c == 2 || c == 3 {
			result = true
		}else{
			result = false
		}

	} else {
		//dead
		if c == 3 {
			result = true
		}else{
			result = false
		}
	}
	
	//fmt.Printf("result:%t\n",result)
	return result
}