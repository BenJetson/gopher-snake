package snake

import "image"

func makeRect(x, y, w, h int) image.Rectangle {
	min := image.Point{x, y}
	max := min.Add(image.Point{w, h})

	return image.Rectangle{min, max}
}
