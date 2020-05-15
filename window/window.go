package window

import (
	"image"
	"os"

	// image/png for png conversion
	_ "image/png"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

// LaunchWindow launches the primary window
func LaunchWindow() {
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

	for !win.Closed() {
		win.Update()
	}
}

// RenderPicture validates and prepares a picture to be a sprite
func RenderPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}

// RenderSprite returns a configured pixel sprite
func RenderSprite(path string) *pixel.Sprite {
	pic, err := RenderPicture(path)
	if err != nil {
		panic(err) // TODO handle picture load errors
	}
	sprite := pixel.NewSprite(pic, pic.Bounds())
	return sprite
}
