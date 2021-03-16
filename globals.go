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

	Board struct {
		x,y int
	}

	//World : Lige Game World
	World struct {
		generaition int
		x, y int //width, height
		w [][] bool
		status Status
     	color tcell.Color
	}

	//Event : Keyboard Input Event
	Event struct {
		Type string
	}

	//Config : Config.toml
	Config struct {
		Gen General `mapstructure:"general"`
	}

	//General : General setting on Config.toml
	General struct {
		Maxgen int `mapstructure:"maxgen"`
		Worldsize int `mapstructure:"worldsize"`
		Color string `mapstructure:"color"`
	}
)

var (
	logger *log.Logger
	screen tcell.Screen
	view   *View
	world *World
	board *Board
)
