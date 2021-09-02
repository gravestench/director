package main

import "image/color"

type Player int

const (
	empty Player = iota - 1
	PlayerX
	PlayerO
	numPlayers
)

func (p Player) String() string {
	switch p {
	case PlayerX:
		return "X"
	case PlayerO:
		return "O"
	default:
		return ""
	}
}

func (p Player) Color() color.Color {
	switch p {
	case PlayerX:
		return color.RGBA{R:0xFF, G: 0xFF, A: 0xFF}
	case PlayerO:
		return color.RGBA{R:0xFF, B: 0xFF, A: 0xFF}
	default:
		return color.RGBA{}
	}
}
