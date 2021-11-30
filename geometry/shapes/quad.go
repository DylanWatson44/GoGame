package shapes

import "example/my-game/geometry/vector"

type Quad struct {
	TopLeft, TopRight, BottomLeft, BottomRight vector.Vector
}
