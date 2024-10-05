package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	texture_icon_file_new = iota
)

func (he *heightmap_editor) init_textures() {
	he.textures = []rl.Texture2D{
		rl.LoadTexture("./textures/icon_file_new.png"),
	}
}
