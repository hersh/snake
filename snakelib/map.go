package snakelib

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/nsf/termbox-go"
)

type Map struct {
	contents []rune
	size IntPos
	filename string
}

func (m *Map) GetSize() IntPos {
	return m.size
}

func (m *Map) PosValid( pos IntPos ) bool {
	return 0 <= pos.X && pos.X < m.size.X &&
		0 <= pos.Y && pos.Y < m.size.Y
}

func LoadNewMap( scanner *bufio.Scanner, filename string ) (*Map, error) {
	var m Map

	lines := make( []string, 0, 100 )
	maxlen := 0
	for scanner.Scan() {
		line := scanner.Text()
		lines = append( lines, line )
		len_line := len( line )
		if len_line > maxlen {
			maxlen = len_line
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	m.size.Y = len( lines )
	m.size.X = maxlen
//	termbox.Close()
//	fmt.Printf( "width: %d, height: %d\n", m.size.X, m.size.Y )
	m.contents = make( []rune, m.size.Y * m.size.X )

	map_index := 0
	for _, line := range lines {
		var x int
		var ch rune
		for x, ch = range line {
			m.contents[ map_index ] = ch
			map_index++
		}
//		fmt.Printf( "y %d, x %d, map_index %d\n", y, x, map_index )
		for x++; x < maxlen; x++ {
			m.contents[ map_index ] = ' '
			map_index++
		}
	}

	m.filename = filename
	return &m, nil
}

func NewEmptyMap( size IntPos ) *Map {
	var m Map
	m.filename = "no file"
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

func (m *Map) Count( target_char rune ) int {
	count := 0
	for _, ch := range( m.contents ) {
		if ch == target_char {
			count++
		}
	}
	return count
}

func (m *Map) Find( target_char rune ) (IntPos, error) {
	for index, ch := range( m.contents ) {
		if ch == target_char {
			return IntPos{ index % m.size.X, index / m.size.X }, nil
		}
	}
	return IntPos{0,0}, errors.New( fmt.Sprintf( "Rune '%c' not found in map file '%s'.", target_char, m.filename ));
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
