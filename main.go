package main

import(
	"fmt"
	"os"
	"time"
	"math"
	"glitch/pkgs/animation"
	"github.com/gdamore/tcell/v2"
)

var startTime time.Time

type CircleTracer struct {
	centreX, centreY, radius, speed, initialAngle float64
}

func (c CircleTracer) PosWhen( t time.Duration ) (x, y float64) {
	x = c.radius * math.Cos( t.Seconds() * 2 * math.Pi * c.speed + c.initialAngle) + c.centreX
	y = c.radius * math.Sin( t.Seconds() * 2 * math.Pi * c.speed + c.initialAngle) + c.centreY
	return 
}

func draw( s tcell.Screen, e animation.Element ) {
	_, h := s.Size()
	x, y := e.PosWhen(time.Since(startTime))
	s.SetContent( int(x), h - int(y), tcell.RuneBlock, nil, tcell.StyleDefault)
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

	s.SetStyle(defStyle)

	//point := NewParticle(0, 0, 0, 0, 1, 0.8)
	//point.Push(Vector2D{10, 10})

	c := CircleTracer{ 80, 20, 18, 1, 0 }
	c1 := CircleTracer{ 80, 20, 18, 1, 0.2 }
	c2 := CircleTracer{ 80, 20, 18, 1, 0.4 }
	c3 := CircleTracer{ 80, 20, 18, 1, 0.6 }

	go func () {
		for ;; {
			ev := s.PollEvent()
			switch ev := ev.(type) {
				case *tcell.EventKey:
					switch ev.Key() {
						case tcell.KeyCtrlC:
							os.Exit(0)
						case tcell.KeyUp:
							c.speed += .1
						case tcell.KeyDown:
							c.speed -= .1
					}
			}
		}
	}()
	
	startTime = time.Now()
	for ;; {
		s.Clear()
		draw(s, c)
		draw(s, c1)
		draw(s, c2)
		draw(s, c3)
		s.Show()
	}
}
