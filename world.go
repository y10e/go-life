package main

import (
	"math/rand"
	"time"
)

//NewWorld :initilize the world
func NewWorld(x,y int) {
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
			if r < 0.3 {
				w[i][j] = true
			}
		}
	} 
	//fmt.Println(w)
	world = &World{0,x,y,w}
	NewView()

}

//Stop : Stop the world
func (world *World) Stop() {
	view.Stop()
}

//Update : update the world
func (world *World) Update() {

	for i:=0; i < len(world.w); i++ {
		for j := 0; j < len(world.w[i]); j++ {
			r := rand.Float32()
			if r < 1.0 {
				world.w[i][j] = !(world.w[i][j])
			}
		}
	} 
	view.Update()
}

