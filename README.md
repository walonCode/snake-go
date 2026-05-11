# snake-go

A classic Snake game built in Go using [raylib-go](https://github.com/gen2brain/raylib-go).

## Features

- Smooth 60 FPS rendering via raylib
- Rounded snake segments
- Sound effects for eating food and hitting a wall
- Score display
- Game resets on collision with walls or the snake's own tail

## Prerequisites

- Go 1.21+
- raylib system libraries (see [raylib-go install guide](https://github.com/gen2brain/raylib-go#requirements))

## Run

```bash
git clone https://github.com/walonCode/snake-go.git
cd snake-go
go run .
```

## Controls

| Key | Action |
|-----|--------|
| Arrow Up | Move up |
| Arrow Down | Move down |
| Arrow Left | Move left |
| Arrow Right | Move right |

Press any arrow key to start the game. The snake resets after a collision.

## Project Structure

```
snake-go/
├── main.go        # Window setup, game loop, input handling
├── game.go        # Game state, collision detection
├── snake.go       # Snake entity
├── food.go        # Food entity and random placement
├── deque/
│   └── deque.go   # Generic double-ended queue used for the snake body
├── graphics/
│   └── food.png
└── sounds/
    ├── eat.mp3
    └── wall.mp3
```

## Acknowledgements

This project was inspired by and based on the C++ Retro Snake Game tutorial by [educ8s](https://github.com/educ8s/Cpp-Retro-Snake-Game-with-raylib). The core game structure, visual style, and collision logic follow that implementation closely — ported to Go with a generic deque replacing `std::deque`.

## License

MIT — see [LICENSE](LICENSE)
