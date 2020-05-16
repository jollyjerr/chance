package main

import (
	// w "chance/window"
	chanceModel "chance/models/chance"
	// "time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
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

	chance := chanceModel.Construct()
	// angle := 0.0

	// last := time.Now()
	for !win.Closed() {
		win.Clear(colornames.Slategray)
		// deltaTime := time.Since(last).Seconds()
		// last = time.Now() 

		// angle += 3 * deltaTime

		mat := pixel.IM
		// mat = mat.Rotated(pixel.ZV, angle)
		mat = mat.Scaled(pixel.ZV, 6)
		mat = mat.Moved(win.Bounds().Center())
		chance.Sprite.Draw(win, mat)


		win.Update()
	}
}

// Use PixelGL to lock-down Go concurrency https://godoc.org/runtime#LockOSThread
func main() {
	pixelgl.Run(run)
}
