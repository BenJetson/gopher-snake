package main

import (
	"github.com/hajimehoshi/ebiten/v2"

	snake "github.com/BenJetson/gopher-snake"
)

func main() {
	// Initialize the game object.
	g := snake.NewGame()

	// Configure game engine parameters.
	ebiten.SetWindowSize(snake.ScreenWidth*3, snake.ScreenHeight*3)
	ebiten.SetWindowTitle("Hello, World!")

	ebiten.SetMaxTPS(8) // FIXME

	// Run the main game loop.
	err := ebiten.RunGame(&g)
	if err != nil {
		panic(err)
	}
}
