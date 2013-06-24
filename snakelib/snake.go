package snakelib

// import "github.com/nsf/termbox-go"

type IntPos struct {
	X, Y int
}

func (p IntPos) Plus( q IntPos ) IntPos {
	return IntPos { p.X + q.X, p.Y + q.Y }
}

func (p IntPos) Minus( q IntPos ) IntPos {
	return IntPos { p.X - q.X, p.Y - q.Y }
}

func (p IntPos) Div( divisor int ) IntPos {
	return IntPos { p.X / divisor, p.Y / divisor }
}

func (p IntPos) Mult( factor int ) IntPos {
	return IntPos { p.X * factor, p.Y * factor }
}

// Return the position "count" steps from p in direction "dir".
func (p IntPos) PlusDir( dir Direction, count int ) IntPos {
	return p.Plus( motions[ dir ].Mult( count ))
}

func (p IntPos) LowerBound( bound IntPos ) IntPos {
	result := p
	if result.X < bound.X {
		result.X = bound.X
	}
	if result.Y < bound.Y {
		result.Y = bound.Y
	}
	return result
}

func (p IntPos) UpperBound( bound IntPos ) IntPos {
	result := p
	if result.X > bound.X {
		result.X = bound.X
	}
	if result.Y > bound.Y {
		result.Y = bound.Y
	}
	return result
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

func NewSnake( pos IntPos, length int ) *Snake {
	var s Snake
	s.body = make( []IntPos, length )
	for i := range( s.body ) {
		s.body[ i ] = pos
	}
	return &s
}

func (s *Snake) HeadPos() IntPos {
	return s.body[ len( s.body ) - 1 ]
}

func (s *Snake) NextPos() IntPos {
	return s.body[ len(s.body)-1 ].Plus( motions[ s.head_dir ])
}

func (s *Snake) Advance( m *Map ) {
	new_pos := s.NextPos()

	pos := s.body[ 0 ]
	m.SetCell( pos, ' ' )
	var i int
	for i = 1; i < len( s.body ); i++ {
		s.body[ i - 1 ] = s.body[ i ]
	}
	s.body[ i-1 ] = new_pos
	m.SetCell( new_pos, '@' )
}

func (s *Snake) Grow( length_change int ) {
	new_body := make([]IntPos, len( s.body ) + length_change)
	for i := 0; i < length_change; i++ {
		new_body[ i ] = s.body[ 0 ]
	}
	copy( new_body[ length_change: ], s.body )
	s.body = new_body
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

func (s *Snake) SetDir( new_dir Direction ) {
	s.head_dir = new_dir
}