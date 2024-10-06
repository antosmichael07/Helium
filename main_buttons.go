package main

import rl "github.com/gen2brain/raylib-go/raylib"

func (he *heightmap_editor) button_new_file() {
	name, err := he.popup_string("New Project Name:")
	if err != nil {
		if err.Error() != "esc" {
			rl.EndDrawing()
			he.popup_error(err.Error())
		}
		return
	}
	w, err := he.popup_uint("Heightmap Width:")
	if err != nil {
		if err.Error() != "esc" {
			rl.EndDrawing()
			he.popup_error(err.Error())
		}
		return
	}
	h, err := he.popup_uint("Heightmap Height:")
	if err != nil {
		if err.Error() != "esc" {
			rl.EndDrawing()
			he.popup_error(err.Error())
		}
		return
	}
	l, err := he.popup_uint("Heightmap Length:")
	if err != nil {
		if err.Error() != "esc" {
			rl.EndDrawing()
			he.popup_error(err.Error())
		}
		return
	}

	he.new_file(name, w, h, l)
}

func (he *heightmap_editor) button_open_file() {
	file, err := he.popup_string("Open File:")
	if err != nil {
		if err.Error() != "esc" {
			rl.EndDrawing()
			he.popup_error(err.Error())
		}
		return
	}

	he.load_heightmap_texture_image(file)
}

func (he *heightmap_editor) button_save_file() {
	err := he.save_heightmap_texture_image(he.project_name + ".mhm")

	if err != nil {
		rl.EndDrawing()
		he.popup_error(err.Error())
	} else {
		rl.EndDrawing()
		he.popup_alert("File saved successfully!")
	}
}
