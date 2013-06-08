package snakelib

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"time"
)

type Level struct {
	_map Map
	allowed_duration_seconds int
	end_time time.Time
	game *Game
}

func NewLevel( game *Game ) *Level {
	var l Level
	l.game = game
	return &l
}

func (l *Level) DrawState() {
	time_left := l.end_time.Sub( time.Now() )
	state := fmt.Sprintf( "Time remaining: %02d:%02d", time_left.Minutes(), time_left.Seconds() )
	screen_width, _ := termbox.Size()
	l.game.DrawString( screen_width - len( state ), 0, state )
}

