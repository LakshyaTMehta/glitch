package shapes

import (
  "glitch/pkgs/render"
  "math"
)

func Elipse(a, b float64, res int) (vertices []render.Point) {
  vertices = make([]render.Point, 0, res)
  root_of_unity := math.Pi * 2 / float64(res)
  for i := 0; i < res; i++ {
    theta := float64(i) * root_of_unity
    x := a / 2 * math.Cos(theta)
    y := b / 2 * math.Sin(theta)
    vertices = append(vertices, render.Point{X: int(x), Y: int(y)})
  }
  return vertices
}
