package actor

import (
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
	State                 int32
	Sprites               map[int32]*sprite.Sprite
	CurrentSprite         *sprite.Sprite
	Facing                FacingDirection
	CurrentFrame          int32
	X, Y                  float64
	VelocityX, VelocityY  float64
	AccelX, AccelY        float64
	MaxAccel, MaxVelocity float64
	Stopping              bool
	StopSpeed             float64
	SubRoutines           []<-chan bool
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
	options.GeoM.Translate(-float64(actor.CurrentSprite.FrameWidth)/2, -float64(actor.CurrentSprite.FrameHeight)/2)
	options.GeoM.Translate(actor.X, actor.Y)

	var imgObj *ebiten.Image = actor.CurrentSprite.GetFrame(currentGameFrame)

	screen.DrawImage(imgObj, options)
}
