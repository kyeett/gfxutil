package main

import (
	"image"
	"image/color"
	"math"

	"golang.org/x/image/colornames"

	"github.com/kyeett/gfxutil"
	"github.com/peterhellberg/gfx"
)

const shadowLength = 100

func main() {
	r := gfx.R(0, 0, 100, 100).Moved(gfx.V(20, 20))

	// Draw input polygons
	polygonImg := gfx.NewImage(200, 200, color.Transparent)
	polygon := gfxutil.NewRectPolygon(r)
	gfx.DrawPolygon(polygonImg, polygon, 0, colornames.Pink)
	gfx.SavePNG("polygons.png", polygonImg)

	// Draw shadows
	shadowImg := gfx.NewImage(200, 200, color.Transparent)
	c := color.Black
	offset := gfx.V(shadowLength, 0).Rotated(3 * math.Pi / 8)
	gfxutil.DrawShadows(shadowImg, polygon, offset, c)
	gfx.SavePNG("shadows.png", shadowImg)

	// Combine
	gfx.DrawOver(shadowImg, shadowImg.Bounds(), polygonImg, image.ZP)
	gfx.SavePNG("combined.png", shadowImg)
}
