package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	init_window()
	heightmap_editor := init_heightmap_editor()

	heightmap_editor.buttons.new_group()

	heightmap_editor.buttons.new_button_texture(button_group_main, &heightmap_editor.textures[texture_icon_file_new], 0, 0, func() {})

	for {
		window_maganer()

		heightmap_editor.tool_window.update()
		heightmap_editor.buttons.update(button_group_main)
		heightmap_editor.buttons.last_update()

		rl.BeginDrawing()
		rl.ClearBackground(heightmap_editor.config.color_config.preview_background)

		heightmap_editor.tool_window.draw(&heightmap_editor.config.color_config)
		heightmap_editor.buttons.draw(button_group_main, &heightmap_editor.config.color_config)

		rl.EndDrawing()
	}
}
