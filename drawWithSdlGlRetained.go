package main

import (
	"encoding/binary"
	"math"
	"unsafe"

	mat "bitbucket.org/zombiezen/math3/mat32"
	"github.com/adam000/Go-SDL/sdl"
	"github.com/go-gl/gl"
)

const basicVertexShader = `
uniform mat4 uProjMatrix;
uniform mat4 uViewMatrix;
uniform mat4 uModelMatrix;

attribute vec3 aPosition;
attribute vec3 aColor;

void main() {
    vec4 vPosition = vec4(aPosition.x, aPosition.y, aPosition.z, 1.0);
    vPosition = uViewMatrix * vPosition;
    gl_Position = uProjMatrix * vPosition;

    gl_FrontColor = vec4(aColor.r, aColor.g, aColor.b, 1.0);
}
`

func getGlError(context string) {
	switch gl.GetError() {
	case gl.NO_ERROR:
	case gl.INVALID_ENUM:
		panic("Invalid enum at " + context)
	case gl.INVALID_VALUE:
		panic("Invalid value at " + context)
	case gl.INVALID_OPERATION:
		panic("Invalid operation at " + context)
	case gl.INVALID_FRAMEBUFFER_OPERATION:
		panic("Invalid framebuffer operation at " + context)
	case gl.OUT_OF_MEMORY:
		panic("Out of memory at " + context)
	case gl.STACK_UNDERFLOW:
		panic("Stack underflow at " + context)
	case gl.STACK_OVERFLOW:
		panic("Stack overflow at " + context)
	}
}

func convertMat32(mat *mat.Matrix) *[16]float32 {
	return (*[16]float32)(unsafe.Pointer(mat))
}

func makeSymmetricProjectionMatrix() mat.Matrix {
	// 80 degrees in radians
	fovy := 1.396
	aspect := 4 / 3.0
	nearPlane := 0.01
	farPlane := 100.0

	screenRange := math.Tan(fovy/2) * nearPlane
	right := screenRange * aspect
	top := screenRange

	var mat mat.Matrix
	mat[0][0] = float32(nearPlane / right)
	mat[1][1] = float32(nearPlane / top)
	mat[2][2] = float32(-(farPlane + nearPlane) / (farPlane - nearPlane))
	mat[2][3] = -1
	mat[3][2] = float32((-2 * farPlane * nearPlane) / (farPlane - nearPlane))

	return mat
}

// Draw something! in SDL using OpenGL with shaders / retained mode
func drawWithSdlGlRetained() {
	const HEIGHT int = 800
	const WIDTH int = 600

	triVertexes := [][3]gl.GLfloat{
		{2.0, 0.0, -5.0},
		{0.0, 4.0, -5.0},
		{-2.0, 0.0, -5.0},
	}

	triColors := [][3]gl.GLfloat{
		{1.0, 0.0, 0.0},
		{0.0, 1.0, 0.0},
		{0.0, 0.0, 1.0},
	}

	// SDL Initialization
	sdl.Init(sdl.INIT_EVERYTHING)
	defer sdl.Quit()

	screen := sdl.SetVideoMode(HEIGHT, WIDTH, 32, sdl.OPENGL)
	_ = screen

	// OpenGL initialization
	if err := gl.Init(); err != 0 {
		panic("Problem in GL initialization")
	}

	// Initialize what color and depth to clear to
	gl.ClearColor(0, 0, 0, 0)
	gl.ClearDepth(1.0)
	gl.DepthFunc(gl.LEQUAL)
	gl.Enable(gl.DEPTH_TEST)

	// Shader initialization
	vertexShader := gl.CreateShader(gl.VERTEX_SHADER)
	vertexShader.Source(basicVertexShader)
	vertexShader.Compile()
	if success := vertexShader.Get(gl.COMPILE_STATUS); success == gl.FALSE {
		panic("Shader failed to compile: " + vertexShader.GetInfoLog())
	}

	program := gl.CreateProgram()
	program.AttachShader(vertexShader)
	program.Link()
	program.Validate()
	if success := program.Get(gl.VALIDATE_STATUS); success == gl.FALSE {
		panic("Program validation failed: " + program.GetInfoLog())
	}

	program.Use()

	posAttrib := program.GetAttribLocation("aPosition")
	colorAttrib := program.GetAttribLocation("aColor")
	projUniform := program.GetUniformLocation("uProjMatrix")
	viewUniform := program.GetUniformLocation("uViewMatrix")
	modelUniform := program.GetUniformLocation("uModelMatrix")

	// Buffer initialization
	posBuffer := gl.GenBuffer()
	posBuffer.Bind(gl.ARRAY_BUFFER)
	gl.BufferData(gl.ARRAY_BUFFER, binary.Size(triVertexes), triVertexes, gl.STATIC_DRAW)

	colorBuffer := gl.GenBuffer()
	colorBuffer.Bind(gl.ARRAY_BUFFER)
	gl.BufferData(gl.ARRAY_BUFFER, binary.Size(triColors), triColors, gl.STATIC_DRAW)


	// Draw - this portion suitable for a loop
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	program.Use()

	// Set uniforms - projection and view only need to be set once, model is
	// done per-object
	projection := makeSymmetricProjectionMatrix()
	projUniform.UniformMatrix4f(false, convertMat32(&projection))

	// Note: The view / model matrices are not really needed for this exercise,
	// but are useful in other applications.
	viewUniform.UniformMatrix4f(false, convertMat32(&mat.Identity))
	modelUniform.UniformMatrix4f(false, convertMat32(&mat.Identity))

	posAttrib.EnableArray()
	posBuffer.Bind(gl.ARRAY_BUFFER)
	posAttrib.AttribPointer(3, gl.FLOAT, false, 0, nil)
	getGlError("posBuffer")

	colorAttrib.EnableArray()
	colorBuffer.Bind(gl.ARRAY_BUFFER)
	colorAttrib.AttribPointer(3, gl.FLOAT, false, 0, nil)
	getGlError("colorBuffer")

	gl.DrawArrays(gl.TRIANGLES, 0, 3)
	getGlError("DrawArrays")

	posAttrib.DisableArray()
	colorAttrib.DisableArray()

	gl.ProgramUnuse()
	sdl.GL_SwapBuffers()

	sdl.Delay(5000)
}
