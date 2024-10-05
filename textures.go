package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	texture_icon_file_new = iota
	texture_icon_file_open
	texture_icon_file_save
)

func (he *heightmap_editor) init_textures() {
	he.textures = []rl.Texture2D{
		rl.LoadTexture("./textures/icon_file_new.png"),
		rl.LoadTexture("./textures/icon_file_open.png"),
		rl.LoadTexture("./textures/icon_file_save.png"),
	}
}
