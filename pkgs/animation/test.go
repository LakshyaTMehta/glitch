package animation

import (
	"time"
)

type Element interface {
	PosWhen( t time.Duration )( x, y float64 )
}

