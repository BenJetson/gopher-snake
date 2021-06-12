package snake

import (
	ebiten "github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var _ ebiten.Game = (*Game)(nil)

type Game struct {
	// now time.Duration
	g grid
}

const margin = 20

func NewGame() Game {
	return Game{
		g: newGrid(margin, margin, ScreenWidth-margin*2, ScreenHeight-margin*2),
	}
}

// Update updates a game by one tick.
func (g *Game) Update() error {
	if g.g.s.dead && inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		*g = NewGame() // FIXME
	}

	return g.g.Update()
}

// Draw draws the game screen. The given argument represents a screen image.
func (g *Game) Draw(screen *ebiten.Image) {
	// screen.Fill(color.RGBA{R: 0xff, G: 0x0, B: 0x0, A: 0xff})
	// const rectOffset = 10
	// ebitenutil.DrawRect(screen,
	// 	rectOffset,
	// 	rectOffset,
	// 	ScreenWidth-2*rectOffset,
	// 	ScreenHeight-2*rectOffset,
	// 	color.RGBA{R: 0x00, G: 0x0, B: 0xff, A: 0xff})
	// ebitenutil.DebugPrint(screen, "Hello, World!")

	g.g.Draw(screen)
}

const ScreenWidth = 320
const ScreenHeight = 240

// Layout accepts a native outside size in device-independent pixels
// and returns the game's logical screen size. On desktops, the outside is a
// window or a monitor (fullscreen mode)
//
// Even though the outside size and the screen size differ, the rendering scale
// is automatically adjusted to fit with the outside.
//
// You can return a fixed screen size if you don't care, or you can also return
// a calculated screen size adjusted with the given outside size.
func (g *Game) Layout(
	outsideWidth, outsideHeight int,
) (screenWidth, screenHeight int) {

	// screenWidth = outsideWidth
	// screenHeight = outsideHeight

	screenWidth = ScreenWidth
	screenHeight = ScreenHeight

	return
}
