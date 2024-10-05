package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type tool_window struct {
	inner_rect     rl.Rectangle
	border_rect    rl.Rectangle
	holding_rect   rl.Rectangle
	resizing_rect  rl.Rectangle
	tabs_rect      rl.Rectangle
	is_holding     bool
	is_resizing    bool
	resizing_mode  byte
	selected_tab   int
	tabs           [3]string
	tab_widths     [3]int32
	tab_offsets    [3]int32
	button_pressed int
}

const (
	tab_heightmap = iota
	tab_texture
	tab_model
)

const (
	resizing_top    byte = 0b00000001
	resizing_bottom byte = 0b00000010
	resizing_left   byte = 0b00000100
	resizing_right  byte = 0b00001000
)

func init_tool_window() tool_window {
	return tool_window{
		inner_rect:     rl.NewRectangle(101, 140, 598, 359),
		border_rect:    rl.NewRectangle(100, 100, 600, 400),
		holding_rect:   rl.NewRectangle(100, 100, 600, 40),
		resizing_rect:  rl.NewRectangle(94, 94, 612, 412),
		tabs_rect:      rl.NewRectangle(101, 140, 598, 30),
		is_holding:     false,
		is_resizing:    false,
		resizing_mode:  0,
		selected_tab:   0,
		tabs:           [3]string{"Heightmap", "Texture", "Models"},
		tab_widths:     [3]int32{108, 94, 78},
		tab_offsets:    [3]int32{0, 108, 202},
		button_pressed: -1,
	}
}

func (tw *tool_window) update() {
	tw.button_press()
	tw.movement()
	tw.resizing()
	tw.tab_selecting()
	if rl.IsMouseButtonReleased(rl.MouseLeftButton) {
		tw.is_holding = false
		tw.is_resizing = false
		tw.resizing_mode = 0
		tw.button_pressed = -1
	}
}

func (tw *tool_window) button_press() {
	if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		for i := 0; i < len(tw.tabs); i++ {
			if rl.CheckCollisionPointRec(rl.GetMousePosition(), rl.NewRectangle(float32(tw.tab_offsets[i])+tw.tabs_rect.X, float32(tw.tabs_rect.Y), float32(rl.MeasureText(tw.tabs[i], 20)+10), 30)) {
				tw.button_pressed = i
				break
			}
		}
	}
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
			tw.tabs_rect.X += mouse_delta.X
			tw.tabs_rect.Y += mouse_delta.Y
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
			if tw.resizing_mode&resizing_top != 0 && tw.border_rect.Height-mouse_delta.Y >= 71 {
				tw.inner_rect.Y += mouse_delta.Y
				tw.border_rect.Y += mouse_delta.Y
				tw.holding_rect.Y += mouse_delta.Y
				tw.resizing_rect.Y += mouse_delta.Y
				tw.tabs_rect.Y += mouse_delta.Y
				tw.inner_rect.Height -= mouse_delta.Y
				tw.border_rect.Height -= mouse_delta.Y
				tw.resizing_rect.Height -= mouse_delta.Y
			} else if tw.resizing_mode&resizing_bottom != 0 && tw.border_rect.Height+mouse_delta.Y >= 71 {
				tw.inner_rect.Height += mouse_delta.Y
				tw.border_rect.Height += mouse_delta.Y
				tw.resizing_rect.Height += mouse_delta.Y
			}
			if tw.resizing_mode&resizing_left != 0 && tw.border_rect.Width-mouse_delta.X >= 282 {
				tw.inner_rect.X += mouse_delta.X
				tw.border_rect.X += mouse_delta.X
				tw.holding_rect.X += mouse_delta.X
				tw.resizing_rect.X += mouse_delta.X
				tw.tabs_rect.X += mouse_delta.X
				tw.inner_rect.Width -= mouse_delta.X
				tw.border_rect.Width -= mouse_delta.X
				tw.holding_rect.Width -= mouse_delta.X
				tw.resizing_rect.Width -= mouse_delta.X
				tw.tabs_rect.Width -= mouse_delta.X
			} else if tw.resizing_mode&resizing_right != 0 && tw.border_rect.Width+mouse_delta.X >= 282 {
				tw.inner_rect.Width += mouse_delta.X
				tw.border_rect.Width += mouse_delta.X
				tw.holding_rect.Width += mouse_delta.X
				tw.resizing_rect.Width += mouse_delta.X
				tw.tabs_rect.Width += mouse_delta.X
			}
		}
	} else {
		rl.SetMouseCursor(rl.MouseCursorArrow)
	}
}

func (tw *tool_window) tab_selecting() {
	if rl.IsMouseButtonReleased(rl.MouseLeftButton) && tw.button_pressed != -1 {
		if rl.CheckCollisionPointRec(rl.GetMousePosition(), rl.NewRectangle(float32(tw.tab_offsets[tw.button_pressed])+tw.tabs_rect.X, float32(tw.tabs_rect.Y), float32(rl.MeasureText(tw.tabs[tw.button_pressed], 20)+10), 30)) {
			tw.selected_tab = tw.button_pressed
		}
	}
}

func (tw *tool_window) draw(cc *color_config) {
	rl.DrawRectangleRec(tw.border_rect, cc.tool_window_border)
	rl.DrawRectangleRec(tw.inner_rect, cc.tool_window_background)
	rl.DrawText("Tool Window", int32(tw.border_rect.X)+5, int32(tw.border_rect.Y)+5, 30, cc.tool_window_text)
	tw.draw_tabs(cc)
}

func (tw *tool_window) draw_tabs(cc *color_config) {
	rl.DrawRectangleRec(tw.tabs_rect, cc.tool_window_tab)
	for i := 0; i < len(tw.tabs); i++ {
		if i == tw.selected_tab {
			if !rl.IsMouseButtonDown(rl.MouseLeftButton) || tw.button_pressed == -1 {
				rl.DrawRectangle(tw.tab_offsets[i]+int32(tw.tabs_rect.X), int32(tw.tabs_rect.Y), tw.tab_widths[i], 30, cc.tool_window_tab_click)
			} else {
				rl.DrawRectangle(tw.tab_offsets[i]+int32(tw.tabs_rect.X), int32(tw.tabs_rect.Y), tw.tab_widths[i], 30, cc.tool_window_tab_hover)
			}
		} else if rl.CheckCollisionPointRec(rl.GetMousePosition(), rl.NewRectangle(float32(tw.tab_offsets[i])+tw.tabs_rect.X, float32(tw.tabs_rect.Y), float32(tw.tab_widths[i]), 30)) {
			if tw.button_pressed == i {
				rl.DrawRectangle(tw.tab_offsets[i]+int32(tw.tabs_rect.X), int32(tw.tabs_rect.Y), tw.tab_widths[i], 30, cc.tool_window_tab_click)
			} else if !rl.IsMouseButtonDown(rl.MouseLeftButton) {
				rl.DrawRectangle(tw.tab_offsets[i]+int32(tw.tabs_rect.X), int32(tw.tabs_rect.Y), tw.tab_widths[i], 30, cc.tool_window_tab_hover)
			}
		}
		rl.DrawText(tw.tabs[i], tw.tab_offsets[i]+int32(tw.tabs_rect.X)+5, int32(tw.tabs_rect.Y)+5, 20, cc.tool_window_text)
	}
}
