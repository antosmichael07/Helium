package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	init_window()
	heightmap_editor := init_heightmap_editor()

	for {
		window_maganer()

		heightmap_editor.tool_window.update()
		heightmap_editor.buttons.update(button_group_main)
		heightmap_editor.buttons.last_update()

		rl.BeginDrawing()
		rl.ClearBackground(heightmap_editor.config.color_config.preview_background)

		rl.DrawRectangle(0, 0, int32(rl.GetScreenWidth()), 30, heightmap_editor.config.color_config.button)
		heightmap_editor.tool_window.draw(&heightmap_editor.config.color_config)
		heightmap_editor.buttons.draw(button_group_main, &heightmap_editor.config.color_config)

		heightmap_editor.draw_project_info()
		rl.EndDrawing()
	}
}

func (he *heightmap_editor) draw_project_info() {
	text := fmt.Sprintf("Project Name: %s, Width: %d, Height: %d", he.project_name, he.heightmap_width, he.heightmap_height)
	rl.DrawText(text, int32(rl.GetScreenWidth())-rl.MeasureText(text, 20)-10, 5, 20, he.config.color_config.text)
}
