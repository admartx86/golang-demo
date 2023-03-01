package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

var (
	spriteImage *ebiten.Image
	sprite      = &Sprite{
		X: screenWidth / 2,
		Y: screenHeight / 2,
	}
)

type Sprite struct {
	X float64
	Y float64
}

func update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	screen.Fill(color.RGBA{10, 66, 202, 191})

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		sprite.X -= 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		sprite.X += 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		sprite.Y += 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		sprite.Y -= 5
	}

	//screen.DrawImage(spriteImage, &ebiten.DrawImageOptions{})

	screen.DrawImage(spriteImage, &ebiten.DrawImageOptions{
		GeoM:          ebiten.TranslateGeo(sprite.X, sprite.Y),
		ColorM:        ebiten.ColorM{},
		CompositeMode: 0,
		Filter:        0,
		ImageParts:    nil,
		//Parts:         []ebiten.ImagePart{},
		//SourceRect:    &image.Rectangle{},
	})

	return nil
}

func main() {
	var err error
	spriteImage, _, err = ebitenutil.NewImageFromFile("assets/DinoSprites-doux.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	if err := ebiten.Run(update, screenWidth, screenHeight, 1, "golang-demo"); err != nil {
		log.Fatal(err)
	}
}
