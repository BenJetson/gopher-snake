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
	extend    bool
}

func (s *snake) readKeyboard() {
	if s.heading != directionSouth &&
		inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		s.heading = directionNorth
	} else if s.heading != directionNorth &&
		inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		s.heading = directionSouth
	} else if s.heading != directionWest &&
		inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {
		s.heading = directionEast
	} else if s.heading != directionEast &&
		inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {
		s.heading = directionWest
	}
}

func (s *snake) head() fragment {
	return *(s.fragments[0])
}

func (s *snake) searchForFood(a *apple) {
	head := s.head()

	if !a.eaten && a.pos == head.pos {
		a.eaten = true
		s.extend = true
	}
}

func (s *snake) move() {
	if s.heading != directionStationary {
		// Make a copy of the head node so we can move it.
		head := s.head()

		// Adjust the coordinates of the new head based on the direction the
		// snake is currently heading.
		switch s.heading {
		case directionNorth:
			head.pos.y--
		case directionSouth:
			head.pos.y++
		case directionEast:
			head.pos.x++
		case directionWest:
			head.pos.x--
		}

		// Check for bounds on new head.
		if head.pos.isOutOfBounds() {
			s.dead = true
			return // don't move!
		}

		// If the snake is extending, add a new fragment to the end.
		// No need to worry about position since it will get filled by
		// the one ahead when it moves.
		if s.extend {
			s.fragments = append(s.fragments, &fragment{color: s.color})
			s.extend = false
		}

		// Move all fragments of the snake. The new position will be the
		// position of the fragment before it.
		for i := len(s.fragments) - 1; i > 0; i-- {
			s.fragments[i].pos = s.fragments[i-1].pos
		}

		// Copy the new head node back into the first slot, now that all other
		// fragments have moved.
		*(s.fragments[0]) = head
	}
}

func (s *snake) preventCannibalism() {
	for i, fa := range s.fragments {
		for j, fb := range s.fragments {
			if i != j && fa.pos == fb.pos {
				s.dead = true
			}
		}
	}
}

func (s *snake) Update(a *apple) error {
	if len(s.fragments) < 1 {
		s.fragments = append(s.fragments, &fragment{
			color: s.color,
			pos:   makeRandomGridPos(s.gd),
		})
	}

	if !s.dead {
		s.readKeyboard()
		s.searchForFood(a)
		s.move()
		s.preventCannibalism()
	} else {
		s.heading = directionStationary
	}

	return nil
}

func (s *snake) Draw(frameImg *ebiten.Image) {
	for _, f := range s.fragments {
		f.Draw(frameImg)
	}
}
