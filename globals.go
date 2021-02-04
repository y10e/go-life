package main

import (
	"log"

	"github.com/gdamore/tcell"
)

const (
	boardXOffset = 4
	boardYOffset = 2
)

type (
	// View is the display engine
	View struct {
	}

	World struct {
		generaition int
		x, y int //width, height
		w [][] bool
	}

)

var (
	logger *log.Logger
	screen tcell.Screen
	view   *View
	world *World
)
