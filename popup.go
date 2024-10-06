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

		he.update_buttons(button_group_popup)

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
			} else if char := rl.GetCharPressed(); ((key_pressed >= rl.KeyZero && key_pressed <= rl.KeyNine) || (key_pressed >= rl.KeyKp0 && key_pressed <= rl.KeyKp9)) && rl.MeasureText(msg_before_cursor+string(char)+"|"+msg_after_cursor, 20) < 520 {
				msg_before_cursor += string(char)
			}
		}

		he.buttons.last_update()

		rl.BeginDrawing()

		rl.DrawRectangle(window_width/2-300, window_height/2-70, 600, 140, he.config.color_config.window_border)
		rl.DrawRectangle(window_width/2-299, window_height/2-69, 598, 138, he.config.color_config.window_background)
		rl.DrawRectangle(window_width/2-260, window_height/2-15, 520, 30, he.config.color_config.button)
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

		he.update_buttons(button_group_popup)

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
			} else if char := rl.GetCharPressed(); char >= 32 && char <= 125 && rl.MeasureText(msg_before_cursor+string(char)+"|"+msg_after_cursor, 20) < 520 {
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
	message_len := []int32{}
	split_text := []string{}

	for i := 0; i < len(message); {
		white_space := i
		for white_space < len(message) && message[white_space] != ' ' && message[white_space] != '\t' && message[white_space] != '\n' && message[white_space] != '\r' && message[white_space] != '\v' && message[white_space] != '\f' {
			white_space++
		}

		if white_space == len(message) {
			split_text = append(split_text, message)
			message_len = append(message_len, rl.MeasureText(split_text[len(split_text)-1], 20))
			break
		}

		if rl.MeasureText(message[:white_space], 20) > 520 {
			split_text = append(split_text, message[:i-1])
			message_len = append(message_len, rl.MeasureText(split_text[len(split_text)-1], 20))
			message = message[i:]
			i = 0
		} else {
			i = white_space + 1
		}
	}

	for {
		window_maganer()

		window_width := int32(rl.GetScreenWidth())
		window_height := int32(rl.GetScreenHeight())

		he.buttons.group[button_group_popup][0].rect.X = float32(window_width)/2 - he.buttons.group[button_group_popup][0].rect.Width/2
		he.buttons.group[button_group_popup][0].rect.Y = float32(window_height)/2 + float32(len(split_text))*15 - 15

		he.update_buttons(button_group_popup)

		if he.buttons.clicked_ok {
			he.buttons.last_update()
			break
		}

		if rl.IsKeyPressed(rl.KeyEnter) || rl.IsKeyPressed(rl.KeyKpEnter) || rl.IsKeyPressed(rl.KeyEscape) {
			break
		}

		he.buttons.last_update()

		rl.BeginDrawing()

		rl.DrawRectangle(window_width/2-300, window_height/2-(15*int32(len(split_text)))-35, 600, (30*int32(len(split_text)))+70, he.config.color_config.window_border)
		rl.DrawRectangle(window_width/2-299, window_height/2-(15*int32(len(split_text)))-34, 598, (30*int32(len(split_text)))+68, he.config.color_config.window_background)
		he.buttons.draw(button_group_popup, &he.config.color_config)
		for i := 0; i < len(split_text); i++ {
			rl.DrawText(split_text[i], window_width/2-message_len[i]/2, window_height/2-(15*int32(len(split_text)))+(int32(i)*30)-15, 20, he.config.color_config.text)
		}

		rl.EndDrawing()
	}
}

func (he *heightmap_editor) popup_error(err string) {
	err = "Error: " + err
	err_len := []int32{}
	split_text := []string{}

	for i := 0; i < len(err); {
		white_space := i
		for white_space < len(err) && err[white_space] != ' ' && err[white_space] != '\t' && err[white_space] != '\n' && err[white_space] != '\r' && err[white_space] != '\v' && err[white_space] != '\f' {
			white_space++
		}

		if white_space == len(err) {
			split_text = append(split_text, err)
			err_len = append(err_len, rl.MeasureText(split_text[len(split_text)-1], 20))
			break
		}

		if rl.MeasureText(err[:white_space], 20) > 520 {
			split_text = append(split_text, err[:i-1])
			err_len = append(err_len, rl.MeasureText(split_text[len(split_text)-1], 20))
			err = err[i:]
			i = 0
		} else {
			i = white_space + 1
		}
	}

	for {
		window_maganer()

		window_width := int32(rl.GetScreenWidth())
		window_height := int32(rl.GetScreenHeight())

		he.buttons.group[button_group_popup][0].rect.X = float32(window_width)/2 - he.buttons.group[button_group_popup][0].rect.Width/2
		he.buttons.group[button_group_popup][0].rect.Y = float32(window_height)/2 + float32(len(split_text))*15 - 15

		he.update_buttons(button_group_popup)

		if he.buttons.clicked_ok {
			he.buttons.last_update()
			break
		}

		if rl.IsKeyPressed(rl.KeyEnter) || rl.IsKeyPressed(rl.KeyKpEnter) || rl.IsKeyPressed(rl.KeyEscape) {
			break
		}

		he.buttons.last_update()

		rl.BeginDrawing()

		rl.DrawRectangle(window_width/2-300, window_height/2-(15*int32(len(split_text)))-35, 600, (30*int32(len(split_text)))+70, he.config.color_config.window_border)
		rl.DrawRectangle(window_width/2-299, window_height/2-(15*int32(len(split_text)))-34, 598, (30*int32(len(split_text)))+68, he.config.color_config.window_background)
		he.buttons.draw(button_group_popup, &he.config.color_config)
		for i := 0; i < len(split_text); i++ {
			rl.DrawText(split_text[i], window_width/2-err_len[i]/2, window_height/2-(15*int32(len(split_text)))+(int32(i)*30)-15, 20, rl.Red)
		}

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
