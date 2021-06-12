package snake

import (
	"image/color"

	ebiten "github.com/hajimehoshi/ebiten/v2"
)

type apple struct {
	pos   gridPosition
	eaten bool
}

func (a *apple) Update() error {
	if a.eaten || a.pos.isZero() {
		a.pos = makeRandomGridPos(a.pos.d)
	}

	return nil
}

func (a *apple) Draw(frameImg *ebiten.Image) {
	img := ebiten.NewImage(gridCellSize, gridCellSize)
	img.Fill(color.RGBA{0xf4, 0x43, 0x36, 0xff})

	frameImg.DrawImage(img, &ebiten.DrawImageOptions{
		GeoM: a.pos.geom(),
	})
}
