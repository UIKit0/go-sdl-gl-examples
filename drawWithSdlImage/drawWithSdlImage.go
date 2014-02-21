package main

import (
	"github.com/adam000/Go-SDL2/sdl"
)

// Draw something! in SDL
func main() {
	var img string = "../imgs/preview2.jpg"

	sdl.Init(sdl.InitEverything)
	defer sdl.Quit()

	window, err := sdl.NewWindow("Hello world!", sdl.WindowPosCentered, sdl.WindowPosCentered, 800, 600, sdl.WindowOpenGL)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	renderer, err := sdl.NewRenderer(window, -1, 0)
	if err != nil {
		panic(err)
	}
	defer renderer.Destroy()

	surf, err := sdl.LoadImage(img)
	if err != nil {
		panic(err)
	}

	tex, err := surf.ToTexture(renderer)
	if err != nil {
		panic(err)
	}
	surf.Free()
	defer tex.Destroy()

	quit := false
	for {
		if e := sdl.PollEvent(); e != nil {

			switch t := e.Type(); t {
			case sdl.QuitEv:
				quit = true
			}

			if quit {
				break
			}
		}

		renderer.Clear()
		renderer.CopyTexture(tex, nil, nil)
		renderer.Present()
	}
}
