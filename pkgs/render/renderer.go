package render

import (
	"math"
	"github.com/gdamore/tcell/v2"
)

type Point struct {
	X, Y int
}

func Render(v []Point, fill rune, s tcell.Screen, style tcell.Style){
	l := len(v)
	if l < 3 {
		return
	}
	// top-left & bottom-right bounds for scan
	scan_tl, scan_br := v[0], v[0]
	for _, vertex := range v {
		switch {
		case vertex.X < scan_tl.X:
			scan_tl.X = vertex.X
		case vertex.X > scan_br.X:
			scan_br.X = vertex.X
		}
		switch {
		case vertex.Y < scan_tl.Y:
			scan_tl.Y = vertex.Y
		case vertex.Y > scan_br.Y:
			scan_br.Y = vertex.Y
		}
	}

  // literal "edge" cases
  scan_br.X += 1
  scan_br.Y += 1
  scan_tl.X -= 1
  scan_tl.Y -= 1

	// scan
  horizontal_intersections := make([]int, 0, scan_br.X - scan_tl.X)
	for y := scan_tl.Y; y <= scan_br.Y; y++{
		prev_on_edge := false
		for x := scan_tl.X; x <= scan_br.X; x++{
			curr_on_edge := false
			for i := 0; i < l; i++{
				p1, p2 := v[i], v[(i + 1) % l]
				// calculate distance (delta) of pixel from the line between p1 & p2
				dx := p2.X - p1.X
				dy := p2.Y - p1.Y
				numerator := math.Abs(float64((dx * (y - p1.Y)) - (dy * (x - p1.X))))
				denominator := math.Sqrt(float64((dx * dx) + (dy * dy)))
				delta := numerator / denominator
				// check if the current scan position intersects with a shape edge
				curr_on_edge =	(delta < 0.5) &&
												(x <= max(p1.X, p2.X)) &&
												(x >= min(p1.X, p2.X)) &&
												(y <= max(p1.Y, p2.Y)) &&
												(y >= min(p1.Y, p2.Y))
				if curr_on_edge {
					s.SetContent(x, y, fill, nil, style)
					break
				} 
			}	
      if prev_on_edge && !curr_on_edge {
        horizontal_intersections = append(horizontal_intersections, x)
      }
      prev_on_edge = curr_on_edge
		}
    for i := 0; (i + 1) < len(horizontal_intersections); i += 2 {
      for x := horizontal_intersections[i]; x < horizontal_intersections[i + 1]; x++ {
        s.SetContent(x, y, fill, nil, style)
      }
    }
  horizontal_intersections = horizontal_intersections[:0]
	}
}
