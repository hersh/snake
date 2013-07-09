package snakelib

type PathPlanner struct {
	dist_from_start_map []int
	expansion_queue []IntPos
	expansion_queue_front int
	expansion_queue_back int
}

func (pp *PathPlanner) DirTowardNearest( _map *Map, start IntPos, target rune ) {
	// resize the distance field and expansion queue if needed.
	map_size := _map.GetSize()
	new_cell_count = map_size.X * map_size.Y
	if len( pp.dist_from_start_map ) != new_cell_count {
		// make() returns a zeroed slice
		pp.dist_from_start_map = make( []int, new_cell_count )
		pp.expansion_queue = make( []IntPos, new_cell_count / 2 )
	} else {
		// zero the existing slice
		for i, _ := range( pp.dist_from_start_map ) {
			pp.dist_from_start_map[ i ] = 0
		}
	}

	pp.expansion_queue_front = 0
	pp.expansion_queue_back = 0

	// grow a distance field out from start to the first instance of target in _map
	pp.pushPos( start )
	found := false
	while( !found && pp.expansion_queue_front != pp.expansion_queue_back ) {
		pos := pp.expansion_queue[ pp.expansion_queue_back ]
		dist_index = pos.X + map_size.X * pos.Y
		pp.expansion_queue_back++
		map_char = _map.GetCell( pos )
		dist_value := dist_from_start_map[ dist_index ]
		if map_char == target {
			found = true
			goal_pos = pos
		}
		if (map_char == ' ' || map_char == '*') && dist_value == 0 {
			dist_from_start_map
		}
	}

	// walk back down the distance field
}

	for _, motion := range( motions ) {
	}