package main

import (
	"strconv"

	"github.com/gdamore/tcell"
	"github.com/mattn/go-runewidth"
)

// CreateStartMenu creates a new game screen
func CreateStartMenu() {
	var err error
	screen, err = tcell.NewScreen()
	if err != nil {
		logger.Fatal("NewScreen error:", err)
	}
	err = screen.Init()
	if err != nil {
		logger.Fatal("screen Init error:", err)
	}

	screen.Clear()

	xd := runewidth.RuneWidth('　')
	for x := 0; x < 30; x++ {
		for y := 0; y < 30; y++ {
			styleBoarder := tcell.StyleDefault.Foreground(tcell.ColorBlack).Background(tcell.ColorBlack)
			screen.SetContent(x*xd, y, '　' , nil , styleBoarder)
		}
	}

	//drawText(3*xd, 11 ,"Game Start",tcell.ColorWhite,tcell.ColorBlack)
	drawText(9*xd, 10 ,"Conway's Life Game",tcell.ColorWhite,tcell.ColorBlack)
	drawText(9*xd, 15 ,"  Press EnterKey  ",tcell.ColorWhite,tcell.ColorBlack)
	drawText(20*xd, 27 ,"       version 0.1",tcell.ColorWhite,tcell.ColorBlack)
	drawText(20*xd, 28 ,"Copyright (C) y10e",tcell.ColorWhite,tcell.ColorBlack)

	screen.Show()
	view = &View{}
}

// StopStartMenu stops Start Menu
func (view *View) StopStartMenu(){
	screen.Fini()
}


// NewView creates a new view
func NewView() {
	//fmt.Print("Call NewView")
	var err error

	screen, err = tcell.NewScreen()
	if err != nil {
		logger.Fatal("NewScreen error:", err)
	}
	err = screen.Init()
	if err != nil {
		logger.Fatal("screen Init error:", err)
	}

	screen.Clear()
	//var cell []rune = []rune{'a'}
	xd := runewidth.RuneWidth('　')
	for x := 0; x < world.x; x++ {
		for y := 0; y < world.y; y++ {
			styleBoarder := tcell.StyleDefault.Foreground(tcell.ColorBlack).Background(tcell.ColorBlack)
			if world.w[x][y] {
				styleBoarder = tcell.StyleDefault.Foreground(tcell.ColorBlack).Background(tcell.ColorSpringGreen)
			}
			screen.SetContent(x*xd, y, '　' , nil , styleBoarder)
			//screen.SetContent(x + xOffset + 1, y, ' ', nil, styleBoarder)
		}
		//xOffset = xOffset + 2
	}
	
	drawText(0, world.y,"Gen:" + strconv.Itoa(world.generaition),tcell.ColorWhite,tcell.ColorBlack)
	drawText(0, world.y+1,"Status:start",tcell.ColorWhite,tcell.ColorBlack)
	screen.Show()
	view = &View{}
}

// Update Update the view
func (view *View) Update() {
	screen.Clear()
	//var cell []rune = []rune{'a'}
	xd := runewidth.RuneWidth('　')
	for x := 0; x < world.x; x ++{
		for y := 0; y < world.y; y++ {
			styleBoarder := tcell.StyleDefault.Foreground(tcell.ColorBlack).Background(tcell.ColorBlack)
			if world.w[x][y] {
				styleBoarder = tcell.StyleDefault.Foreground(tcell.ColorBlack).Background(tcell.ColorSpringGreen)
			}
			screen.SetContent(x*xd, y, '　', nil, styleBoarder)
		}
	}
	drawText(0, world.y,"Gen:" + strconv.Itoa(world.generaition),tcell.ColorWhite,tcell.ColorBlack)
	drawText(0, world.y+1,"Status:"+world.status.String(),tcell.ColorWhite,tcell.ColorBlack)
	screen.Show()
}

// Stop stops the view
func (view *View) Stop() {
	//fmt.Println("View Stop start")
	screen.Fini()
	//fmt.Println("View Stop end")
}

func (view *View) drawTexts() {
	/*
		xOffset := boardXOffset + board.width*2 + 8
		yOffset := boardYOffset + 7

		view.drawText(xOffset, yOffset, "SCORE:", tcell.ColorLightGray, tcell.ColorDarkBlue)
	*/
}

func drawText(x int, y int, text string, fg tcell.Color, bg tcell.Color) {
	style := tcell.StyleDefault.Foreground(fg).Background(bg)
	var xd int = 0 
	for _ , char := range text {
		screen.SetContent(x+xd, y, rune(char), nil, style)
		xd += runewidth.RuneWidth(char)
	}
}
