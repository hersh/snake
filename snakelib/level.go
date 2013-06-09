package snakelib

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"math/rand"
	"time"
)

type Result int8

const (
	Win Result = iota
	Lose
	Quit
	Start
)

type Level struct {
	_map *Map
	allowed_duration time.Duration
	end_time time.Time
	game *Game
	player_snake *Snake
	apples_remaining int
}

func NewLevel( game *Game ) *Level {
	var l Level
	l.game = game
	l.player_snake = NewSnake( IntPos{ 10, 10 }, 30 )
	l.allowed_duration = 100 * time.Second
	l.end_time = time.Now().Add( l.allowed_duration )
	width := 80
	height := 40
	l._map = NewEmptyMap( IntPos { width, height })
	l.apples_remaining = 10
	for i := 0; i < l.apples_remaining; i++ {
		apple_pos := IntPos{ rand.Intn( width - 2 ) + 1, rand.Intn( height - 2 ) + 1 }
		l._map.SetCell( apple_pos, '*' )
	}

	return &l
}

func (level *Level) Run() Result {
	snake := level.player_snake

	quit := false
	go func() {
		for {
			event := termbox.PollEvent()
			if event.Type == termbox.EventKey {
				switch event.Ch {
				case 'q':
					quit = true
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

	stopped := false

	for {
		if quit {
			return Quit
		}

		if level.end_time.Sub( time.Now() ).Seconds() < 0 {
			return Lose
		}

		switch level._map.GetCell( snake.NextPos() ) {
		case '#', '@':
			if !stopped {
				level.game.AddScore( -10 )
				stopped = true
			}
		case '*':
			level.game.AddScore( 5 )
			snake.Grow( 5 )
			level.apples_remaining--
			if level.apples_remaining <= 0 {
				return Win
			}
			stopped = false
		default:
			stopped = false
		}

		if !stopped {
			snake.Advance( level._map )
		}
		level._map.DrawCentered( snake.HeadPos() )
		level.game.DrawState()
		level.DrawState()

		if termbox.Flush() != nil {
			break
		}
		time.Sleep( time.Millisecond * 100 )
	}

	return Quit
}

func (l *Level) DrawState() {
	time_left := l.end_time.Sub( time.Now() )
	minutes := int( time_left.Minutes() )
	seconds := int( time_left.Seconds() ) - minutes * 60
	state := fmt.Sprintf( "Apples: %d Time: %02d:%02d",
		l.apples_remaining, minutes, seconds )
	screen_width, _ := termbox.Size()
	DrawString( screen_width - len( state ), 0, state )
}

