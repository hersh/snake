package main

import (
	"github.com/nsf/termbox-go"
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

	level_map := snakelib.NewEmptyMap( snakelib.IntPos { 100, 100 })

	for ; !done ; {
		snake.Update( level_map )
		level_map.DrawCentered( snake.HeadPos() )
		
		if termbox.Flush() != nil {
			break
		}
		time.Sleep( time.Millisecond * 50 )
	}
}
