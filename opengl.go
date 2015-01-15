package main

import (
	"log"

	"github.com/go-gl/gl"
	"github.com/nylar/go-sdl-examples/base"
	"github.com/veandco/go-sdl2/sdl"
)

func InitGL() {
	gl.ClearColor(0, 0, 0, 0)
	gl.ClearDepth(1)
	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LEQUAL)
	gl.Viewport(0, 0, base.WindowWidth, base.WindowHeight)
}

func Scene() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.Begin(gl.TRIANGLES)
	gl.Color3f(1, 0, 0)
	gl.Vertex3f(0, .75, 0)
	gl.Color3f(0, 1, 0)
	gl.Vertex3f(-.75, -.75, 0)
	gl.Color3f(0, 0, 1)
	gl.Vertex3f(.75, -.75, 0)
	gl.End()
}

func main() {
	base.WindowTitle = "OpenGL Window"
	var event sdl.Event

	game, err := base.NewGame()
	if err != nil {
		log.Fatalln("Could not create game instance. Error: %s", err.Error())
	}
	defer game.Destroy()

	InitGL()

	game.Running = true

	for game.Running {
		Scene()
		sdl.GL_SwapWindow(game.Window)

		game.HandleEvents(event)
	}

}
