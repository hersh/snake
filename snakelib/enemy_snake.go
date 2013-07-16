package snakelib

import (
	"math/rand"
)

type EnemySnake struct {
	Snake
	PathPlanner
}

func (es *EnemySnake) update( _map *Map ) {

	new_dir, err := es.DirTowardsNearest( _map, es.HeadPos(), '*' )
	if err == nil {
		es.SetDir( new_dir )
	} else {
		es.Turn( rand.Intn( 3 ) - 1) // turn left, go straight, or turn right randomly
	}
	es.Advance( _map )
}
