package snakelib

import (
	"errors"
	"fmt"
)

type posAndDist struct {
	IntPos
	dist int
}

type PathPlanner struct {
	size IntPos
	dist_from_start_map []int
	expansion_queue chan posAndDist
	goal_pos IntPos
	found bool
}

const NOT_VISITED int = -1

func (pp *PathPlanner) queueIsEmpty() bool {
	return len( pp.expansion_queue ) == 0
}

func (pp *PathPlanner) clearQueue() {
	for len( pp.expansion_queue ) > 0 {
		<-pp.expansion_queue
	}
}

func (pp *PathPlanner) pushPos( pos IntPos, dist int ) {
	var pad posAndDist
	pad.IntPos = pos
	pad.dist = dist

	pp.expansion_queue <- pad

	pp.dist_from_start_map[ pos.X + pp.size.X * pos.Y ] = dist
}

func (pp *PathPlanner) popPos() posAndDist {
	return <-pp.expansion_queue
}

func (pp *PathPlanner) fillDistanceMap( _map *Map, start IntPos, target rune ) {
	// resize the distance field and expansion queue if needed.
	pp.size = _map.GetSize()
	new_cell_count := pp.size.X * pp.size.Y
	if len( pp.dist_from_start_map ) != new_cell_count {
		// make() returns a zeroed slice
		pp.dist_from_start_map = make( []int, new_cell_count )
		pp.expansion_queue = make( chan posAndDist, new_cell_count / 2 )
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
		new_dist := current.dist + 1
		for _, motion := range( motions ) {
			new_pos := current.IntPos.Plus( motion )
			if _map.PosValid( new_pos ) {
				cell_char := _map.GetCell( new_pos )
				if (cell_char == ' ' || cell_char == target) &&
					pp.dist_from_start_map[ new_pos.Y * pp.size.X + new_pos.X ] == NOT_VISITED {

					pp.pushPos( new_pos, new_dist )

					if _map.GetCell( new_pos ) == target {
						pp.found = true
						pp.goal_pos = new_pos
					}
				}
			}
		}	
	}
}

func (pp *PathPlanner) DirTowardsNearest( _map *Map, start IntPos, target rune ) ( Direction, error ) {
	// First, fill the distance map from the start
	pp.fillDistanceMap( _map, start, target )

	// If we didn't find the target, fail.
	if !pp.found {
		return Left, errors.New( fmt.Sprintf( "No target '%c' found reachable from %d,%d in map %s",
			target, start.X, start.Y, _map.filename ))
	}

	// If we did find the target, start there, then walk down the
	// gradient back to 1 step away from the start.
	current := pp.goal_pos
	cur_dist := pp.dist_from_start_map[ current.X + pp.size.X * current.Y ]
	var walk_dir int
	for cur_dist > 0 {
		next_dist := cur_dist - 1
		var motion IntPos
		for walk_dir, motion = range( motions ) {
			new_pos := current.Plus( motion )
			if _map.PosValid( new_pos ) &&
				pp.dist_from_start_map[ new_pos.X + pp.size.X * new_pos.Y ] == next_dist {
		
				current = new_pos
				cur_dist = next_dist
				break
			}
		}
	}

	if cur_dist != 0 {
		return Left, errors.New( fmt.Sprintf( "Walking back, ended with cur_dist = %d instead of 0.", cur_dist ))
	}

	// Direction to go towards target is the opposite of the last
	// direction we walked, since we walked from target back to
	// start.
	return Direction( (walk_dir + 2) % 4 ), nil

}
