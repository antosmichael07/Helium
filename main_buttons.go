package main

func (he *heightmap_editor) button_new_file() {
	name, esc := he.popup_string("New Project Name:")
	if esc {
		return
	}
	x, esc := he.popup_uint("Heightmap Width:")
	if esc {
		return
	}
	y, esc := he.popup_uint("Heightmap Height:")
	if esc {
		return
	}

	he.new_file(name, x, y)
}

func (he *heightmap_editor) button_open_file() {
	file, esc := he.popup_string("Open File:")
	if esc {
		return
	}

	he.load_heightmap_texture_image(file)
}

func (he *heightmap_editor) button_save_file() {
	err := he.save_heightmap_texture_image(he.project_name + ".mhm")

	if err != nil {
		he.popup_alert(err.Error())
	} else {
		he.popup_alert("File saved successfully!")
	}
}
