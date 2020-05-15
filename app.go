package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	w "chance/window"
	"golang.org/x/image/colornames"
)

func run() {
	// w.LaunchWindow()
	cfg := pixelgl.WindowConfig{
		Title:  "Chance",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	win.Clear(colornames.Slategray)
	sprite := w.RenderSprite("./sprites/characters/helloworld.png")
	sprite.Draw(win, pixel.IM.Moved(win.Bounds().Center()))

	for !win.Closed() {
		win.Update()
	}
}

// Use PixelGL to lock-down Go concurrency https://godoc.org/runtime#LockOSThread
func main() {
	pixelgl.Run(run)
}
