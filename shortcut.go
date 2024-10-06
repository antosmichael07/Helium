package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var is_shortcut_pressed = []func() bool{
	func() bool { return rl.IsKeyPressed(rl.KeyN) && rl.IsKeyDown(rl.KeyLeftControl) },
	func() bool { return rl.IsKeyPressed(rl.KeyO) && rl.IsKeyDown(rl.KeyLeftControl) },
	func() bool { return rl.IsKeyPressed(rl.KeyS) && rl.IsKeyDown(rl.KeyLeftControl) },
}
