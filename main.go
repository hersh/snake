package main

import (
	"github.com/nsf/termbox-go"
	"math/rand"
	"time"
	"snake/snakelib"
)

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	snake := snakelib.NewSnake( snakelib.IntPos{ 10, 10 }, 30 )
	done := false

	go func() {
		for {
			event := termbox.PollEvent()
			if event.Type == termbox.EventKey {
				switch event.Ch {
				case 'q':
					done = true
				case 'l':
					snake.Turn( -1 )
				case 'j':
					snake.Turn( 1 )
				}
			}
		}
	}()

	width := 300
	height := 150
	level_map := snakelib.NewEmptyMap( snakelib.IntPos { width, height })
	for i := 0; i < 30; i++ {
		apple_pos := snakelib.IntPos{ rand.Intn( width - 2 ) + 1, rand.Intn( height - 2 ) + 1 }
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
