package actor

import (
	"example/my-game/geometry/shapes"
	"example/my-game/geometry/vector"
	"example/my-game/sprite"

	"github.com/hajimehoshi/ebiten/v2"
)

// type ActorI interface {
// 	initialize()
// }

type FacingDirection int32

const (
	Left FacingDirection = iota
	Right
)

type Actor struct {
	State                            int32
	Sprites                          map[int32]*sprite.Sprite
	CurrentSprite                    *sprite.Sprite
	Facing                           FacingDirection
	CurrentFrame                     int32
	Position, Velocity, Acceleration vector.Vector
	MaxAccel, MaxVelocity            float64
	Stopping                         bool
	StopSpeed                        float64
	SubRoutines                      []<-chan bool
}

func (actor *Actor) ManageSubRoutines() {
	i := 0
	for _, subroutine := range actor.SubRoutines {
		subroutineIsDone := <-subroutine
		if !subroutineIsDone {
			// if its not done, keep it
			actor.SubRoutines[i] = subroutine
			i++
		}
	}
	// Prevent memory leak by erasing truncated values
	for j := i; j < len(actor.SubRoutines); j++ {
		actor.SubRoutines[j] = nil
	}
	actor.SubRoutines = actor.SubRoutines[:i]
}

func (actor *Actor) Draw(screen *ebiten.Image, currentGameFrame int) {
	options := &ebiten.DrawImageOptions{}
	if actor.Facing == Left {
		options.GeoM.Scale(-1, 1)
	}
	var imgObj *ebiten.Image = actor.CurrentSprite.GetFrame(currentGameFrame)

	width, height := imgObj.Size()
	options.GeoM.Translate(-float64(width)/2, -float64(height)/2)
	options.GeoM.Translate(actor.Position.X, actor.Position.Y)


	// options.ColorM.Scale(1, 1, 1, 0)
	// // Set color
	// options.ColorM.Translate(0, 0, 0, 1)

	screen.DrawImage(imgObj, options)
}

func (actor *Actor) GetHitBoxAt(position vector.Vector) shapes.Quad {
	var horizontalOffset float64 = float64(actor.CurrentSprite.FrameWidth-actor.CurrentSprite.HotizontalPadding) / 2
	var verticalOffset float64 = float64(actor.CurrentSprite.FrameHeight-actor.CurrentSprite.VerticalPadding) / 2

	topLeft := vector.Vector{X: (position.X - horizontalOffset), Y: (position.Y - verticalOffset)}

	topRight := vector.Vector{X: (position.X + horizontalOffset), Y: (position.Y - verticalOffset)}

	bottomLeft := vector.Vector{X: (position.X - horizontalOffset), Y: (position.Y + verticalOffset)}

	bottomRight := vector.Vector{X: (position.X + horizontalOffset), Y: (position.Y + verticalOffset)}

	hitbox := shapes.Quad{
		TopLeft:     topLeft,
		TopRight:    topRight,
		BottomLeft:  bottomLeft,
		BottomRight: bottomRight,
	}

	return hitbox
}
