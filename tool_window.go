package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type tool_window struct {
	x            int32
	y            int32
	width        int32
	height       int32
	selected_tab int
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
		x:            100,
		y:            100,
		width:        600,
		height:       400,
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

func (tw *tool_window) draw(cc *color_config) {
	rl.DrawRectangle(tw.x, tw.y, tw.width, tw.height, cc.tool_window_background)
	rl.DrawRectangle(tw.x, tw.y, tw.width, 30, cc.tool_window_border)
	rl.DrawRectangleLines(tw.x, tw.y, tw.width, tw.height, cc.tool_window_border)
}
