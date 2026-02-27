package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

type Point struct {
	x, y int
}

type Game struct {
	player Point
	goal   Point
	walls  []Point
	width  int
	height int
	won    bool
}

func NewGame() *Game {
	return &Game{
		player: Point{1, 1},
		goal:   Point{8, 8},
		walls: []Point{
			{3, 3}, {3, 4}, {3, 5},
			{6, 2}, {6, 3}, {6, 4},
			{2, 7}, {3, 7}, {4, 7},
		},
		width:  10,
		height: 10,
	}
}

func (g *Game) clearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func (g *Game) isWall(p Point) bool {
	for _, w := range g.walls {
		if w.x == p.x && w.y == p.y {
			return true
		}
	}
	return false
}

func (g *Game) Draw() {
	g.clearScreen()
	fmt.Println("Terminal Game: Move '@' to 'G' using WASD (followed by Enter)")
	fmt.Println("Walls: '#', Goal: 'G'")
	fmt.Println(strings.Repeat("-", g.width+2))
	for y := 0; y < g.height; y++ {
		fmt.Print("|")
		for x := 0; x < g.width; x++ {
			p := Point{x, y}
			if p == g.player {
				fmt.Print("@")
			} else if p == g.goal {
				fmt.Print("G")
			} else if g.isWall(p) {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println("|")
	}
	fmt.Println(strings.Repeat("-", g.width+2))
}

func (g *Game) Move(dx, dy int) {
	newPos := Point{g.player.x + dx, g.player.y + dy}
	if newPos.x >= 0 && newPos.x < g.width && newPos.y >= 0 && newPos.y < g.height && !g.isWall(newPos) {
		g.player = newPos
	}
	if g.player == g.goal {
		g.won = true
	}
}

func main() {
	game := NewGame()
	reader := bufio.NewReader(os.Stdin)

	for !game.won {
		game.Draw()
		fmt.Print("Move (WASD): ")
		input, _ := reader.ReadString('\n')
		input = strings.ToLower(strings.TrimSpace(input))

		if len(input) == 0 {
			continue
		}

		char := input[0]
		switch char {
		case 'w':
			game.Move(0, -1)
		case 's':
			game.Move(0, 1)
		case 'a':
			game.Move(-1, 0)
		case 'd':
			game.Move(1, 0)
		case 'q':
			fmt.Println("Quitting game...")
			return
		}
	}

	game.Draw()
	fmt.Println("CONGRATULATIONS! You reached the goal!")
}
