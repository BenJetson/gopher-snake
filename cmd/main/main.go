package main

import (
	"github.com/hajimehoshi/ebiten/v2"

	snake "github.com/BenJetson/gopher-snake"
)

func main() {
	// Initialize the game object.
	var g snake.Game

	// Configure game engine parameters.
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")

	// Run the main game loop.
	err := ebiten.RunGame(&g)
	if err != nil {
		panic(err)
	}
}
