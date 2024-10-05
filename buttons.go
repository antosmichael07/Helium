package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type buttons struct {
	group          [][]button
	group_pressed  int
	button_pressed int
	empty_texture  rl.Texture2D
	clicked_ok     bool
}

type button struct {
	text     string
	texture  *rl.Texture2D
	rect     rl.Rectangle
	function func()
}

const (
	button_group_main = iota
	button_group_popup
)

func (he *heightmap_editor) init_buttons() {
	he.buttons = buttons{}

	he.buttons.group = [][]button{}
	he.buttons.group_pressed = -1
	he.buttons.button_pressed = -1
	he.buttons.empty_texture = rl.LoadTextureFromImage(rl.GenImageColor(1, 1, rl.Black))

	he.buttons.new_group()
	he.buttons.new_group()

	he.init_main_buttons()
	he.buttons.new_button(button_group_popup, "OK", 0, 0, func() {
		he.buttons.clicked_ok = true
	})
}

func (bs *buttons) new_group() {
	bs.group = append(bs.group, []button{})
}

func (bs *buttons) new_button(g int, text string, x, y int, function func()) {
	bs.group[g] = append(bs.group[g], button{
		text:     text,
		texture:  &bs.empty_texture,
		rect:     rl.NewRectangle(float32(x), float32(y), float32(rl.MeasureText(text, 20)+10), float32(30)),
		function: function,
	})
}

func (bs *buttons) new_button_texture(g int, texture *rl.Texture2D, x, y int, function func()) {
	bs.group[g] = append(bs.group[g], button{
		text:     "",
		texture:  texture,
		rect:     rl.NewRectangle(float32(x), float32(y), float32(30), float32(30)),
		function: function,
	})
}

func (bs *buttons) update(g int) {
	if bs.group_pressed == -1 && rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		for i := 0; i < len(bs.group[g]); i++ {
			if rl.CheckCollisionPointRec(rl.GetMousePosition(), bs.group[g][i].rect) {
				bs.group_pressed = g
				bs.button_pressed = i
				break
			}
		}
	}

	if bs.group_pressed == g && rl.IsMouseButtonReleased(rl.MouseLeftButton) {
		if rl.CheckCollisionPointRec(rl.GetMousePosition(), bs.group[g][bs.button_pressed].rect) {
			bs.group[g][bs.button_pressed].function()
		}
	}
}

func (bs *buttons) last_update() {
	if rl.IsMouseButtonReleased(rl.MouseLeftButton) {
		bs.group_pressed = -1
		bs.button_pressed = -1
		bs.clicked_ok = false
	}
}

func (bs *buttons) draw(g int, cc *color_config) {
	for i := 0; i < len(bs.group[g]); i++ {
		rl.DrawRectangleRec(bs.group[g][i].rect, cc.button)
		if rl.CheckCollisionPointRec(rl.GetMousePosition(), bs.group[g][i].rect) {
			if bs.group_pressed == g && bs.button_pressed == i {
				rl.DrawRectangleRec(bs.group[g][i].rect, cc.button_click)
			} else if !rl.IsMouseButtonDown(rl.MouseLeftButton) {
				rl.DrawRectangleRec(bs.group[g][i].rect, cc.button_hover)
			}
		}
		if bs.group[g][i].texture == &bs.empty_texture {
			rl.DrawText(bs.group[g][i].text, int32(bs.group[g][i].rect.X)+5, int32(bs.group[g][i].rect.Y)+5, 20, cc.text)
		} else {
			rl.DrawTexture(*bs.group[g][i].texture, int32(bs.group[g][i].rect.X)+5, int32(bs.group[g][i].rect.Y)+5, rl.White)
		}
	}
}

func (he *heightmap_editor) init_main_buttons() {
	he.buttons.new_button_texture(button_group_main, &he.textures[texture_icon_file_new], 0, 0, func() {
		he.button_new_file()
	})

	he.buttons.new_button_texture(button_group_main, &he.textures[texture_icon_file_open], 31, 0, func() {
		he.button_open_file()
	})

	he.buttons.new_button_texture(button_group_main, &he.textures[texture_icon_file_save], 62, 0, func() {
		he.button_save_file()
	})
}
