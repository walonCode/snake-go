package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/walonCode/snake-go/deque"
)

type food struct {
	position rl.Vector2
	texture rl.Texture2D
}

func newFood(snakeBody deque.Deque[rl.Vector2]) *food{
	image := rl.LoadImage("")
	food :=  food{
		texture: rl.LoadTextureFromImage(image),
	}

	food.position = food.generateRandomPos(snakeBody)

	return &food
}

func(f *food)getRandomCell()rl.Vector2{
	x := rl.GetRandomValue(0, int32(cellcount) - 1)
	y := rl.GetRandomValue(0, int32(cellcount) - 1)

	return rl.Vector2{
		X: float32(x),
		Y: float32(y),
	}
}

func (f *food )draw(){
	rl.DrawTexture(f.texture, int32(offset) + int32(f.position.X) * int32(cellsize), int32(offset) + int32(f.position.Y) * int32(cellsize), rl.White)
}

func (f *food)generateRandomPos(body deque.Deque[rl.Vector2])rl.Vector2 {
	position := f.getRandomCell()

	for elementInDeque(position, body) {
		position = f.getRandomCell()
	}

	return position
}

func(f *food)close(){
	rl.UnloadTexture(f.texture)
}