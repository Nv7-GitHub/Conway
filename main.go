package main

import r "github.com/lachee/raylib-goplus/raylib"

func main() {
	r.InitWindow(800, 450, "Raylib Go Plus")
	for !r.WindowShouldClose() {
		r.BeginDrawing()
		r.ClearBackground(r.RayWhite)
		r.DrawText("Woo! Raylib-Go-Plus! Now with ++", 20, 20, 20, r.GopherBlue)
		r.EndDrawing()
	}
	r.CloseWindow()
}
