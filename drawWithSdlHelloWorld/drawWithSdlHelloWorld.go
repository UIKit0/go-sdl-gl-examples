package main

import (
	"github.com/adam000/Go-SDL2/sdl"
)

// Draw something! in SDL
func main() {
	sdl.Init(sdl.InitEverything)
	defer sdl.Quit()

	window, err := sdl.NewWindow("Hello world!", sdl.WindowPosCentered, sdl.WindowPosCentered,
		800, 600, sdl.WindowOpenGL)

	if err != nil {
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
