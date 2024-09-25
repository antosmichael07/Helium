package main

import (
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var refresh_rate int32

func init_window() {
	rl.SetTraceLogLevel(rl.LogError)
	rl.InitWindow(1280, 720, "Heightmap Editor")
	rl.SetExitKey(-1)
	rl.SetWindowState(rl.FlagWindowResizable)
	rl.MaximizeWindow()
	refresh_rate = int32(rl.GetMonitorRefreshRate(rl.GetCurrentMonitor()))
	rl.SetTargetFPS(refresh_rate)
}

func window_maganer() {
	if rl.WindowShouldClose() {
		os.Exit(0)
	}

	if rl.IsWindowFocused() {
		rl.SetTargetFPS(refresh_rate)
	} else {
		rl.SetTargetFPS(30)
	}
}
