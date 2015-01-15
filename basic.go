package main

import (
	"log"

	"github.com/nylar/go-sdl-examples/base"
	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	base.WindowTitle = "Basic Window"
	var event sdl.Event

	game, err := base.NewGame()
	if err != nil {
		log.Fatalln("Could not create game instance. Error: %s", err.Error())
	}
	defer game.Destroy()

	game.Running = true

	for game.Running {
		game.Renderer.Clear()
		game.Renderer.Present()

		game.HandleEvents(event)
	}
}
