package snakelib

import (
	"github.com/nsf/termbox-go"
	"fmt"
	"math/rand"
	"time"
)

type Game struct {
        dir string
	current_level_num, score int
//	current_level Level
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

	snake := NewSnake( IntPos{ 10, 10 }, 30 )
	done := false

	go func() {
		for {
			event := termbox.PollEvent()
			if event.Type == termbox.EventKey {
				switch event.Ch {
				case 'q':
					done = true
				case 'l':
					snake.SetDir( Right )
				case 'j':
					snake.SetDir( Left )
				case 'k':
					snake.SetDir( Down )
				case 'i':
					snake.SetDir( Up )
				}
			}
		}
	}()

	width := 300
	height := 150
	level_map := NewEmptyMap( IntPos { width, height })
	for i := 0; i < 30; i++ {
		apple_pos := IntPos{ rand.Intn( width - 2 ) + 1, rand.Intn( height - 2 ) + 1 }
		level_map.SetCell( apple_pos, '*' )
	}

	for ; !done ; {
		snake.Update( level_map )
		level_map.DrawCentered( snake.HeadPos() )
		
		if termbox.Flush() != nil {
			break
		}
		time.Sleep( time.Millisecond * 100 )
	}
}

