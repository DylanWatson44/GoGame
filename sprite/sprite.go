package sprite

import (
	"fmt"
	"image"
	_ "image/png"
	"log"
	"path/filepath"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Sprite struct {
	SpriteSheet                        *ebiten.Image
	FrameWidth, FrameHeight            int
	numFrames                          int
	rows                               int
	hotizontalPadding, verticalPadding int
	frameSpeed                         int
}

func NewSprite(spriteSheetFileLocation string, frameWidth, frameHeight, numFrames, rows, hotizontalPadding, verticalPadding, frameSpeed int) *Sprite {
	newSprite := new(Sprite)
	var err error
	var img *ebiten.Image

	assetFilePath := filepath.Join("assets", spriteSheetFileLocation)
	img, _, err = ebitenutil.NewImageFromFile(assetFilePath)
	if err != nil {
		log.Fatal("Error making image ", err)
	} else {
		fmt.Println("loaded the image")
	}
	newSprite.SpriteSheet = img
	newSprite.FrameHeight = frameHeight
	newSprite.FrameWidth = frameWidth
	newSprite.numFrames = numFrames
	newSprite.rows = rows
	newSprite.hotizontalPadding = hotizontalPadding
	newSprite.verticalPadding = verticalPadding
	newSprite.frameSpeed = frameSpeed

	return newSprite
}

func (sprite Sprite) GetFrame(currentGameFrame int) *ebiten.Image {
	currentFrame := (currentGameFrame / sprite.frameSpeed) % sprite.numFrames

	currentFrameX := currentFrame * sprite.FrameWidth
	currentFrameY := 0
	var currentFrameRect image.Rectangle = image.Rect(currentFrameX, currentFrameY, currentFrameX+sprite.FrameWidth, currentFrameY+sprite.FrameHeight)
	var subImg image.Image = sprite.SpriteSheet.SubImage(currentFrameRect)
	var imgObj *ebiten.Image = subImg.(*ebiten.Image)

	return imgObj
}

//TODO make customer frame controller, i.e. allow some frames to have longer frame times than others
