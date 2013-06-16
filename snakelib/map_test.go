package snakelib

import (
	"io/ioutil"
	"testing"
)

func testCell( t *testing.T, m *Map, x, y int, expected rune ) {
	actual := m.GetCell( IntPos{ x, y })
	if( actual != expected ) {
		t.Errorf( "Map cell %d, %d was '%c' but should have been '%c'\n",
			x, y, actual, expected )
	}
}

func TestNewEmptyMap( t* testing.T ) {
	m := NewEmptyMap( IntPos{ 4, 3 })
	testCell( t, m, 0, 0, '#' )
	testCell( t, m, 0, 1, '#' )
	testCell( t, m, 2, 2, '#' )
	testCell( t, m, 3, 2, '#' )
	testCell( t, m, 1, 0, '#' )
	testCell( t, m, 1, 1, ' ' )
	testCell( t, m, 2, 1, ' ' )
}

func TestLoadMap( t* testing.T ) {
	mapstr :=
`#####
# @ #
#   #
#####`
	file, err := ioutil.TempFile( "", "" )
	if err != nil {
		t.Errorf( "Failed to create tempfile.\n" )
		return
	}
	file.WriteString( mapstr )
	file.Close()

	m, err := LoadNewMap( file.Name() )
	if err != nil {
		t.Errorf( "Failed to load test map: %v", err )
	}
	testCell( t, m, 0, 0, '#' )
	testCell( t, m, 2, 1, '@' )
	testCell( t, m, 1, 2, ' ' )
	testCell( t, m, 4, 3, '#' )
	if 5 != m.size.X {
		t.Errorf( "Map width is %d but should have been 5.\n", m.size.X )
	}
	if 4 != m.size.Y {
		t.Errorf( "Map height is %d but should have been 4.\n", m.size.Y )
	}
}
