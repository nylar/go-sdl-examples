package main

import (
	"log"

	"github.com/veandco/go-sdl2/sdl"
)

var (
	WindowTitle  = "Game"
	WindowWidth  = 720
	WindowHeight = 640
	Fullscreen   = false
)

type Game struct {
	Window   *sdl.Window
	Renderer *sdl.Renderer

	Running bool
}

func NewGame() (*Game, error) {
	var windowOpt uint32 = sdl.WINDOW_SHOWN
	if Fullscreen {
		windowOpt = sdl.WINDOW_FULLSCREEN
	}

	window, err := sdl.CreateWindow(WindowTitle, 0, 0, WindowWidth, WindowHeight, windowOpt)
	if err != nil {
		return nil, err
	}

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		return nil, err
	}

	return &Game{
		Window:   window,
		Renderer: renderer,
	}, nil
}

func (g *Game) Destroy() {
	g.Renderer.Destroy()
	g.Window.Destroy()
}

func (g *Game) HandleEvents(event sdl.Event) {
	for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.QuitEvent:
			g.Running = false
		}
	}
}

func main() {
	var event sdl.Event

	game, err := NewGame()
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
