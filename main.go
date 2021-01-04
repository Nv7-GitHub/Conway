package main

import (
	"strconv"

	r "github.com/lachee/raylib-goplus/raylib"
)

const fps = 60
const width = 10
const height = 10

func main() {
	b := newBoard(width, height)
	b.Board[2][0] = true
	b.Board[2][2] = true
	b.Board[1][2] = true
	b.Board[3][1] = true
	b.Board[3][2] = true
	b.Width = 500
	b.Height = 500
	b.Y = 30
	gameStarted := false
	timeSince := 0
	speed := 60

	r.InitWindow(500, 530, "Conway's Game Of Life")
	r.SetTargetFPS(fps)
	for !r.WindowShouldClose() {
		r.SetMouseScale(1, 1)
		r.BeginDrawing()

		r.ClearBackground(r.Black)
		b.draw()
		if !gameStarted {
			b.events()
		} else if timeSince >= speed {
			b.generation()
			timeSince = 0
		}

		if gameStarted {
			timeSince++
		}

		text := "Stop!"
		if !gameStarted {
			text = "Start!"
		}
		if r.GuiButton(r.Rectangle{X: 0, Y: 0, Width: 100, Height: 30}, text) {
			gameStarted = !gameStarted
		}
		speed = int(fps - r.GuiSlider(r.Rectangle{X: 150, Y: 0, Width: 100, Height: 30}, "Speed: 0", "1", float32(fps-speed)/fps, 0, 1)*fps)

		r.DrawText(strconv.Itoa(r.GetFPS()), 0, 500, 30, r.RayWhite)

		r.EndDrawing()
	}
	r.CloseWindow()
}
