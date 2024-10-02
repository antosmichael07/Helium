package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type tool_window struct {
	inner_rect    rl.Rectangle
	border_rect   rl.Rectangle
	holding_rect  rl.Rectangle
	resizing_rect rl.Rectangle
	is_holding    bool
	is_resizing   bool
	resizing_mode byte
	selected_tab  byte
	tabs          []tab
}

type tab struct {
	name string
	mode int32
}

const (
	tabmode_heightmap = iota
	tabmode_texture
	tabmode_model
)

const (
	resizing_top    byte = 0b00000001
	resizing_bottom byte = 0b00000010
	resizing_left   byte = 0b00000100
	resizing_right  byte = 0b00001000
)

func init_tool_window() tool_window {
	return tool_window{
		inner_rect:    rl.NewRectangle(101, 130, 598, 369),
		border_rect:   rl.NewRectangle(100, 100, 600, 400),
		holding_rect:  rl.NewRectangle(100, 100, 600, 30),
		resizing_rect: rl.NewRectangle(94, 94, 612, 412),
		is_holding:    false,
		is_resizing:   false,
		selected_tab:  0,
		tabs: []tab{
			{
				name: "Heightmap",
				mode: tabmode_heightmap,
			},
			{
				name: "Texture 1",
				mode: tabmode_texture,
			},
		},
	}
}

func (tw *tool_window) update() {
	if rl.IsMouseButtonReleased(rl.MouseLeftButton) {
		tw.is_holding = false
		tw.is_resizing = false
		tw.resizing_mode = 0
	}
	tw.movement()
	tw.resizing()
}

func (tw *tool_window) movement() {
	if tw.is_holding || (rl.IsMouseButtonPressed(rl.MouseLeftButton) && rl.CheckCollisionPointRec(rl.GetMousePosition(), tw.holding_rect)) {
		tw.is_holding = true
		if mouse_delta := rl.GetMouseDelta(); mouse_delta.X != 0 || mouse_delta.Y != 0 {
			tw.inner_rect.X += mouse_delta.X
			tw.inner_rect.Y += mouse_delta.Y
			tw.border_rect.X += mouse_delta.X
			tw.border_rect.Y += mouse_delta.Y
			tw.holding_rect.X += mouse_delta.X
			tw.holding_rect.Y += mouse_delta.Y
			tw.resizing_rect.X += mouse_delta.X
			tw.resizing_rect.Y += mouse_delta.Y
		}
	}
}

func (tw *tool_window) resizing() {
	if mouse_pos := rl.GetMousePosition(); tw.is_resizing || (rl.CheckCollisionPointRec(mouse_pos, tw.resizing_rect) && !rl.CheckCollisionPointRec(mouse_pos, tw.border_rect)) {
		rl.SetMouseCursor(rl.MouseCursorResizeNESW)
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			if !tw.is_resizing {
				if mouse_pos.X < tw.border_rect.X {
					tw.resizing_mode |= resizing_left
				}
				if mouse_pos.X > tw.border_rect.X+tw.border_rect.Width {
					tw.resizing_mode |= resizing_right
				}
				if mouse_pos.Y < tw.border_rect.Y {
					tw.resizing_mode |= resizing_top
				}
				if mouse_pos.Y > tw.border_rect.Y+tw.border_rect.Height {
					tw.resizing_mode |= resizing_bottom
				}
			}
			tw.is_resizing = true
		}
		if mouse_delta := rl.GetMouseDelta(); mouse_delta.X != 0 || mouse_delta.Y != 0 {
			if tw.resizing_mode&resizing_top != 0 && tw.border_rect.Height-mouse_delta.Y > 50 {
				tw.inner_rect.Y += mouse_delta.Y
				tw.border_rect.Y += mouse_delta.Y
				tw.holding_rect.Y += mouse_delta.Y
				tw.resizing_rect.Y += mouse_delta.Y
				tw.inner_rect.Height -= mouse_delta.Y
				tw.border_rect.Height -= mouse_delta.Y
				tw.resizing_rect.Height -= mouse_delta.Y
			} else if tw.resizing_mode&resizing_bottom != 0 && tw.border_rect.Height+mouse_delta.Y > 50 {
				tw.inner_rect.Height += mouse_delta.Y
				tw.border_rect.Height += mouse_delta.Y
				tw.resizing_rect.Height += mouse_delta.Y
			}
			if tw.resizing_mode&resizing_left != 0 && tw.border_rect.Width-mouse_delta.X > 50 {
				tw.inner_rect.X += mouse_delta.X
				tw.border_rect.X += mouse_delta.X
				tw.holding_rect.X += mouse_delta.X
				tw.resizing_rect.X += mouse_delta.X
				tw.inner_rect.Width -= mouse_delta.X
				tw.border_rect.Width -= mouse_delta.X
				tw.holding_rect.Width -= mouse_delta.X
				tw.resizing_rect.Width -= mouse_delta.X
			} else if tw.resizing_mode&resizing_right != 0 && tw.border_rect.Width+mouse_delta.X > 50 {
				tw.inner_rect.Width += mouse_delta.X
				tw.border_rect.Width += mouse_delta.X
				tw.holding_rect.Width += mouse_delta.X
				tw.resizing_rect.Width += mouse_delta.X
			}
		}
	} else {
		rl.SetMouseCursor(rl.MouseCursorArrow)
	}
}

func (tw *tool_window) draw(cc *color_config) {
	rl.DrawRectangleRec(tw.border_rect, cc.tool_window_border)
	rl.DrawRectangleRec(tw.inner_rect, cc.tool_window_background)
}
