package player

import (
	"example/my-game/actor"
	"example/my-game/geometry/shapes"
	"example/my-game/geometry/vector"
	"example/my-game/mymathutil"
	"example/my-game/sprite"
	"fmt"
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
	Actor   actor.Actor
	gravity float64
}

func InstantiatePlayer() *Player {

	var PlayerSpritesMap map[int32]*sprite.Sprite
	PlayerSpritesMap = make(map[int32]*sprite.Sprite)

	PlayerSpritesMap[Idle] = sprite.NewSprite("Cyborg/Cyborg_idle.png", 48, 48, 4, 1, 26, 14, 5)
	PlayerSpritesMap[Run] = sprite.NewSprite("Cyborg/Cyborg_run.png", 48, 48, 6, 1, 26, 14, 5)

	var player Player = Player{
		Actor: actor.Actor{
			State:         Idle,
			Sprites:       PlayerSpritesMap,
			CurrentSprite: PlayerSpritesMap[Idle],
			Position: vector.Vector{
				X: 100,
				Y: 100},
			Velocity: vector.Vector{
				X: 0,
				Y: 0},
			MaxAccel:    4,
			MaxVelocity: 10,
			Facing:      actor.Right,
			Stopping:    false,
			StopSpeed:   0.15,
		},
		gravity: 0,
	}

	return &player
}

func (player *Player) Update(aHitbox shapes.Quad) {
	player.checkInput()
	player.updatePhysics(aHitbox)
	player.Actor.ManageSubRoutines()
}

func (player *Player) checkInput() {
	requestToMoveLeftOrRight := ebiten.IsKeyPressed(ebiten.KeyArrowRight) || ebiten.IsKeyPressed(ebiten.KeyArrowLeft)
	if requestToMoveLeftOrRight {
		player.startWalking()
	} else {
		player.stopWalking()
	}
}

func (player *Player) updatePhysics(aHitbox shapes.Quad) {
	player.Actor.Velocity.X += player.Actor.Acceleration.X
	player.Actor.Velocity.X = mymathutil.ClampFloat64(player.Actor.Velocity.X, -player.Actor.MaxVelocity, player.Actor.MaxVelocity)

	integer, decimal := math.Modf(player.Actor.Velocity.X)
	interactions := integer
	if decimal != 0 {
		interactions++
	}

	direction := mymathutil.ReduceToSignedUnit(integer)
	for i := 0; i < int(math.Abs(player.Actor.Velocity.X)); i++ {
		var amountToMove float64
		if decimal != 0 {
			if i == 0 {
				amountToMove = decimal
			}
		} else {
			amountToMove = float64(direction)
		}

		var nextXposition = player.Actor.Position.X + amountToMove

		newPosition := vector.Vector{X: nextXposition, Y: player.Actor.Position.Y}
		if !IsColliding(player.Actor.GetHitBoxAt(newPosition), aHitbox) {
			player.Actor.Position.X = nextXposition
		} else {
			//moveback
			fmt.Println("collided")
			player.stopWalking() //replace
			break
		}
	}

	player.Actor.Velocity.Y += player.Actor.Acceleration.Y + player.gravity
	player.Actor.Velocity.Y = mymathutil.ClampFloat64(player.Actor.Velocity.Y, -player.Actor.MaxVelocity, player.Actor.MaxVelocity)

	integerY, decimalY := math.Modf(player.Actor.Velocity.Y)
	interactionsY := integerY
	if decimalY != 0 {
		interactionsY++
	}

	directionY := mymathutil.ReduceToSignedUnit(integerY)
	for i := 0; i < int(math.Abs(player.Actor.Velocity.Y)); i++ {
		var amountToMove float64
		if decimalY != 0 {
			if i == 0 {
				amountToMove = decimalY
			}
		} else {
			amountToMove = float64(directionY)
		}

		var nextYposition = player.Actor.Position.Y + amountToMove

		newPositionY := vector.Vector{X: player.Actor.Position.X, Y: nextYposition}
		if !IsColliding(player.Actor.GetHitBoxAt(newPositionY), aHitbox) {
			player.Actor.Position.Y = nextYposition
		} else {
			//moveback
			fmt.Println("collided")
			player.Actor.Velocity.Y = 0
			player.stopWalking() //replace
			break
		}
	}

}

// func (player *Player) changePos(valToChange *float64, velocity float64, aHitbox shapes.Quad) {
// 	integer, decimal := math.Modf(velocity)
// 	interactions := integer
// 	if decimal != 0 {
// 		interactions++
// 	}

// 	direction := mymathutil.ReduceToSignedUnit(integer)
// 	for i := 0; i < int(math.Abs(velocity)); i++ {
// 		var amountToMove float64
// 		if decimal != 0 {
// 			if i == 0 {
// 				amountToMove = decimal
// 			}
// 		} else {
// 			amountToMove = float64(direction)
// 		}

// 		var nextXposition = *valToChange + amountToMove

// 		newPosition := vector.Vector{X: nextXposition, Y: player.Actor.Position.Y}
// 		if !IsColliding(player.Actor.GetHitBoxAt(newPosition), aHitbox) {
// 			*valToChange = nextXposition
// 		} else {
// 			//moveback
// 			fmt.Println("collided")
// 			player.stopWalking() //replace
// 			break
// 		}
// 	}
// }

func IsColliding(hitboxA, hitboxB shapes.Quad) bool {
	return CheckCollision(hitboxA, hitboxB) || CheckCollision(hitboxB, hitboxA)
}

func CheckCollision(hitboxA, hitboxB shapes.Quad) bool {
	listOfPoints := [4]vector.Vector{hitboxA.TopLeft, hitboxA.TopRight, hitboxA.BottomLeft, hitboxA.BottomRight}

	for _, point := range listOfPoints {
		if point.X >= hitboxB.TopLeft.X && point.X <= hitboxB.TopRight.X && point.Y >= hitboxB.TopLeft.Y && point.Y <= hitboxB.BottomRight.Y {
			return true
		}
	}

	return false
}

//fix switch directions
func (player *Player) startWalking() {
	if player.Actor.State != Run {
		player.Actor.Stopping = false
		player.updateState(Run)
		if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
			player.Actor.Velocity.X = 2
			player.Actor.Acceleration.X = 1
			player.Actor.Facing = actor.Right
		} else {
			player.Actor.Velocity.X = -2
			player.Actor.Acceleration.X = -1
			player.Actor.Facing = actor.Left
		}
	}
}

func (player *Player) stopWalking() {
	if player.Actor.State != Idle {
		player.updateState(Idle)
		player.Actor.Stopping = true
		player.Actor.Acceleration.X = 0
		player.Actor.Acceleration.Y = 0
		// player.Actor.Velocity.X = 0
		stopRoutineChannel := comeToAStop(&player.Actor)
		player.Actor.SubRoutines = append(player.Actor.SubRoutines, stopRoutineChannel)

		// if player.Actor.Acceleration.X > 0 {
		// 	player.Actor.Acceleration.X = -1
		// } else if player.Actor.Acceleration.X < 0 {
		// 	player.Actor.Acceleration.X = 1
		// }
	}
}

func comeToAStop(actor *actor.Actor) <-chan bool {
	isDoneChannel := make(chan bool)

	go func() {
		for actor.Stopping {
			actor.Velocity.X *= (1 - actor.StopSpeed)
			if math.Abs(actor.Velocity.X) <= 1 {
				actor.Stopping = false
				actor.Velocity.X = 0
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
