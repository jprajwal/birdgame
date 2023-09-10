package animation

import (
	"github.com/jprajwal/birdgame/space"
)

type Animation2D interface {
	Animate(space.Object2D, space.Space2D)
	Done() bool
}
