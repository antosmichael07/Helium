package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type tool_window struct {
	rect         rl.Rectangle
	border_rect  rl.Rectangle
	holding_rect rl.Rectangle
	holding      bool
	selected_tab byte
	tabs         []tab
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

func init_tool_window() tool_window {
	return tool_window{
		rect:         rl.NewRectangle(101, 130, 598, 369),
		border_rect:  rl.NewRectangle(100, 100, 600, 400),
		holding_rect: rl.NewRectangle(100, 100, 600, 30),
		holding:      false,
		selected_tab: 0,
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
		tw.holding = false
	} else if (rl.IsMouseButtonPressed(rl.MouseLeftButton) && rl.CheckCollisionPointRec(rl.GetMousePosition(), tw.holding_rect)) || tw.holding {
		tw.holding = true
		if mouse_delta := rl.GetMouseDelta(); mouse_delta.X != 0 || mouse_delta.Y != 0 {
			tw.rect.X += mouse_delta.X
			tw.rect.Y += mouse_delta.Y
			tw.border_rect.X += mouse_delta.X
			tw.border_rect.Y += mouse_delta.Y
			tw.holding_rect.X += mouse_delta.X
			tw.holding_rect.Y += mouse_delta.Y
		}
	}
}

func (tw *tool_window) draw(cc *color_config) {
	rl.DrawRectangleRec(tw.border_rect, cc.tool_window_border)
	rl.DrawRectangleRec(tw.rect, cc.tool_window_background)
}
