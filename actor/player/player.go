package player

import (
	. "example/my-game/actor"
	"example/my-game/sprite"
)

const (
	Unset = iota
	Idle
	Attack1
	Attack2
	Attack3
	Climb
	Death
	DoubleJump
	Hurt
	Jump
	Punch
	Run
	RunAttack
)

func InstantiatePlayer() *Actor {

	var PlayerSpritesMap map[int32]*sprite.Sprite
	PlayerSpritesMap = make(map[int32]*sprite.Sprite)

	PlayerSpritesMap[Idle] = sprite.NewSprite("Cyborg/Cyborg_idle.png", 48, 48, 4, 1, 0, 0, 5)
	PlayerSpritesMap[Run] = sprite.NewSprite("Cyborg/Cyborg_run.png", 48, 48, 6, 1, 0, 0, 5)

	var player Actor = Actor{
		State:         Idle,
		Sprites:       PlayerSpritesMap,
		CurrentSprite: PlayerSpritesMap[Idle],
		X:             100,
		Y:             100,
	}

	return &player
}
