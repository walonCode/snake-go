package main

import rl "github.com/gen2brain/raylib-go/raylib"

var allowMove = false
var green = rl.Color{173, 204, 96, 255}
var darkGreen = rl.Color{43, 51, 24, 255}

var cellsize = 30
var cellcount = 25
var offset = 75

var lastUpdateTime float64 = 0
