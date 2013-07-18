package snakelib

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/nsf/termbox-go"
	"math/rand"
	"os"
	"strconv"
	"strings"
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
	enemy_snake *EnemySnake
	apples_remaining int
}

func NewLevel( game *Game ) *Level {
	var l Level
	l.game = game
	l.player_snake = NewSnake( IntPos{ 10, 10 }, 30, '@' )
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
	l.apples_remaining = l._map.Count( '*' )

	return &l
}

func valuesFromLevelFile( scanner *bufio.Scanner, filename string ) (map[string] string, error) {

	values := make( map[string] string )

	for scanner.Scan() {
		line := scanner.Text()
		nameval := strings.Split( line, ":" )
		if len( nameval ) == 2 {
			name := strings.TrimSpace( nameval[ 0 ])
			value := strings.TrimSpace( nameval[ 1 ])
			values[ name ] = value
		}
		if line == "map:" {
			return values, nil
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return nil, errors.New( fmt.Sprintf( "'map:' key not found in level file '%s'", filename ))
}

func LoadNewLevel( game *Game, filename string ) (*Level, error) {
	var l Level
	l.game = game
	
	file, err := os.Open( filename )
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner( file )
	values, err := valuesFromLevelFile( scanner, filename )
	if err != nil {
		return nil, err
	}

	seconds, err := strconv.Atoi( values[ "time" ])
	if err != nil {
		return nil, err
	}

	l.allowed_duration = time.Duration( seconds ) * time.Second
	l.end_time = time.Now().Add( l.allowed_duration )

	l._map, err = LoadNewMap( scanner, filename )
	if err != nil {
		return nil, err
	}
	l.apples_remaining = l._map.Count( '*' )
	snake_start, err := l._map.Find( '@' )
	if err != nil {
		return nil, err
	}
	snake_len, err := strconv.Atoi( values[ "snake len" ])
	if err != nil {
		return nil, err
	}
	l.player_snake = NewSnake( snake_start, snake_len, '@' )

	enemy_start, err := l._map.Find( '&' )
	if err != nil {
		return nil, err
	}
	enemy_len, err := strconv.Atoi( values[ "enemy len" ])
	if err != nil {
		return nil, err
	}
	l.enemy_snake = NewEnemySnake( enemy_start, enemy_len, '&' )
	return &l, nil
}

func isStopper( r rune ) bool {
	return r != ' ' && r != '*'
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

		ch := level._map.GetCell( snake.NextPos() )
		switch {
		case isStopper( ch ):
			if !stopped {
				level.game.AddScore( -10 )
				stopped = true
				if isStopper( level._map.GetCell( snake.HeadPos().PlusDir( Down, 1 ))) &&
					isStopper( level._map.GetCell( snake.HeadPos().PlusDir( Up, 1 ))) &&
					isStopper( level._map.GetCell( snake.HeadPos().PlusDir( Left, 1 ))) &&
					isStopper( level._map.GetCell( snake.HeadPos().PlusDir( Right, 1 ))) {

					return Lose
				}
			}
		case ch == '*':
			level.game.AddScore( 5 )
			snake.Grow( 5 )
			level.apples_remaining = level._map.Count( '*' ) - 1
			if level.apples_remaining == 0 {
				return Win
			}
			stopped = false
		default:
			stopped = false
		}

		if !stopped {
			snake.Advance( level._map )
		}
		
		level.enemy_snake.Update( level._map )

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

