package main

import (
	"errors"
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (he *heightmap_editor) popup_uint(message string) (int, error) {
	msg_before_cursor := ""
	msg_after_cursor := ""

	message_len := rl.MeasureText(message, 20)

	for {
		window_maganer()

		window_width := int32(rl.GetScreenWidth())
		window_height := int32(rl.GetScreenHeight())

		he.buttons.group[button_group_popup][0].rect.X = float32(window_width)/2 - he.buttons.group[button_group_popup][0].rect.Width/2
		he.buttons.group[button_group_popup][0].rect.Y = float32(window_height)/2 + 30

		he.buttons.update(button_group_popup)

		if he.buttons.clicked_ok {
			he.buttons.last_update()
			break
		}

		if key_pressed := rl.GetKeyPressed(); key_pressed != 0 {
			if key_pressed == rl.KeyBackspace {
				if len(msg_before_cursor) > 0 {
					msg_before_cursor = msg_before_cursor[:len(msg_before_cursor)-1]
				}
			} else if key_pressed == rl.KeyEnter {
				break
			} else if key_pressed == rl.KeyKpEnter {
				break
			} else if key_pressed == rl.KeyEscape {
				return 0, errors.New("esc")
			} else if key_pressed == rl.KeyRight {
				if len(msg_after_cursor) > 0 {
					msg_before_cursor += string(msg_after_cursor[0])
					msg_after_cursor = msg_after_cursor[1:]
				}
			} else if key_pressed == rl.KeyLeft {
				if len(msg_before_cursor) > 0 {
					msg_after_cursor = string(msg_before_cursor[len(msg_before_cursor)-1]) + msg_after_cursor
					msg_before_cursor = msg_before_cursor[:len(msg_before_cursor)-1]
				}
			} else if (key_pressed >= rl.KeyZero && key_pressed <= rl.KeyNine) || (key_pressed >= rl.KeyKp0 && key_pressed <= rl.KeyKp9) {
				msg_before_cursor += string(rl.GetCharPressed())
			}
		}

		he.buttons.last_update()

		rl.BeginDrawing()

		rl.DrawRectangle(window_width/2-300, window_height/2-70, 600, 140, he.config.color_config.window_border)
		rl.DrawRectangle(window_width/2-299, window_height/2-69, 598, 138, he.config.color_config.window_background)
		rl.DrawRectangle(window_width/2-259, window_height/2-15, 518, 30, he.config.color_config.button)
		he.buttons.draw(button_group_popup, &he.config.color_config)
		rl.DrawText(message, window_width/2-message_len/2, window_height/2-50, 20, he.config.color_config.text)
		rl.DrawText(msg_before_cursor+"|"+msg_after_cursor, window_width/2-rl.MeasureText(msg_before_cursor+"|"+msg_after_cursor, 20)/2, window_height/2-10, 20, he.config.color_config.text)

		rl.EndDrawing()
	}

	if check_white_space(msg_before_cursor + msg_after_cursor) {
		return 0, errors.New("empty input")
	}

	value, err := strconv.Atoi(msg_before_cursor + msg_after_cursor)
	if err != nil || value <= 0 {
		return 0, errors.New("invalid input")
	}

	return value, nil
}

func (he *heightmap_editor) popup_string(message string) (string, error) {
	msg_before_cursor := ""
	msg_after_cursor := ""

	message_len := rl.MeasureText(message, 20)

	for {
		window_maganer()

		window_width := int32(rl.GetScreenWidth())
		window_height := int32(rl.GetScreenHeight())

		he.buttons.group[button_group_popup][0].rect.X = float32(window_width)/2 - he.buttons.group[button_group_popup][0].rect.Width/2
		he.buttons.group[button_group_popup][0].rect.Y = float32(window_height)/2 + 30

		he.buttons.update(button_group_popup)

		if he.buttons.clicked_ok {
			he.buttons.last_update()
			break
		}

		if key_pressed := rl.GetKeyPressed(); key_pressed != 0 {
			if key_pressed == rl.KeyBackspace {
				if len(msg_before_cursor) > 0 {
					msg_before_cursor = msg_before_cursor[:len(msg_before_cursor)-1]
				}
			} else if key_pressed == rl.KeyEnter {
				break
			} else if key_pressed == rl.KeyKpEnter {
				break
			} else if key_pressed == rl.KeyEscape {
				return "", errors.New("esc")
			} else if key_pressed == rl.KeyRight {
				if len(msg_after_cursor) > 0 {
					msg_before_cursor += string(msg_after_cursor[0])
					msg_after_cursor = msg_after_cursor[1:]
				}
			} else if key_pressed == rl.KeyLeft {
				if len(msg_before_cursor) > 0 {
					msg_after_cursor = string(msg_before_cursor[len(msg_before_cursor)-1]) + msg_after_cursor
					msg_before_cursor = msg_before_cursor[:len(msg_before_cursor)-1]
				}
			} else if char := rl.GetCharPressed(); char >= 32 && char <= 125 {
				msg_before_cursor += string(char)
			}
		}

		he.buttons.last_update()

		rl.BeginDrawing()

		rl.DrawRectangle(window_width/2-300, window_height/2-70, 600, 140, he.config.color_config.window_border)
		rl.DrawRectangle(window_width/2-299, window_height/2-69, 598, 138, he.config.color_config.window_background)
		rl.DrawRectangle(window_width/2-259, window_height/2-15, 518, 30, he.config.color_config.button)
		he.buttons.draw(button_group_popup, &he.config.color_config)
		rl.DrawText(message, window_width/2-message_len/2, window_height/2-50, 20, he.config.color_config.text)
		rl.DrawText(msg_before_cursor+"|"+msg_after_cursor, window_width/2-rl.MeasureText(msg_before_cursor+"|"+msg_after_cursor, 20)/2, window_height/2-10, 20, he.config.color_config.text)

		rl.EndDrawing()
	}

	if check_white_space(msg_before_cursor + msg_after_cursor) {
		return "", errors.New("empty input")
	}

	return msg_before_cursor + msg_after_cursor, nil
}

func (he *heightmap_editor) popup_alert(message string) {
	message_len := rl.MeasureText(message, 20)

	for {
		window_maganer()

		window_width := int32(rl.GetScreenWidth())
		window_height := int32(rl.GetScreenHeight())

		he.buttons.group[button_group_popup][0].rect.X = float32(window_width)/2 - he.buttons.group[button_group_popup][0].rect.Width/2
		he.buttons.group[button_group_popup][0].rect.Y = float32(window_height) / 2

		he.buttons.update(button_group_popup)

		if he.buttons.clicked_ok {
			he.buttons.last_update()
			break
		}

		if rl.IsKeyPressed(rl.KeyEnter) || rl.IsKeyPressed(rl.KeyKpEnter) || rl.IsKeyPressed(rl.KeyEscape) {
			break
		}

		he.buttons.last_update()

		rl.BeginDrawing()

		rl.DrawRectangle(window_width/2-300, window_height/2-50, 600, 100, he.config.color_config.window_border)
		rl.DrawRectangle(window_width/2-299, window_height/2-49, 598, 98, he.config.color_config.window_background)
		he.buttons.draw(button_group_popup, &he.config.color_config)
		rl.DrawText(message, window_width/2-message_len/2, window_height/2-30, 20, he.config.color_config.text)

		rl.EndDrawing()
	}
}

func check_white_space(str string) bool {
	for i := 0; i < len(str); i++ {
		if str[i] != ' ' && str[i] != '\t' && str[i] != '\n' && str[i] != '\r' && str[i] != '\v' && str[i] != '\f' {
			return false
		}
	}
	return true
}
