package player

import (
	"example/my-game/actor"
	"example/my-game/mymathutil"
	"example/my-game/sprite"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
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

type Player struct {
	Actor actor.Actor
}

func InstantiatePlayer() *Player {

	var PlayerSpritesMap map[int32]*sprite.Sprite
	PlayerSpritesMap = make(map[int32]*sprite.Sprite)

	PlayerSpritesMap[Idle] = sprite.NewSprite("Cyborg/Cyborg_idle.png", 48, 48, 4, 1, 0, 0, 5)
	PlayerSpritesMap[Run] = sprite.NewSprite("Cyborg/Cyborg_run.png", 48, 48, 6, 1, 0, 0, 5)

	var player Player = Player{Actor: actor.Actor{
		State:         Idle,
		Sprites:       PlayerSpritesMap,
		CurrentSprite: PlayerSpritesMap[Idle],
		X:             100,
		Y:             100,
		MaxAccel:      4,
		MaxVelocity:   10,
		Facing:        actor.Right,
		Stopping:      false,
		StopSpeed:     0.15,
	}}

	return &player
}

func (player *Player) Update() {
	player.checkInput()
	player.updatePhysics()
	player.Actor.ManageSubRoutines()
}

func (player *Player) checkInput() {
	requestToMoveLeftOrRight := ebiten.IsKeyPressed(ebiten.KeyArrowRight) || ebiten.IsKeyPressed(ebiten.KeyArrowLeft)
	if requestToMoveLeftOrRight {
		player.start()
	} else {
		player.stop()
	}
}

func (player *Player) updatePhysics() {
	// if player.Actor.Stopping {
	// 	player.Actor.AccelX *= (1 - player.Actor.StopSpeed)
	// 	if math.Abs(player.Actor.AccelX) <= 0.01 {
	// 		player.Actor.Stopping = false
	// 		player.Actor.AccelX = 0
	// 	}
	// }
	player.Actor.VelocityX += player.Actor.AccelX
	player.Actor.VelocityX = mymathutil.ClampFloat64(player.Actor.VelocityX, -player.Actor.MaxVelocity, player.Actor.MaxVelocity)
	player.Actor.X += player.Actor.VelocityX

	player.Actor.VelocityY += player.Actor.AccelY
	player.Actor.VelocityY = mymathutil.ClampFloat64(player.Actor.VelocityY, -player.Actor.MaxVelocity, player.Actor.MaxVelocity)
	player.Actor.Y += player.Actor.VelocityY
}

func (player *Player) start() {
	if player.Actor.State != Run {
		player.Actor.Stopping = false
		player.updateState(Run)
		if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
			player.Actor.VelocityX = 2
			player.Actor.AccelX = 1
			player.Actor.Facing = actor.Right
		} else {
			player.Actor.VelocityX = -2
			player.Actor.AccelX = -1
			player.Actor.Facing = actor.Left
		}
	}
}

func (player *Player) stop() {
	if player.Actor.State != Idle {
		player.updateState(Idle)
		player.Actor.Stopping = true
		player.Actor.AccelX = 0
		// player.Actor.VelocityX = 0
		stopRoutineChannel := comeToAStop(&player.Actor)
		player.Actor.SubRoutines = append(player.Actor.SubRoutines, stopRoutineChannel)

		// if player.Actor.AccelX > 0 {
		// 	player.Actor.AccelX = -1
		// } else if player.Actor.AccelX < 0 {
		// 	player.Actor.AccelX = 1
		// }
	}
}

func comeToAStop(actor *actor.Actor) <-chan bool {
	isDoneChannel := make(chan bool)

	go func() {
		for actor.Stopping {
			actor.VelocityX *= (1 - actor.StopSpeed)
			if math.Abs(actor.VelocityX) <= 1 {
				actor.Stopping = false
				actor.VelocityX = 0
				isDoneChannel <- true
				//			close(isDoneChannel)
			} else {
				isDoneChannel <- false
			}
		}
		isDoneChannel <- true
		close(isDoneChannel)
	}()

	return isDoneChannel
}

func (player *Player) updateState(state int32) {
	player.Actor.State = state
	player.Actor.CurrentSprite = player.Actor.Sprites[state]
	player.Actor.CurrentFrame = 0
}
