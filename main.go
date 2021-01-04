package main

import r "github.com/lachee/raylib-goplus/raylib"

func main() {
	b := newBoard(10, 10)
	b.Board[2][0] = true
	b.Board[2][2] = true
	b.Board[1][2] = true
	b.Board[3][1] = true
	b.Board[3][2] = true
	b.Width = 500
	b.Height = 500
	b.Y = 30
	gameStarted := false

	r.InitWindow(500, 530, "Conway's Game Of Life")
	for !r.WindowShouldClose() {
		r.SetMouseScale(1, 1)
		r.BeginDrawing()

		r.ClearBackground(r.Black)
		b.draw()
		if !gameStarted {
			b.events()
		}

		if r.GuiButton(r.Rectangle{X: 0, Y: 0, Width: 100, Height: 30}, "Start!") {
			gameStarted = true
		}

		r.EndDrawing()
	}
	r.CloseWindow()
}
