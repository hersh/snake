package snakelib

import "github.com/nsf/termbox-go"

type Map struct {
	contents []rune
	size IntPos
}

func NewEmptyMap( size IntPos ) *Map {
	var m Map
	m.size = size
	m.contents = make( []rune, size.X * size.Y )

	for i := range( m.contents ) {
		m.contents[ i ] = ' '
	}

	for x := 0; x < size.X; x++ {
		m.SetCell( IntPos{ x, 0 }, '#' )
		m.SetCell( IntPos{ x, size.Y - 1 }, '#' )
	}
	for y := 0; y < size.Y; y++ {
		m.SetCell( IntPos{ 0, y }, '#' )
		m.SetCell( IntPos{ size.X - 1, y }, '#' )
	}

	return &m
}

func (m *Map) SetCell( pos IntPos, ch rune ) {
	if pos.X < 0 || pos.Y < 0 || pos.X >= m.size.X || pos.Y >= m.size.Y {
		return
	}
	m.contents[ pos.X + pos.Y * m.size.X ] = ch
}

func (m *Map) GetCell( pos IntPos ) rune {
	return m.contents[ pos.X + pos.Y * m.size.X ]
}

// Draw the map to the terminal, centered around the given point in the map.
func (m *Map) DrawCentered( point_rel_map IntPos ) {
	termbox.Clear( termbox.ColorWhite, termbox.ColorBlack )
 
	var term_size IntPos
	term_size.X, term_size.Y = termbox.Size()

	offset := point_rel_map.Minus( term_size.Div( 2 ))

	map_min := offset.LowerBound( IntPos{ 0, 0 })
	map_max := offset.Plus( term_size ).UpperBound( m.size )
	term_min := map_min.Minus( offset )
	term_max := map_max.Minus( offset )

	for map_y, term_y := map_min.Y, term_min.Y; map_y < map_max.Y; map_y, term_y = map_y + 1, term_y + 1 {
		map_index := map_min.X + map_y * m.size.X
		for term_x := term_min.X; term_x < term_max.X; term_x++ {
			termbox.SetCell( term_x, term_y, m.contents[ map_index ],
				termbox.ColorWhite, termbox.ColorBlack )
			map_index++
		}
	}
}
