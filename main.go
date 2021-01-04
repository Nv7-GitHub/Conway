package main

import (
	"strconv"

	r "github.com/lachee/raylib-goplus/raylib"
)

const fps = 60

// WIDTH is the width of the window
const WIDTH = 650

// HEIGHT is the height of the window
const HEIGHT = 650

func main() {
	size := 10
	b := newBoard(size, size)
	b.Board[2][0] = true
	b.Board[2][2] = true
	b.Board[1][2] = true
	b.Board[3][1] = true
	b.Board[3][2] = true
	b.Width = WIDTH
	b.Height = HEIGHT
	b.Y = 30

	gameStarted := false
	timeSince := 0
	speed := 60

	nS := 10
	nSt := 0
	nSo := 10

	r.InitWindow(WIDTH, HEIGHT+30, "Conway's Game Of Life")
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

		timeSince++

		text := "Stop!"
		if !gameStarted {
			text = "Start!"
		}
		if r.GuiButton(r.Rectangle{X: 0, Y: 0, Width: 100, Height: 30}, text) {
			gameStarted = !gameStarted
		}
		speed = int(fps - r.GuiSlider(r.Rectangle{X: 150, Y: 0, Width: 100, Height: 30}, "Speed: 0", "1", float32(fps-speed)/fps, 0, 1)*fps)

		if !gameStarted {
			nS = int(r.GuiSlider(r.Rectangle{X: 310, Y: 0, Width: 100, Height: 30}, "Size: 5", "100", float32(nS), 5, 100))

			if nS != nSo {
				nSt = timeSince
				nSo = nS
			}

			if nS != size && (timeSince-nSt) > fps/2 {
				newBoard := newBoard(nS, nS)
				for y := 0; y < len(b.Board); y++ {
					for x := 0; x < len(b.Board[y]); x++ {
						if (y >= 0) && (y < len(newBoard.Board)) && (x >= 0) && (x < len(newBoard.Board[y])) {
							if b.Board[y][x] {
								newBoard.Board[y][x] = true
							}
						} else {
							break
						}
					}
				}
				b = newBoard
				size = nS
				nSo = nS
				b.Width = WIDTH
				b.Height = HEIGHT
				b.Y = 30
			}
		}

		r.DrawText(strconv.Itoa(r.GetFPS()), 0, HEIGHT, 30, r.RayWhite)

		r.EndDrawing()
	}
	r.CloseWindow()
}
