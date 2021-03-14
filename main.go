package main

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

func main() {

	// Load setting from config.toml 
	conf, err := load()
	if err != nil {
		panic(err)
	}

	var maxGen = conf.Gen.Maxgen
	var x = conf.Gen.Worldsize
	var y = x

	// Create Start Menu
	NewBoard()
	
	// key event loop start
	event := make(chan Event)
	defer close(event)
	go inputLoop(event)

	// Loop until clicking Enter
	startLoop(event)
	board.StopBoard()
	//close(event)

	// Create World
	NewWorld(x, y)
	defer world.Stop()

	// key event loop start
	//event := make(chan Event)
	go inputLoop(event)
	//defer close(event)

	isFin := false
	isPause := false
	for i := 0; i < maxGen; i++ {

		//Interval & world update
		time.Sleep(time.Millisecond * 100)
		world.Update()

		//pause mode
		if isPause {
			pauseLoop(event)	
			world.Restart()
			isPause = false
		}

		// Complete this game
		if isFin {
			break
		}

		select {
			case keyin := <-event:

				if keyin.Type == "fin" {
					// clicking Ctrl + C
					isFin = true
				}else if keyin.Type == "pause" {
					// clicking Space
					world.Pause()
					isPause = true
					keyin.Type = "restart"
				}
			default:
				continue
		}
	}

}

// Lopping unitl clicking EnterKey for Game Start
func startLoop(event chan Event){
	
	start := false
	for {		
		select {
			case keyin := <-event:
				if keyin.Type == "start" {
					// clicking Enter Key, start Game
					start = true
				}
			default:
				continue
		}

		if start {
			break
		}
		time.Sleep(time.Millisecond * 100)
	}
}

// Lopping unitl clicking EnterKey for Pause Mode
func pauseLoop(event chan Event){
	
	restart := false
	for {		
		select {
			case keyin := <-event:
				if keyin.Type == "pause" {
					// clicking Enter Key, cancel pause mode
					restart = true
				}
			default:
				continue
		}

		if restart {
			break
		}

		time.Sleep(time.Millisecond * 100)
	}
}

func load() (*Config, error){

	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	var c Config
	err := viper.ReadInConfig()

	if err != nil {
		return nil, fmt.Errorf("can't read the config file : %s", err)
	}

	err = viper.Unmarshal(&c)
	if err != nil {
		fmt.Println("NG")
		return nil, fmt.Errorf("unable to decode into struct : %s", err)
	}

	return &c, nil
}