package snakelib

import (
	"fmt"
)

type posAndDist struct {
	IntPos
	dist int
}

type PathPlanner struct {
	size IntPos
	dist_from_start_map []int
	expansion_queue []posAndDist
	expansion_queue_front int
	expansion_queue_back int
	goal_pos IntPos
	found bool
}

const NOT_VISITED int = -1

func (pp *PathPlanner) queueIsEmpty() bool {
	return pp.expansion_queue_front == pp.expansion_queue_back
}

func (pp *PathPlanner) clearQueue() {
	pp.expansion_queue_front = 0
	pp.expansion_queue_back = 0
}

func (pp *PathPlanner) pushPos( pos IntPos, dist int ) {
	var pad posAndDist
	pad.IntPos = pos
	pad.dist = dist

	pp.expansion_queue[ pp.expansion_queue_back ] = pad

	pp.expansion_queue_back++
	if pp.expansion_queue_back >= len( pp.expansion_queue ) {
		pp.expansion_queue_back = 0
	}
}

func (pp *PathPlanner) popPos() posAndDist {
	result := pp.expansion_queue[ pp.expansion_queue_front ]

	pp.expansion_queue_front++
	if pp.expansion_queue_front >= len( pp.expansion_queue ) {
		pp.expansion_queue_front = 0
	}
	return result
}

func (pp *PathPlanner) fillDistanceMap( _map *Map, start IntPos, target rune ) {
	// resize the distance field and expansion queue if needed.
	pp.size = _map.GetSize()
	new_cell_count := pp.size.X * pp.size.Y
	if len( pp.dist_from_start_map ) != new_cell_count {
		// make() returns a zeroed slice
		pp.dist_from_start_map = make( []int, new_cell_count )
		pp.expansion_queue = make( []posAndDist, new_cell_count / 2 )
	}
	// clear the dist map to NOT_VISITED in every cell
	for i, _ := range( pp.dist_from_start_map ) {
		pp.dist_from_start_map[ i ] = NOT_VISITED
	}

	pp.clearQueue()

	// grow a distance field out from start to the first instance of target in _map
	pp.pushPos( start, 0 )
	pp.found = false
	for !pp.found && !pp.queueIsEmpty() {
		current := pp.popPos()
		dist_index := current.X + pp.size.X * current.Y
		map_char := _map.GetCell( current.IntPos )
		dist_value := pp.dist_from_start_map[ dist_index ]
		fmt.Printf( "in loop, current = %d, %d dist %d, map_char = '%c', dist_value = %d\n",
			current.X, current.Y, current.dist, map_char, dist_value )
		if map_char == target {
			fmt.Printf("found target\n")
			pp.found = true
			pp.goal_pos = current.IntPos
		}
		pp.dist_from_start_map[ dist_index ] = current.dist
		new_dist := current.dist + 1
		for _, motion := range( motions ) {
			new_pos := current.IntPos.Plus( motion )
			fmt.Printf("expanding motion to new position %d,%d\n", new_pos.X, new_pos.Y )
			if _map.PosValid( new_pos ) {
				cell_char := _map.GetCell( new_pos )
				if (cell_char == ' ' || cell_char == target) &&
					pp.dist_from_start_map[ new_pos.Y * pp.size.X + new_pos.X ] == NOT_VISITED {

					fmt.Printf("pushing new pos with dist %d.\n", new_dist )
					pp.pushPos( new_pos, new_dist )
				}
			}
		}	
	}

	// walk back down the distance field
}

