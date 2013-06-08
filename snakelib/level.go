package snakelib

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"math/rand"
	"time"
)

type Level struct {
	_map *Map
	allowed_duration time.Duration
	end_time time.Time
	game *Game
	player_snake *Snake
}

func NewLevel( game *Game ) *Level {
	var l Level
	l.game = game
	l.player_snake = NewSnake( IntPos{ 10, 10 }, 30 )
	l.allowed_duration = 100 * time.Second
	l.end_time = time.Now().Add( l.allowed_duration )
	width := 300
	height := 150
	l._map = NewEmptyMap( IntPos { width, height })
	for i := 0; i < 30; i++ {
		apple_pos := IntPos{ rand.Intn( width - 2 ) + 1, rand.Intn( height - 2 ) + 1 }
		l._map.SetCell( apple_pos, '*' )
	}

	return &l
}

func (l *Level) GetPlayerSnake() *Snake {
	return l.player_snake
}

func (l *Level) DrawState() {
	time_left := l.end_time.Sub( time.Now() )
	minutes := int( time_left.Minutes() )
	seconds := int( time_left.Seconds() ) - minutes * 60
	state := fmt.Sprintf( "Time remaining: %02d:%02d", minutes, seconds )
	screen_width, _ := termbox.Size()
	l.game.DrawString( screen_width - len( state ), 0, state )
}

