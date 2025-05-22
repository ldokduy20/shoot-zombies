package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Game struct {
	CurrentScene Scene
	Finished     bool
}

func NewGame() Game {
	game := Game{
		CurrentScene: Scene_Menu(),
		Finished:     false,
	}
	game.CurrentScene.SetupFn(&game.CurrentScene)
	return game
}

func (g *Game) Update() {
	g.CurrentScene.UpdateFn(&g.CurrentScene)
}
func (g *Game) Draw() {
	rl.ClearBackground(rl.Black)
	rl.BeginDrawing()
	g.CurrentScene.DrawFn(&g.CurrentScene)
	rl.EndDrawing()
}

func (g *Game) Run() {
	g.Update()
	g.Draw()
	if g.CurrentScene.Finished {
		g.CurrentScene.CleanupFn(&g.CurrentScene)
		g.CurrentScene.Finished = false
		switch g.CurrentScene.Type {
		case Scene_Kind_Menu:
			g.CurrentScene = Scene_Playing()
		case Scene_Kind_Playing:
			g.CurrentScene = Scene_Over()
		case Scene_Kind_Over:
			g.Finished = true
		}
	}
}
