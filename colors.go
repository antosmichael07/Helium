package main

import "image/color"

type color_config struct {
	tool_window_background color.RGBA
	tool_window_border     color.RGBA
}

func init_color_config() color_config {
	return color_config{
		tool_window_background: color.RGBA{0, 0, 0, 255},
		tool_window_border:     color.RGBA{20, 20, 20, 255},
	}
}
