package shapes

import (
  "glitch/pkgs/render"
)

func Translate (vertices []render.Point, x, y float64) {
  for i := range vertices {
    vertices[i].X += int(x)
    vertices[i].Y += int(y)
  }
}

func Scale (vertices []render.Point, x, y float64) {
  for i := range vertices {
    vertices[i].X *= int(x)
    vertices[i].Y *= int(y)
  }
}

