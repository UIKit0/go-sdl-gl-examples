package main

import (
	"github.com/adam000/Go-SDL2/sdl"
	"github.com/go-gl/gl"
)

// Draw something! in SDL using OpenGL immediate
func drawWithSdlGlImmediate() {
	const HEIGHT int = 800
	const WIDTH int = 600

	sdl.Init(sdl.INIT_EVERYTHING)
	defer sdl.Quit()

	screen := sdl.SetVideoMode(HEIGHT, WIDTH, 32, sdl.OPENGL)
	_ = screen

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

	sdl.GL_SwapBuffers()
	sdl.Delay(5000)
}

