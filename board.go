package main

//NewWorld :initilize the world
func NewBoard() {
	CreateStartMenu()
}

func (board *Board) StopBoard() {
	view.StopStartMenu()
}