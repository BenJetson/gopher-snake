package snake

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type fragment struct {
	color color.RGBA
	pos   gridPosition
}

func (f *fragment) Draw(frameImg *ebiten.Image) {
	img := ebiten.NewImage(gridCellSize, gridCellSize)
	img.Fill(f.color)

	frameImg.DrawImage(img, &ebiten.DrawImageOptions{
		GeoM: f.pos.geom(),
	})
}

type snake struct {
	fragments []*fragment
	color     color.RGBA
	heading   direction
	gd        gridDimensions
	dead      bool
}

func (s *snake) readKeyboard() {
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		s.heading = directionNorth
	} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		s.heading = directionSouth
	} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {
		s.heading = directionEast
	} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {
		s.heading = directionWest
	}
}

func (s *snake) move() {
	if s.heading != directionStationary {
		// Make a copy of the head node so we can move it.
		headCopy := *(s.fragments[0])

		// Adjust the coordinates of the new head based on the direction the
		// snake is currently heading.
		switch s.heading {
		case directionNorth:
			headCopy.pos.y--
		case directionSouth:
			headCopy.pos.y++
		case directionEast:
			headCopy.pos.x++
		case directionWest:
			headCopy.pos.x--
		}

		// Check for bounds on new head.
		if headCopy.pos.isOutOfBounds() {
			s.dead = true
			headCopy = *(s.fragments[0]) // FIXME
		}

		// Move all fragments of the snake. The new position will be the
		// position of the fragment before it.
		for i := 1; i < len(s.fragments); i++ {
			s.fragments[i].pos = s.fragments[i-1].pos
		}

		// Copy the new head node back into the first slot, now that all other
		// fragments have moved.
		*(s.fragments[0]) = headCopy
	}
}

func (s *snake) Update(a *apple) error {
	if len(s.fragments) < 1 {
		s.fragments = append(s.fragments, &fragment{
			color: s.color,
			pos:   makeRandomGridPos(s.gd),
		})
	}

	s.readKeyboard()
	s.move()

	return nil
}

func (s *snake) Draw(frameImg *ebiten.Image) {
	for _, f := range s.fragments {
		f.Draw(frameImg)
	}
}
