package snakelib

import (
	"github.com/nsf/termbox-go"
)

type IntPos struct {
	X, Y int
}

func (p IntPos) plus( q IntPos ) IntPos {
	return IntPos { p.X + q.X, p.Y + q.Y }
}

type Direction int8

const (
	Right Direction = iota
	Up
        Left
	Down
)

// motions are IntPos offsets corresponding to Directions.
var motions = [...]IntPos { {1, 0}, {0, -1}, {-1, 0}, {0, 1} }

type Snake struct {
	body []IntPos
	head_dir Direction
}

func NewSnake() *Snake {
	var s Snake
	s.body = []IntPos{ {10, 10}, {11, 10}, {12, 10}, {12, 11} }
	return &s
}

func (s *Snake) Update() {
	pos := s.body[ 0 ]
	termbox.SetCell( pos.X, pos.Y, ' ', termbox.ColorWhite, termbox.ColorBlack )
	var i int
	for i = 1; i < len( s.body ); i++ {
		s.body[ i - 1 ] = s.body[ i ]
	}
	pos = s.body[ i-1 ].plus( motions[ s.head_dir ])
	s.body[ i-1 ] = pos
	termbox.SetCell( pos.X, pos.Y, '@', termbox.ColorWhite, termbox.ColorBlack )
}

func (s *Snake) Turn( dir_change int ) {
	new_dir := s.head_dir + Direction( dir_change )
	if new_dir < Right {
		new_dir = Down
	} else if new_dir > Down {
		new_dir = Right
	}
	s.head_dir = new_dir
}