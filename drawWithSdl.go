package main

import (
	"fmt"

	"github.com/banthar/Go-SDL/sdl"
)

// Draw something! in SDL
func drawWithSdl() {
	var img string = "imgs/preview2.jpg"

	sdl.Init(sdl.INIT_EVERYTHING)
	defer sdl.Quit()

	screen := sdl.SetVideoMode(800, 600, 32, sdl.SWSURFACE)

	surf := sdl.Load(img)
	if surf == nil {
		fmt.Printf("Error: %s\n", sdl.GetError())
	}
	screen.Blit(nil, surf, nil)

	screen.Flip()
	sdl.Delay(5000)
}
