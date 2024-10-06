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
	shortcut       bool
}

type button struct {
	text     string
	texture  *rl.Texture2D
	rect     rl.Rectangle
	function func(*heightmap_editor)
}

const (
	button_group_main = iota
	button_group_popup
)

func (he_main *heightmap_editor) init_buttons() {
	he_main.buttons = buttons{}

	he_main.buttons.group = [][]button{}
	he_main.buttons.group_pressed = -1
	he_main.buttons.button_pressed = -1
	he_main.buttons.empty_texture = rl.LoadTextureFromImage(rl.GenImageColor(1, 1, rl.Black))
	he_main.buttons.clicked_ok = false
	he_main.buttons.shortcut = false

	he_main.buttons.new_group()
	he_main.buttons.new_group()

	he_main.init_main_buttons()
	he_main.buttons.new_button(button_group_popup, "OK", 0, 0, func(he *heightmap_editor) {
		he.buttons.clicked_ok = true
	})
}

func (bs *buttons) new_group() {
	bs.group = append(bs.group, []button{})
}

func (bs *buttons) new_button(g int, text string, x, y int, function func(*heightmap_editor)) {
	bs.group[g] = append(bs.group[g], button{
		text:     text,
		texture:  &bs.empty_texture,
		rect:     rl.NewRectangle(float32(x), float32(y), float32(rl.MeasureText(text, 20)+10), float32(30)),
		function: function,
	})
}

func (bs *buttons) new_button_texture(g int, texture *rl.Texture2D, x, y int, function func(*heightmap_editor)) {
	bs.group[g] = append(bs.group[g], button{
		text:     "",
		texture:  texture,
		rect:     rl.NewRectangle(float32(x), float32(y), float32(30), float32(30)),
		function: function,
	})
}

func (he *heightmap_editor) update_buttons(g int) {
	if he.buttons.group_pressed == -1 {
		for i := 0; i < len(he.buttons.group[g]); i++ {
			if (rl.IsMouseButtonPressed(rl.MouseLeftButton) && rl.CheckCollisionPointRec(rl.GetMousePosition(), he.buttons.group[g][i].rect)) || (g == button_group_main && is_shortcut_pressed[i]()) {
				he.buttons.group_pressed = g
				he.buttons.button_pressed = i
				break
			}
		}
	}

	if he.buttons.button_pressed != -1 && g == button_group_main && is_shortcut_pressed[he.buttons.button_pressed]() {
		he.buttons.shortcut = true
		he.buttons.group[g][he.buttons.button_pressed].function(he)
		return
	}

	if he.buttons.group_pressed == g && rl.IsMouseButtonReleased(rl.MouseLeftButton) {
		if rl.CheckCollisionPointRec(rl.GetMousePosition(), he.buttons.group[g][he.buttons.button_pressed].rect) {
			he.buttons.group[g][he.buttons.button_pressed].function(he)
		}
	}
}

func (bs *buttons) last_update() {
	if rl.IsMouseButtonReleased(rl.MouseLeftButton) || bs.shortcut {
		bs.group_pressed = -1
		bs.button_pressed = -1
		bs.clicked_ok = false
		bs.shortcut = false
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
		if *bs.group[g][i].texture == bs.empty_texture {
			rl.DrawText(bs.group[g][i].text, int32(bs.group[g][i].rect.X)+5, int32(bs.group[g][i].rect.Y)+5, 20, cc.text)
		} else {
			rl.DrawTexture(*bs.group[g][i].texture, int32(bs.group[g][i].rect.X)+5, int32(bs.group[g][i].rect.Y)+5, rl.White)
		}
	}
}
