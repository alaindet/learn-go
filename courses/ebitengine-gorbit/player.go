package main

import (
	"gorbit/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	sprite *ebiten.Image
}

func NewPlayer(game *Game) *Player {
	return &Player{
		sprite: assets.PlayerSprite,
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}
	screen.DrawImage(p.sprite, options)
}
