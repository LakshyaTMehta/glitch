package face

import (
	"glitch/pkgs/render"
	"glitch/pkgs/shapes"
	"math"
	"time"

	"github.com/gdamore/tcell/v2"
)

type Eye struct {
  Pos render.Point
  Width, Height int
  res int
  eyeball_size, eyeball_d, eyeball_theta float64
  // eyeball_r and eyeball_size must be strictly between 0..1
  // eyeball_theta must be between 0..2 * math.Pi
  eye_vertices []render.Point
  eyeball_vertices []render.Point
}

func NewEye(xpos, ypos, w, h int) (e Eye) {
  e.Pos = render.Point{X: xpos, Y: ypos}
  e.Width = w
  e.Height = h
  e.res = 20
  e.eye_vertices = shapes.Elipse( float64(e.Width), 
                                  float64(e.Height), 
                                  e.res,
                                )
  shapes.Translate(e.eye_vertices, float64(e.Pos.X), float64(e.Pos.Y))
  // eyeball
  e.eyeball_size = 0.5
  eyeball_width := e.eyeball_size * float64(e.Width)
  eyeball_height := e.eyeball_size * float64(e.Height)
  e.eyeball_vertices = shapes.Elipse( eyeball_width,
                                      eyeball_height,
                                      e.res,
                                    )
  shapes.Translate(e.eyeball_vertices, float64(e.Pos.X), float64(e.Pos.Y))
  return e
}

func (e *Eye) Update() {
  e.eye_vertices = shapes.Elipse( float64(e.Width), 
                                  float64(e.Height), 
                                  e.res,
                                )
  shapes.Translate(e.eye_vertices, float64(e.Pos.X), float64(e.Pos.Y))
  // eyeball variable boundary
  e.eyeball_d = math.Max(math.Min(e.eyeball_d, 1), -1)
  e.eyeball_size = math.Max(math.Min(e.eyeball_size, 1), 0)
  e.eyeball_theta = math.Max(math.Min(e.eyeball_theta, math.Pi * 2), 0)
  // eyeball draw parameters
  ebx := float64(e.Pos.X) + e.eyeball_d * float64(e.Width) / 2 * math.Cos(e.eyeball_theta) * (1 - e.eyeball_size)
  eby := float64(e.Pos.Y) + e.eyeball_d * float64(e.Height) / 2 * math.Sin(e.eyeball_theta) * (1 - e.eyeball_size)
  eyeball_width := e.eyeball_size * float64(e.Width)
  eyeball_height := e.eyeball_size * float64(e.Height)
  e.eyeball_vertices = shapes.Elipse( eyeball_width,
                                      eyeball_height,
                                      e.res,
                                    )
  shapes.Translate(e.eyeball_vertices, ebx, eby)
}

func (e *Eye) Watchout() {
  theta := 0.0
  for ;; {
    if theta >= (math.Pi * 2) {
      theta = 0
    }
    e.eyeball_d = math.Sin(theta)
    theta += 0.01
    time.Sleep(time.Millisecond * 5)
  }
}

func (e Eye) Draw (s tcell.Screen, style tcell.Style) {
   render.Render(e.eye_vertices, tcell.RuneBlock, s, style)
   render.Render(e.eyeball_vertices, ' ', s, style)
}
