package main

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type color_config struct {
	preview_background     color.RGBA
	tool_window_background color.RGBA
	tool_window_border     color.RGBA
	button_text            color.RGBA
	button                 color.RGBA
	button_hover           color.RGBA
	button_click           color.RGBA
}

func init_color_config() color_config {
	cc := color_config{}

	cc.preview_background = color.RGBA{0, 0, 0, 255}
	cc.tool_window_background = color.RGBA{0, 0, 0, 255}
	cc.tool_window_border = color.RGBA{30, 30, 30, 255}
	cc.button_text = color.RGBA{255, 255, 255, 255}
	cc.button = color.RGBA{15, 15, 15, 255}
	cc.button_hover = rl.ColorBrightness(cc.button, 0.1)
	cc.button_click = rl.ColorBrightness(cc.button, 0.15)

	return cc
}
