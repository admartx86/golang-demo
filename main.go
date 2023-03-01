package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"image/color"
	"log"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

func update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	screen.Fill(color.RGBA{10, 66, 202, 191})

	ebitenutil.DebugPrint(screen, "Hello World!")

	return nil
}

func main() {
	if err := ebiten.Run(update, screenWidth, screenHeight, 1, "Hello World"); err != nil {
		log.Fatal(err)
	}
}
