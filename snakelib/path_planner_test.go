package snakelib

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

func testDistMap(t *testing.T, pp *PathPlanner, expected []int) {
	if pp.size.X*pp.size.Y != len(expected) {
		t.Errorf("Expected slice is a different size (%d) than planner: %d x %d.\n",
			len(expected), pp.size.X, pp.size.Y)
	}
	for i, d := range expected {
		if d != pp.dist_from_start_map[i] {
			t.Errorf("At position %d, %d expected dist %d but found %d.\n",
				i%pp.size.X, i/pp.size.X, d, pp.dist_from_start_map[i])
		}
	}
}

func printDistMap(pp *PathPlanner) {
	for y := 0; y < pp.size.Y; y++ {
		for x := 0; x < pp.size.X; x++ {
			fmt.Printf("%2d, ", pp.dist_from_start_map[y*pp.size.X+x])
		}
		fmt.Printf("\n")
	}
}

func TestDistanceMapSimple(t *testing.T) {
	mapstr :=
		`##########
#  #     #
# @#  #  #
#     #  #
#  ####  #
#        #
#       !#
##########`

	reader := strings.NewReader(mapstr)
	scanner := bufio.NewScanner(reader)

	_map, err := LoadNewMap(scanner, "test_map")
	if err != nil {
		t.Errorf("Failed to load test map: %v", err)
	}

	pp := new(PathPlanner)
	pp.fillDistanceMap(_map, IntPos{2, 2}, '!')

	printDistMap(pp)

	testDistMap(t, pp, []int{
		-1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
		-1, 2, 1, -1, 5, 6, 7, 8, 9, -1,
		-1, 1, 0, -1, 4, 5, -1, 9, 10, -1,
		-1, 2, 1, 2, 3, 4, -1, 10, -1, -1,
		-1, 3, 2, -1, -1, -1, -1, 9, 10, -1,
		-1, 4, 3, 4, 5, 6, 7, 8, 9, -1,
		-1, 5, 4, 5, 6, 7, 8, 9, 10, -1,
		-1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
	})
}

func TestDirTowardsNearest1(t *testing.T) {
	mapstr :=
		`##########
#  #  *  #
#  #  #  #
# @   #  #
#  ####  #
#   *    #
#       *#
##########`

	reader := strings.NewReader(mapstr)
	scanner := bufio.NewScanner(reader)

	_map, err := LoadNewMap(scanner, "test_map")
	if err != nil {
		t.Errorf("Failed to load test map: %v", err)
	}

	pp := new(PathPlanner)
	dir, err := pp.DirTowardsNearest(_map, IntPos{2, 3}, '*')
	if Down != dir {
		t.Errorf("Direction towards nearest should be Down (3), not %d.\n", int(dir))
	}
}

func TestDirTowardsNearest2(t *testing.T) {
	mapstr :=
		`     #
  @* #
     #`
	reader := strings.NewReader(mapstr)
	scanner := bufio.NewScanner(reader)

	_map, err := LoadNewMap(scanner, "test_map")
	if err != nil {
		t.Errorf("Failed to load test map: %v", err)
	}

	pp := new(PathPlanner)
	dir, err := pp.DirTowardsNearest(_map, IntPos{2, 1}, '*')
	printDistMap(pp)
	if Right != dir {
		t.Errorf("Direction towards nearest should be Right (0), not %d.\n", int(dir))
	}

}

func TestDirTowardsNearest3(t *testing.T) {
	mapstr := "*@ *"

	reader := strings.NewReader(mapstr)
	scanner := bufio.NewScanner(reader)

	_map, err := LoadNewMap(scanner, "test_map")
	if err != nil {
		t.Errorf("Failed to load test map: %v", err)
	}

	pp := new(PathPlanner)
	dir, err := pp.DirTowardsNearest(_map, IntPos{1, 0}, '*')
	if Left != dir {
		t.Errorf("Direction towards nearest should be Left (2), not %d.\n", int(dir))
	}

}

func TestDirTowardsNearestBig(t *testing.T) {
	mapstr :=
		`#############################################################################################################################################################
#                                                                                                                                                           #
#                                                                                                                                                           #
#                                                                                                                                                           #
#                                                                                                                                                           #
#                                                                                                                                                           #
#                                                                                                                                                           #
#                                                                                                                                                           #
#                                                                                                                                                           #
#                                                                                                                                                           #
#                                                                                                                                                           #
#                                                                                                                                                           #
#                       #                                                                                                    #                              #
#                       #                                                 *                                                  #                              #
#                       #                                                                                                    #                              #
#                       #                                                                                                    #                              #
#                       #                                                                                                    #                              #
#                       #                                                                                                    #                              #
#                       #                                                                                                    #                              #
#                       #                                                                                                    #                              #
#                       #                                                                                                    #                              #
#                       #                                                                                                    #                              #
#                       #                                                                                                    #                              #
#                       #                                                                                                    #                              #
#                       #                                                                                                    #                              #
#                       #                                                                                                    #                              #
#                       #                                                                                                    #                              #
#                       ######################################################################################################                              #
#                       #&                                                                                                   #                              #
#                       #                                                                                                    #                              #
#                       #                                                                                                    #                              #
#                       #                                                                                                    #                              #
#                       #                                                                                                    #                              #
#                       #                                                                                                    #                              #
#                       #                                                                                                    #                              #
#                       #                                                                                                    #                              #
#                       #                                                                                                    #                              #
#                       #                                                                                                    #                              #
#                       #                                                                                                    #                              #
#                       #                                                                                                    #                              #
#                       #                                                                                                    #                              #
#                       #                                                                                                    #                              #
#                       #                                                                                                    #                              #
#                       #                                                                                                    #                              #
#                       #                                                                                                    #                              #
#                                                                                                                                                           #
#                                                                                                                                                           #
#                                                                                                                                                           #
#                                                                                                                                                           #
#                                                                                                                                                           #
#                                                                                                                                                           #
#                                                                                                                                                           #
#                                                                                                                                                           #
#                                                                                                                                                           #
#                                                                                                                                                           #
#                                                                                                                                                           #
#                                                                                                                                                           #
#############################################################################################################################################################
`

	reader := strings.NewReader(mapstr)
	scanner := bufio.NewScanner(reader)

	_map, err := LoadNewMap(scanner, "test_map")
	if err != nil {
		t.Errorf("Failed to load test map: %v", err)
	}

	start, err := _map.Find('&')
	if err != nil {
		t.Errorf("Failed to find '&': %v", err)
	}

	pp := new(PathPlanner)
	dir, err := pp.DirTowardsNearest(_map, start, '*')
	if Down != dir {
		t.Errorf("Direction towards nearest should be Down (3), not %d.\n", int(dir))
	}
}
