package main

import(
	"fmt"
	"os"
	"math"
  "time"
	"glitch/pkgs/render"
	"github.com/gdamore/tcell/v2"
)

type Circle struct {
  radius, center_x, center_y float64
  resolution int
}

func (c Circle) vertices() (out []render.Point) {
  step := math.Pi * 2 / float64(c.resolution)
  out = make([]render.Point, c.resolution)
  for i := 0 ; i < c.resolution; i++ {
    x := c.radius * math.Cos(step * float64(i)) + c.center_x
    y := c.radius * math.Sin(step * float64(i)) + c.center_y
    out[i] = render.Point{X: int(x), Y: int(y)}
  }
  return out
}

func main() {
	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	s, err := tcell.NewScreen()
	if err != nil {
			fmt.Println("Screen not found")
			os.Exit(0)
	}
	
	if err := s.Init(); err != nil {
			fmt.Println("Screen not initialised")
	}
	defer s.Fini()
  
  w, h := s.Size()
	s.SetStyle(defStyle)
  c := Circle{  radius: 5,
                center_x: float64(w) / 2,
                center_y: float64(h) / 2,
                resolution: 6,
              }

	go func () {
		for ;; {
			ev := s.PollEvent()
			switch ev := ev.(type) {
				case *tcell.EventKey:
					switch ev.Key() {
						case tcell.KeyCtrlC:
							os.Exit(0)
            case tcell.KeyUp:
              c.resolution += 1
            case tcell.KeyDown:
              c.resolution -= 1
					}
			}
		}
	}()
	
  go func () {
    scaling_factor := 10.0
    for i := 0.0;; i += 0.05 {
      c.radius = scaling_factor * math.Cos(i)
      if i >= math.Pi * 2 {
        i = 0
      }
      time.Sleep(time.Millisecond * 10)
    }
  }()

	for ;; {
		s.Clear()
    render.Render(c.vertices(), true, s, defStyle)
		s.Show()
	}
}
