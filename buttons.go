package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type buttons struct {
	group          [][]button
	group_pressed  int
	button_pressed int
	empty_texture  rl.Texture2D
}

type button struct {
	text     string
	texture  *rl.Texture2D
	rect     rl.Rectangle
	function func()
}

const (
	button_group_main = iota
)

func init_buttons() buttons {
	return buttons{
		group:          [][]button{},
		group_pressed:  -1,
		button_pressed: -1,
		empty_texture:  rl.LoadTextureFromImage(rl.GenImageColor(1, 1, rl.Black)),
	}
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

func (bs *buttons) new_button_texture(g int, texture *rl.Texture2D, x, y, width, height int, function func()) {
	bs.group[g] = append(bs.group[g], button{
		text:     "",
		texture:  texture,
		rect:     rl.NewRectangle(float32(x), float32(y), float32(width), float32(height)),
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
			rl.DrawText(bs.group[g][i].text, int32(bs.group[g][i].rect.X)+5, int32(bs.group[g][i].rect.Y)+5, 20, cc.button_text)
		} else {
			rl.DrawTexture(*bs.group[g][i].texture, int32(bs.group[g][i].rect.X), int32(bs.group[g][i].rect.Y), rl.White)
		}
	}
}
