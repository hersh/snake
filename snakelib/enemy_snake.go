package snakelib

import (
//	"math/rand"
)

type EnemySnake struct {
	Snake
	PathPlanner
}

func (es *EnemySnake) update( _map *Map ) {
/*
	new_dir, err = es.DirTowardNearest( _map, es.HeadPos(), '*' )
	if err == nil {
		es.SetDir( new_dir )
	} else {
		es.Turn( rand.Intn( 4 ))
	}
	es.Advance( _map )
*/
}
