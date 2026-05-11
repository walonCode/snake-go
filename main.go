package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/walonCode/snake-go/deque"
)

var allowMove = false
var green = rl.Color{R:173, G:204, B:96, A:255}
var darkGreen = rl.Color{R:43, G:51, B:24, A:255}

var cellsize = 30
var cellcount = 25
var offset = 75

var lastUpdateTime float64 = 0

func elementInDeque( element rl.Vector2, deque deque.Deque[rl.Vector2] )bool {
	for i:=0;i < deque.Size(); i++ {
		if rl.Vector2Equals(deque.At(i), element){
			return true
		}
	}

	return false
}

func eventTriggered(interval float64)bool {
	currentTime := rl.GetTime()
	if currentTime - lastUpdateTime >= interval {
		lastUpdateTime = currentTime
		return true
	}

	return false
}

func main(){
	fmt.Println("Starting the game....")

	rl.SetConfigFlags(rl.FlagWindowResizable | rl.FlagWindowMaximized)
	rl.InitWindow(int32(2 * offset + cellsize * cellcount), int32(2 * offset + cellsize * cellcount), "Snake Game")
	rl.SetTargetFPS(60)
	defer rl.CloseWindow()

	game := newGame()
	defer game.close()

	for !rl.WindowShouldClose(){
		rl.BeginDrawing()

		if eventTriggered(0.2){
			allowMove = true
			game.update()
		}

		if rl.IsKeyPressed(rl.KeyUp) && game.snake.direction.Y != 1 && allowMove {
	  		game.snake.direction = rl.Vector2{X:0, Y:-1};
            game.running = true;
            allowMove = false;
		}

		if rl.IsKeyPressed(rl.KeyDown) && game.snake.direction.Y != -1 && allowMove {
	  		game.snake.direction = rl.Vector2{X:0, Y:1};
            game.running = true;
            allowMove = false;
		}

		if rl.IsKeyPressed(rl.KeyLeft) && game.snake.direction.Y != 1 && allowMove {
	  		game.snake.direction = rl.Vector2{X:-1, Y:0};
            game.running = true;
            allowMove = false;
		}

		if rl.IsKeyPressed(rl.KeyRight) && game.snake.direction.Y != -1 && allowMove {
	  		game.snake.direction = rl.Vector2{X:1, Y:0};
            game.running = true;
            allowMove = false;
		}

		rl.ClearBackground(rl.Green)
		rl.DrawRectangleLinesEx(rl.Rectangle{
			X: float32(offset) -5,
			Y: float32(offset) -5,
			Width: float32(cellsize) * float32(cellcount) + 10,
			Height: float32(cellsize) * float32(cellcount) + 10,
		}, 5, rl.DarkGreen)
		rl.DrawText("Snake Game", int32(offset) - 5, 20, 40 , rl.DarkGreen)
		rl.DrawText(fmt.Sprintf("%d", game.score), int32(offset) - 5, int32(offset) + int32(cellsize) + 10, 40, rl.DarkGreen)

		rl.EndDrawing()
	}
}