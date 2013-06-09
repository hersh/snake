package snakelib

import (
	"github.com/nsf/termbox-go"
	"fmt"
)

type Game struct {
        dir string
	current_level_num, score int
	current_level *Level
}

func LoadNewGame( game_dir string ) *Game {
	var game Game
	game.dir = game_dir
	game.current_level_num = 0
	game.score = 0
	return &game
}

func (g *Game) AddScore( amount int ) {
	g.score += amount
}

func (g *Game) DrawState() {
	state := fmt.Sprintf( "Score: %d, Level %d", g.score, g.current_level_num )
	g.DrawString( 0, 0, state )
}

func (g *Game) DrawString( x, y int, str string ) {
	runes := []rune( str )
	for i := 0; i < len( runes ); i++ {
		termbox.SetCell( x + i, y, runes[ i ], termbox.ColorWhite, termbox.ColorBlack )
	}
}

func (game *Game) Run() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	// while show-intro-screen() != quit:
	//   level_num = 0
	//   winning = true
	//   while winning
	//     level = loadlevel( level_num )
        //     winning = level.run()
	//   if show-high-scores() == quit
	//     return

	game.current_level = NewLevel( game )
	game.current_level.Run()
}

