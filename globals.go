package main

import (
	"log"

	"github.com/gdamore/tcell"
)

const (
	boardXOffset = 4
	boardYOffset = 2
)

//Status : world status
type Status int

const (
	stop Status = iota
	running
	pause
)

func (s Status) String() string {
    switch s {
		case stop:
			return "stop"
		case running:
			return "running"
		case pause:
			return "pause"
		default:
			return "Unknown"
    }
}

type (
	// View is the display engine
	View struct {
	}

	//World : Lige Game World
	World struct {
		generaition int
		x, y int //width, height
		w [][] bool
		status Status
	}

	//Event : Keyboard Input Event
	Event struct {
		Type string
	}

)



var (
	logger *log.Logger
	screen tcell.Screen
	view   *View
	world *World
)
