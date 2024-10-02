package main

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type color_config struct {
	tool_window_background color.RGBA
	tool_window_border     color.RGBA
	tool_window_text       color.RGBA
	tool_window_tab        color.RGBA
	tool_window_tab_hover  color.RGBA
	tool_window_tab_click  color.RGBA
}

func init_color_config() color_config {
	cc := color_config{}

	cc.tool_window_background = color.RGBA{0, 0, 0, 255}
	cc.tool_window_border = color.RGBA{30, 30, 30, 255}
	cc.tool_window_text = color.RGBA{255, 255, 255, 255}
	cc.tool_window_tab = color.RGBA{15, 15, 15, 255}
	cc.tool_window_tab_hover = rl.ColorBrightness(cc.tool_window_tab, 0.1)
	cc.tool_window_tab_click = rl.ColorBrightness(cc.tool_window_tab, 0.15)

	return cc
}
