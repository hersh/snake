package snakelib

import (
	"github.com/nsf/termbox-go"
	"fmt"
	"time"
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
	DrawString( 0, 0, state )
}

func DrawString( x, y int, str string ) {
	runes := []rune( str )
	for i := 0; i < len( runes ); i++ {
		termbox.SetCell( x + i, y, runes[ i ], termbox.ColorWhite, termbox.ColorBlack )
	}
}

func DrawCentered( x, y int, str string ) {
	DrawString( x - len( str ) / 2, y, str )
}

func ShowIntroScreen() Result {
	width, height := termbox.Size()
	x := width / 2
	y := height / 2 - 5
	DrawCentered( x, y, "Welcome to GoSnake!" )
	y++
	DrawCentered( x, y, "    I    " ); y++
	DrawCentered( x, y, "    ^    " ); y++
	DrawCentered( x, y, "J <   > L" ); y++
	DrawCentered( x, y, "    v    " ); y++
	DrawCentered( x, y, "    K    " ); y++
	y++
	DrawCentered( x, y, "Press Q to quit," ); y++
	y++
	DrawCentered( x, y, "Any other key to start." ); y++

	if termbox.Flush() != nil {
		return Quit
	}
	for {
		event := termbox.PollEvent()
		if event.Type == termbox.EventKey {
			switch event.Ch {
			case 'q':
				return Quit
			default:
				return Start
			}
		}
		time.Sleep( time.Millisecond * 100 )
	}
	return Start
}

func (game *Game) Run() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	for ; ShowIntroScreen() == Start; {
		game.current_level_num = 1
		winning := true
		for ; winning; {
			game.current_level = NewLevel( game )
			switch game.current_level.Run() {
			case Win:
				game.current_level_num++
			case Lose:
				winning = false
			case Quit:
				return
			}
		}
		//   if show-high-scores() == quit
		//     return
	}
}

