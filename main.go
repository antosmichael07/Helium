package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	init_window()
	heightmap_editor := init_heightmap_editor()

	for {
		window_maganer()

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		heightmap_editor.tool_window.draw(&heightmap_editor.config.color_config)

		rl.DrawFPS(10, 10)
		rl.EndDrawing()
	}
}
