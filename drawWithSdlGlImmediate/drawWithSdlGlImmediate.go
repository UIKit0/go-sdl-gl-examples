package main

import (
	"time"

	"github.com/adam000/Go-SDL2/sdl"
	"github.com/go-gl/gl"
)

// Draw something! in SDL using OpenGL immediate
func main() {
	const HEIGHT int = 800
	const WIDTH int = 600

	sdl.Init(sdl.InitEverything)
	defer sdl.Quit()

	window, err := sdl.NewWindow("Hello world!", sdl.WindowPosCentered, sdl.WindowPosCentered, 800, 600, sdl.WindowOpenGL)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	context, err := sdl.NewGLContext(window)
	if err != nil {
		panic(err)
	}
	defer context.Destroy()

	if err := gl.Init(); err != 0 {
		panic("Problem in GL initialization")
	}

	gl.Enable(gl.TEXTURE_2D)
	gl.Viewport(0, 0, WIDTH, HEIGHT)
	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	gl.Ortho(-float64(WIDTH)/2.0, float64(WIDTH)/2.0, -float64(HEIGHT)/2.0, float64(HEIGHT)/2.0, -100, 100)

	gl.ClearColor(0, 0, 0, 0)
	gl.Clear(gl.COLOR_BUFFER_BIT)

	gl.Begin(gl.TRIANGLES)

	gl.Color3f(1.0, 0.0, 0.0)
	gl.Vertex2f(0.0, 100.0)

	gl.Color3f(1.0, 0.0, 0.0)
	gl.Vertex2f(-100.0, 0.0)

	gl.Color3f(1.0, 0.0, 0.0)
	gl.Vertex2f(100.0, 0.0)

	gl.End()

	window.GLSwap()
	time.Sleep(time.Second * 7)
}
