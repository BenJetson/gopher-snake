package snake

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

func coordinatesToGeo(x, y int) (geo ebiten.GeoM) {
	geo.Translate(float64(x), float64(y))
	return
}

func rectToGeo(r image.Rectangle, offsetX, offsetY int) ebiten.GeoM {
	return coordinatesToGeo(r.Min.X+offsetX, r.Min.Y+offsetY)
}
