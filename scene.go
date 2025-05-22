package main

import (
	"log"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Scene_Kind int

const (
	Scene_Kind_Menu = iota
	Scene_Kind_Playing
	Scene_Kind_Over
)

type Scene struct {
	Type      Scene_Kind
	SetupFn   func(s *Scene)
	DrawFn    func(s *Scene)
	UpdateFn  func(s *Scene)
	CleanupFn func(s *Scene)
	Data      map[string]any
	Finished  bool
}

// SCENES
const (
	FONT_SIZE = 50
)

func GetValue[T any](m map[string]any, key string) T {
	val, ok := m[key]
	if !ok {
		log.Fatalf("Key '%s' not found!", key)
	}
	typedVal, ok := val.(T)
	if !ok {
		log.Fatalf("Key '%s' has wrong type: got %T", key, val)
	}
	return typedVal
}

func Scene_Menu() Scene {
	const (
		grass_tex_path      = "assets/grass-tile.png"
		zombie_text_offset  = 100
		shooter_text_offset = 45
		play_button_offset  = -100
		play_button_w       = 300
		play_button_h       = 75
	)
	play_rect := rl.NewRectangle(float32(CENTER_X)-play_button_w/2, float32(CENTER_Y)-play_button_offset, play_button_w, play_button_h)
	return Scene{
		Type: Scene_Kind_Menu,
		SetupFn: func(s *Scene) {
			s.Data["grass-texture"] = rl.LoadTexture(grass_tex_path)
		},
		DrawFn: func(s *Scene) {
			// Draw the background
			grass_tex := GetValue[rl.Texture2D](s.Data, "grass-texture")
			grass_tex_w := grass_tex.Width
			grass_tex_h := grass_tex.Height

			for row := int32(0); row < SCREEN_WIDTH/grass_tex_w+1; row++ {
				for col := int32(0); col < SCREEN_HEIGHT/grass_tex_h+1; col++ {
					rl.DrawTexture(grass_tex, row*grass_tex_w, col*grass_tex_h, rl.White)
				}
			}
			rl.DrawText("Zombie", int32(CENTER_X)-int32((rl.MeasureText("Zombie", FONT_SIZE)/2)), int32(CENTER_Y)-zombie_text_offset, FONT_SIZE, rl.White)
			rl.DrawText("Shooter", int32(CENTER_X)-int32((rl.MeasureText("Shooter", FONT_SIZE+15)/2)), int32(CENTER_Y)-shooter_text_offset, FONT_SIZE+15, rl.Red)

			alpha := GetValue[uint8](s.Data, "play-button-alpha")
			color := rl.NewColor(255, 255, 255, alpha)
			rl.DrawRectangleRounded(play_rect, 5.0, 4, color)

			button_center_x := play_rect.ToInt32().X + int32(play_button_w/2)
			button_center_y := play_rect.ToInt32().Y + int32(play_button_h/2)
			rl.DrawText("Play", button_center_x-int32(rl.MeasureText("Play", FONT_SIZE)/2), button_center_y-int32(FONT_SIZE/2), FONT_SIZE, rl.Black)
		},
		UpdateFn: func(s *Scene) {
			if rl.CheckCollisionPointRec(rl.GetMousePosition(), play_rect) {
				s.Data["play-button-alpha"] = uint8(255)
				if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
					s.Finished = true
				}
			} else {
				s.Data["play-button-alpha"] = uint8(127)
			}

		},
		CleanupFn: func(s *Scene) {
			rl.UnloadTexture(GetValue[rl.Texture2D](s.Data, "grass-texture"))
		},
		Data: map[string]any{
			"grass-texture":     rl.Texture2D{},
			"play-button-alpha": uint8(127),
		},
	}
}

func Scene_Playing() Scene {
	return Scene{
		Type:     Scene_Kind_Playing,
		DrawFn:   func(s *Scene) {},
		UpdateFn: func(s *Scene) {},
		Data:     map[string]any{},
	}
}

func Scene_Over() Scene {
	return Scene{
		Type:     Scene_Kind_Over,
		DrawFn:   func(s *Scene) {},
		UpdateFn: func(s *Scene) {},
		Data:     map[string]any{},
	}
}
