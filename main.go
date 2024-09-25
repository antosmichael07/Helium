package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	init_window()

	for {
		window_maganer()

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		rl.DrawFPS(10, 10)
		rl.EndDrawing()
	}
}
