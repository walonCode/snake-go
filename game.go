package main

import rl "github.com/gen2brain/raylib-go/raylib"

type game struct {
	running bool
	score int
	eatSound rl.Sound
	wallSound rl.Sound
	food food
	snake snake
}

func newGame()*game {
	rl.InitAudioDevice()
	return &game {
		eatSound: rl.LoadSound("sounds/eat.mp3"),
		wallSound: rl.LoadSound("sounds/wall.mp3"),
	}
}

func (g *game)close(){
	rl.UnloadSound(g.eatSound)
	rl.UnloadSound(g.wallSound)
	rl.CloseAudioDevice()
}

func (g *game)draw(){
	g.food.draw()
	g.snake.draw()
}

func (g *game)update(){
	if g.running{
		g.snake.update()
  		g.checkCollisionWithFood();
        g.checkCollisionWithEdges();
        g.checkCollisionWithTail();
	}
}

func (g *game)checkCollisionWithFood(){
	if rl.Vector2Equals(g.snake.body.At(0), g.food.position) {
		g.food.position = g.food.generateRandomPos(g.snake.body)
		g.snake.addSegment = true	
		g.score++
		rl.PlaySound(g.eatSound)
	}
}

func (g *game)checkCollisionWithTail(){
	headlessBody := g.snake.body
	headlessBody.PopFront()

	if elementInDeque(g.snake.body.At(0), headlessBody){
		g.gameOver()
	}
}

func (g *game)checkCollisionWithEdges(){
	if g.snake.body.At(0).X == float32(cellcount) || g.snake.body.At(0).X == -1 {
		g.gameOver()
	}
	if g.snake.body.At(0).Y == float32(cellcount) || g.snake.body.At(0).Y == -1 {
		g.gameOver()
	}
}

func (g *game)gameOver(){
	g.snake.reset()
	g.food.position = g.food.generateRandomPos(g.snake.body)
	g.running = false
	g.score = 0
	rl.PauseSound(g.wallSound)
}