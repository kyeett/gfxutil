package gfxutil

import (
	"image"
	"image/color"

	"github.com/peterhellberg/gfx"
)

// DrawShadows draws a shadow with color and an offset on img
func DrawShadows(img *image.NRGBA, polygon gfx.Polygon, offset gfx.Vec, c color.Color) {
	for i := range polygon {
		p1 := polygon[i]
		p2 := polygon[(i+1)%len(polygon)]
		poly1 := []gfx.Vec{p1, p2, p1.Add(offset)}
		poly2 := []gfx.Vec{p2, p2.Add(offset), p1.Add(offset)}
		gfx.DrawPolygon(img, poly1, 0, c)
		gfx.DrawPolygon(img, poly2, 0, c)
	}
	gfx.DrawPolygon(img, polygon, 0, color.NRGBA{0, 0, 0, 1})
}

// NewRectPolygon returns a polygon with the four corners as points
func NewRectPolygon(r gfx.Rect) gfx.Polygon {
	return []gfx.Vec{
		r.Min,
		r.Min.Add(gfx.V(r.W(), 0)),
		r.Max,
		r.Max.Add(gfx.V(-r.W(), 0)),
	}
}
