package main

import(
	"fmt"
	"os"
	//"math"
  //"time"
	//"glitch/pkgs/render"
	"glitch/pkgs/face"
	//"glitch/pkgs/shapes"
	"github.com/gdamore/tcell/v2"
)

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
  
  left_eye_xpos := w / 4
  left_eye_ypos := h / 3
  left_eye_width := w / 4
  left_eye_height := h / 2

  right_eye_xpos := w / 4 * 3
  right_eye_ypos := left_eye_ypos
  right_eye_width := left_eye_width
  right_eye_height := left_eye_height

	s.SetStyle(defStyle)
  left_eye := face.NewEye(  left_eye_xpos,
                            left_eye_ypos,
                            left_eye_width,
                            left_eye_height,
                          )
  right_eye := face.NewEye( right_eye_xpos,
                            right_eye_ypos,
                            right_eye_width,
                            right_eye_height,
                          )

	go func () {
		for ;; {
			ev := s.PollEvent()
			switch ev := ev.(type) {
				case *tcell.EventKey:
					switch ev.Key() {
						case tcell.KeyCtrlC:
							os.Exit(0)
            case tcell.KeyUp:
            case tcell.KeyDown:
					}
			}
		}
	}()
	
  go left_eye.Watchout()
  go right_eye.Watchout()

	for ;; {
		s.Clear()
    left_eye.Draw(s, defStyle)
    left_eye.Update()
    right_eye.Update()
    right_eye.Draw(s, defStyle)
		s.Show()
	}
}
