package main

import (
	"math"

	r "github.com/lachee/raylib-goplus/raylib"
)

type board struct {
	Board        [][]bool
	WasMouseDown bool
	Width        int
	Height       int
	X            int
	Y            int
}

func newBoard(w, h int) board {
	b := make([][]bool, h)
	for i := 0; i < len(b); i++ {
		b[i] = make([]bool, w)
	}
	return board{
		Board: b,
	}
}

func (b *board) draw() {
	height := b.Width / len(b.Board)
	width := b.Height / len(b.Board[0])
	for y := 0; y < len(b.Board); y++ {
		for x := 0; x < len(b.Board[y]); x++ {
			if b.Board[y][x] {
				r.DrawRectangle(x*width+b.X, y*height+b.Y, width, height, r.RayWhite)
			}
		}
	}
}

func (b *board) events() {
	if r.IsMouseButtonDown(r.MouseLeftButton) && (!b.WasMouseDown) {
		x := float64(r.GetMouseX() - b.X)
		y := float64(r.GetMouseY() - b.Y)
		if int(x) <= b.Width && int(y) <= b.Height && (int(x) >= 0) && (int(y) >= 0) {
			height := float64(b.Height / len(b.Board))
			width := float64(b.Width / len(b.Board[0]))
			posX := int(math.Floor(x / width))
			posY := int(math.Floor(y / height))
			b.Board[posY][posX] = !b.Board[posY][posX]
			b.WasMouseDown = true
		}
	}
	b.WasMouseDown = r.IsMouseButtonDown(r.MouseLeftButton)
}
