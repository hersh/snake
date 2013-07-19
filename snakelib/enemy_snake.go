package snakelib

import (
	"math/rand"
)

type EnemySnake struct {
	*Snake
	*PathPlanner
}

func NewEnemySnake( start IntPos, length int, body_char rune ) *EnemySnake {
	var es EnemySnake
	es.Snake = NewSnake( start, length, body_char )
	es.PathPlanner = new( PathPlanner )
	return &es
}

func (es *EnemySnake) Update( _map *Map ) {

	new_dir, err := es.DirTowardsNearest( _map, es.HeadPos(), '*' )

	if err == nil {
		es.SetDir( new_dir )
	} else {
		es.Turn( rand.Intn( 3 ) - 1) // turn left, go straight, or turn right randomly
	}

	es.Snake.Update( _map )
}
