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

func (c *config) init_color_config() {
	c.color_config = color_config{}

	c.color_config.preview_background = color.RGBA{0, 0, 0, 255}
	c.color_config.tool_window_background = color.RGBA{0, 0, 0, 255}
	c.color_config.tool_window_border = color.RGBA{30, 30, 30, 255}
	c.color_config.button_text = color.RGBA{255, 255, 255, 255}
	c.color_config.button = color.RGBA{15, 15, 15, 255}
	c.color_config.button_hover = rl.ColorBrightness(c.color_config.button, 0.1)
	c.color_config.button_click = rl.ColorBrightness(c.color_config.button, 0.15)
}
