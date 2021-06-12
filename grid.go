package snake

import (
	"image"
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

const gridCellSize = 5

type gridDimensions struct {
	cols, rows int
}

func makeGridDimensions(width, height int) gridDimensions {
	return gridDimensions{
		cols: width / gridCellSize,
		rows: height / gridCellSize,
	}
}

type grid struct {
	s snake
	a apple

	d gridDimensions
	r image.Rectangle
}

func newGrid(x, y, w, h int) grid {
	d := makeGridDimensions(w, h)

	return grid{
		r: makeRect(x, y, w, h),
		d: d,

		s: snake{gd: d, color: color.RGBA{0x7c, 0x4d, 0xff, 0xff}},
		a: apple{pos: gridPosition{d: d}},
	}
}

func (g *grid) Update() error {
	g.s.Update(&g.a)
	g.a.Update()

	return nil
}

func (g *grid) Draw(screen *ebiten.Image) {
	img := ebiten.NewImage(g.r.Dx(), g.r.Dy())

	// Debugging grid.
	for x := 0; x <= g.r.Dx(); x += gridCellSize {
		for y := 0; y <= g.r.Dy(); y += gridCellSize {
			img.Set(x, y, color.Gray{Y: 0x88})
		}
	}

	// for x := 0; x < g.r.Max.X; x += gridCellSize {
	// 	ebitenutil.DrawLine(img, float64(x), 0, float64(x), float64(g.r.Max.Y), color.Gray{Y: 0xff})
	// 	fmt.Println(x)
	// }

	// Draw snake and apple.
	g.a.Draw(img)
	g.s.Draw(img)

	screen.DrawImage(img, &ebiten.DrawImageOptions{
		GeoM: rectToGeo(g.r, 0, 0),
	})
}

type gridPosition struct {
	d    gridDimensions
	x, y int
}

func makeRandomGridPos(d gridDimensions) gridPosition {
	return gridPosition{
		x: rand.Intn(d.rows + 1),
		y: rand.Intn(d.cols + 1),
		d: d,
	}
}

func (p gridPosition) isZero() bool {
	return p.x == 0 && p.y == 0
}

func (p gridPosition) rect() image.Rectangle {
	return makeRect(
		p.x*gridCellSize, p.y*gridCellSize,
		gridCellSize, gridCellSize,
	)
}

func (p gridPosition) geom() ebiten.GeoM {
	return coordinatesToGeo(p.x*gridCellSize, p.y*gridCellSize)
}

func (p *gridPosition) isOutOfBounds() bool {
	return p.x < 0 ||
		p.y < 0 ||
		p.x >= p.d.cols ||
		p.y >= p.d.rows
}
