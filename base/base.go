package base

import (
	"errors"

	"github.com/veandco/go-sdl2/sdl"
)

var (
	WindowTitle         = "Game"
	WindowWidth         = 720
	WindowHeight        = 640
	WindowOpt    uint32 = sdl.WINDOW_OPENGL
)

type Game struct {
	Window   *sdl.Window
	Renderer *sdl.Renderer
	Context  sdl.GLContext

	Running bool
}

func NewGame() (*Game, error) {
	window, err := sdl.CreateWindow(WindowTitle, 0, 0, WindowWidth, WindowHeight, WindowOpt)
	if err != nil {
		return nil, err
	}

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		return nil, err
	}

	context := sdl.GL_CreateContext(window)
	if context == nil {
		return nil, errors.New("Could not create OpenGL context")
	}

	return &Game{
		Window:   window,
		Renderer: renderer,
		Context:  context,
	}, nil
}

func (g *Game) Destroy() {
	sdl.GL_DeleteContext(g.Context)
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
