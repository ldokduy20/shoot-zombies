package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	SCREEN_WIDTH  = 800
	SCREEN_HEIGHT = 450
	CENTER_X      = int(SCREEN_WIDTH / 2)
	CENTER_Y      = int(SCREEN_HEIGHT / 2)
)

func main() {
	rl.InitWindow(800, 450, "Hi")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	game := NewGame()

	for !rl.WindowShouldClose() && !game.Finished {
		game.Run()
	}
}
