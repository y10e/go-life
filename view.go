package main

import (
	"strconv"

	"github.com/gdamore/tcell"
	"github.com/mattn/go-runewidth"
)

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
	for index, char := range text {
		screen.SetContent(x+index, y, rune(char), nil, style)
	}
}
