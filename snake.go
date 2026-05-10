package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/walonCode/snake-go/deque"
)

type snake struct {
	addSegment bool 
	direction rl.Vector2
	body deque.Deque[rl.Vector2]
}

func(s *snake)draw(){
	for i:=0; i<s.body.Size(); i++ {
		x := s.body.At(i).X
		y := s.body.At(i).Y
		rectangle := rl.Rectangle{
			X: float32(offset) + x * float32(cellsize),
			Y: float32(offset) + y * float32(cellsize),
			Width: float32(cellsize),
			Height: float32(cellsize),
		}
		rl.DrawRectangleRounded(rectangle, 0.5, 6, rl.DarkGreen)
	}
}

func (s *snake)update(){
	s.body.PushFront(rl.Vector2Add(s.body.At(0),s.direction))
	if s.addSegment == true {
		s.addSegment = false
	} else {
		s.body.PopBack()
	}
}

func (s *snake) reset() {
	s.body = deque.New(
		rl.Vector2{6,9},
		rl.Vector2{5,9},
		rl.Vector2{4,9},
	)
	s.direction = rl.Vector2{1,0}
}

func new() *snake{
	return &snake{
		addSegment: false,
		direction: rl.Vector2{X:1, Y: 0},
		body: deque.New(
			rl.Vector2{X: 6, Y: 9},
			rl.Vector2{X: 5, Y: 9},
			rl.Vector2{X: 4, Y: 9},
		),
	}
}