package snakelib

import (
	"bufio"
	"strings"
	"testing"
)

func testCell(t *testing.T, m *Map, x, y int, expected rune) {
	actual := m.GetCell(IntPos{x, y})
	if actual != expected {
		t.Errorf("Map cell %d, %d was '%c' but should have been '%c'\n",
			x, y, actual, expected)
	}
}

func TestPosValid(t *testing.T) {
	var m Map
	m.size = IntPos{2, 3}
	if true != m.PosValid(IntPos{1, 2}) {
		t.Errorf("1,2 should be a valid position, but PosValid() says it is not.\n")
	}
	if true != m.PosValid(IntPos{0, 2}) {
		t.Errorf("0,2 should be a valid position, but PosValid() says it is not.\n")
	}
	if false != m.PosValid(IntPos{-1, 2}) {
		t.Errorf("-1,2 should be an invalid position, but PosValid() says it is valid.\n")
	}
	if false != m.PosValid(IntPos{2, 2}) {
		t.Errorf("2,2 should be an invalid position, but PosValid() says it is valid.\n")
	}
	if false != m.PosValid(IntPos{1, 3}) {
		t.Errorf("1,3 should be an invalid position, but PosValid() says it is valid.\n")
	}
	if false != m.PosValid(IntPos{1, -1}) {
		t.Errorf("1,-1 should be an invalid position, but PosValid() says it is valid.\n")
	}
}

func TestNewEmptyMap(t *testing.T) {
	m := NewEmptyMap(IntPos{4, 3})
	testCell(t, m, 0, 0, '#')
	testCell(t, m, 0, 1, '#')
	testCell(t, m, 2, 2, '#')
	testCell(t, m, 3, 2, '#')
	testCell(t, m, 1, 0, '#')
	testCell(t, m, 1, 1, ' ')
	testCell(t, m, 2, 1, ' ')
}

func TestLoadMap(t *testing.T) {
	mapstr :=
		`#####
# @ #
#   #
#####`
	reader := strings.NewReader(mapstr)
	scanner := bufio.NewScanner(reader)

	m, err := LoadNewMap(scanner, "test_map")
	if err != nil {
		t.Errorf("Failed to load test map: %v", err)
	}
	testCell(t, m, 0, 0, '#')
	testCell(t, m, 2, 1, '@')
	testCell(t, m, 1, 2, ' ')
	testCell(t, m, 4, 3, '#')
	if 5 != m.size.X {
		t.Errorf("Map width is %d but should have been 5.\n", m.size.X)
	}
	if 4 != m.size.Y {
		t.Errorf("Map height is %d but should have been 4.\n", m.size.Y)
	}
}
