package main

import (
	"github.com/adam000/Go-SDL2/sdl"
)

// Draw something! in SDL
func main() {
	sdl.Init(sdl.INIT_EVERYTHING)
	defer sdl.Quit()

	window, err := sdl.CreateWindow("Hello world!", sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED,
		800, 600, sdl.WINDOW_OPENGL)

	if (err != nil) {
		panic(err)
	}
	defer window.Destroy()

	// Polling loop
	for {
		event := sdl.PollEvent()
		if event == nil {
			continue
		}
	}
}
