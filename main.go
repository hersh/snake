package main

import "snake/snakelib"

func main() {
	game := snakelib.LoadNewGame( "." )
	game.Run()
}
