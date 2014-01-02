package main

import (
	"fmt"

	"github.com/adam000/Go-SDL2/sdl"
)

// Draw something! in SDL
func drawWithSdl() {
	var img string = "imgs/preview2.jpg"

	sdl.Init(sdl.INIT_EVERYTHING)
	defer sdl.Quit()

	window := sdl.CreateWindow("Hello world!", sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED, 800, 600, sdl.WINDOW_OPENGL)

	if (window == nil) {
		panic(sdl.GetError())
	}
	defer window.Destroy()

	renderer := sdl.CreateRenderer(window, -1, 0)
	defer renderer.Destroy()

	surf := sdl.Load(img)
	if surf == nil {
		panic(sdl.GetError())
	}

	tex := sdl.CreateTextureFromSurface(renderer, surf)
	surf.Free() // TODO or is it sdl.FreeSurface(surf)?
	defer tex.Destroy()

	for {
		if e := sdl.PollEvent; e != nil && e.type == sdl.QUIT {
			break
		}

		renderer.Clear()
		renderer.Copy(tex, nil, nil)
		renderer.Present()
	}
}
