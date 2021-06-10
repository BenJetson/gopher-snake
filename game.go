package snake

import (
	ebiten "github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var _ ebiten.Game = (*Game)(nil)

type Game struct {
	// now time.Duration
}

// Update updates a game by one tick.
func (g *Game) Update() error {
	return nil
}

// Draw draws the game screen. The given argument represents a screen image.
func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!")
}

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

	screenWidth = 320
	screenHeight = 240

	return
}
