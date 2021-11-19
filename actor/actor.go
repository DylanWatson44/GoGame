package actor

import (
	"example/my-game/sprite"

	"github.com/hajimehoshi/ebiten/v2"
)

// type ActorI interface {
// 	initialize()
// }

type Actor struct {
	State                int32
	Sprites              map[int32]*sprite.Sprite
	CurrentSprite        *sprite.Sprite
	currentFrame         int32
	X, Y                 float64
	VelocityX, VelocityY float64
	AccelX, AccelY       float64
}

func (actor Actor) Draw(screen *ebiten.Image, currentGameFrame int) {
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(-float64(actor.CurrentSprite.FrameWidth)/2, -float64(actor.CurrentSprite.FrameHeight)/2)
	options.GeoM.Translate(actor.X, actor.Y)

	var imgObj *ebiten.Image = actor.CurrentSprite.GetFrame(currentGameFrame)

	screen.DrawImage(imgObj, options)
}
