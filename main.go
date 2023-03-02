package main

import (
	"fmt"
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	screenWidth  = 1280
	screenHeight = 720
)

var (
	spriteImage *ebiten.Image
	sprite      = &Sprite{
		X: 373,
		Y: 73,
		//X: screenWidth / 2,
		//Y: screenHeight / 2,
	}
)

type Sprite struct {
	X float64
	Y float64
}

var (
	meatImage *ebiten.Image
	meats     []*Meat = []*Meat{
		{X: 415, Y: 300},
		{X: 500, Y: 500},
		{X: 700, Y: 100},
		{X: 950, Y: 500},
		{X: 820, Y: 200},
	}
)

type Meat struct {
	X float64
	Y float64
}

func update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	screen.Fill(color.RGBA{10, 66, 202, 191})

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		sprite.X -= 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		sprite.X += 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		sprite.Y += 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		sprite.Y -= 2
	}

	for _, meat := range meats {
		geo2 := ebiten.GeoM{}
		geo2.Translate(meat.X, meat.Y)

		screen.DrawImage(meatImage, &ebiten.DrawImageOptions{
			GeoM:          geo2,
			ColorM:        ebiten.ColorM{},
			CompositeMode: 0,
		})
	}

	geo := ebiten.GeoM{}
	geo.Translate(-float64(spriteImage.Bounds().Dx())/2, -float64(spriteImage.Bounds().Dy())/2)
	geo.Translate(sprite.X, sprite.Y)

	scaleFactor := 4.0
	geo.Scale(scaleFactor, scaleFactor)

	width := float64(spriteImage.Bounds().Size().X)
	height := float64(spriteImage.Bounds().Size().Y)
	//width := float64(spriteImage.Bounds().Size().X) * scaleFactor
	//height := float64(spriteImage.Bounds().Size().Y) * scaleFactor
	geo.Translate(sprite.X-width/2-float64(spriteImage.Bounds().Dx())/2, sprite.Y-height/2-float64(spriteImage.Bounds().Dy())/2)

	screen.DrawImage(spriteImage.SubImage(image.Rect(0, 0, 23, 23)).(*ebiten.Image), &ebiten.DrawImageOptions{
		GeoM:          geo,
		ColorM:        ebiten.ColorM{},
		CompositeMode: 0,
	})

	//screen.DrawImage(spriteImage, &ebiten.DrawImageOptions{
	//	GeoM:          geo,
	//	ColorM:        ebiten.ColorM{},
	//	CompositeMode: 0,
	//})

	positionText := fmt.Sprintf("Sprite position: (%v, %v)", sprite.X, sprite.Y)
	ebitenutil.DebugPrint(screen, positionText)

	for i := len(meats) - 1; i >= 0; i-- {
		meat := meats[i]
		//meatRect := image.Rect(int(meat.X), int(meat.Y), int(meat.X)+meatImage.Bounds().Dx(), int(meat.Y)+meatImage.Bounds().Dy())
		meatRect := image.Rect(int(meat.X), int(meat.Y), int(meat.X+float64(meatImage.Bounds().Dx())), int(meat.Y+float64(meatImage.Bounds().Dy())))
		//meatRect := image.Rect(int(meat.X), int(meat.Y), int(meat.X+meatImage.Bounds().Size().X), int(meat.Y+meatImage.Bounds().Size().Y))
		spriteRect := image.Rect(int(sprite.X), int(sprite.Y), int(sprite.X)+spriteImage.Bounds().Dx(), int(sprite.Y)+spriteImage.Bounds().Dy())

		if spriteRect.Overlaps(meatRect) {
			// remove the meat from the list
			meats = append(meats[:i], meats[i+1:]...)
		} else {
			// draw the meat
			geo3 := ebiten.GeoM{}
			geo3.Translate(meat.X, meat.Y)
			screen.DrawImage(meatImage, &ebiten.DrawImageOptions{
				GeoM:          geo3,
				ColorM:        ebiten.ColorM{},
				CompositeMode: 0,
			})
		}
	}

	return nil

}

func main() {

	var err error
	spriteImage, _, err = ebitenutil.NewImageFromFile("assets/DinoSprites-doux.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	meatImage, _, err = ebitenutil.NewImageFromFile("assets/meat3.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	if err := ebiten.Run(update, screenWidth, screenHeight, 1, "golang-demo"); err != nil {
		log.Fatal(err)
	}
}
