package main

import (
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type heightmap_editor struct {
	tool_window      tool_window
	buttons          buttons
	heightmap_image  *rl.Image
	texture_image    *rl.Image
	heightmap_width  int
	heightmap_height int
	project_name     string
	config           config
	textures         []rl.Texture2D
}

func init_heightmap_editor() heightmap_editor {
	he := heightmap_editor{}

	he.init_textures()
	he.init_tool_window()
	he.init_buttons()
	he.init_config()
	he.new_file("New Project", 16, 16)

	return he
}

func (he *heightmap_editor) new_file(name string, width, height int) {
	he.project_name = name
	he.heightmap_image = rl.GenImageColor(width, height, rl.Black)
	he.texture_image = rl.GenImageColor(width, height, rl.White)
	he.heightmap_width = width
	he.heightmap_height = height
}

func (he *heightmap_editor) save_heightmap_texture_image(file string) error {
	heightmap_data := rl.ExportImageToMemory(*he.heightmap_image, ".png")
	texture_data := rl.ExportImageToMemory(*he.texture_image, ".png")

	data := []byte{byte(len(heightmap_data) >> 24), byte(len(heightmap_data) >> 16), byte(len(heightmap_data) >> 8), byte(len(heightmap_data)), byte(len(texture_data) >> 24), byte(len(texture_data) >> 16), byte(len(texture_data) >> 8), byte(len(texture_data))}
	data = append(data, heightmap_data...)
	data = append(data, texture_data...)

	return os.WriteFile(file, data, 0644)
}

func (he *heightmap_editor) load_heightmap_texture_image(file string) error {
	data, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	heightmap_length := int32(data[0])<<24 | int32(data[1])<<16 | int32(data[2])<<8 | int32(data[3])
	he.heightmap_image = rl.LoadImageFromMemory(".png", data[8:8+int32(heightmap_length)], int32(heightmap_length))
	texture_length := int32(data[4])<<24 | int32(data[5])<<16 | int32(data[6])<<8 | int32(data[7])
	he.texture_image = rl.LoadImageFromMemory(".png", data[8+int32(heightmap_length):8+int32(heightmap_length)+int32(texture_length)], int32(texture_length))

	return nil
}
