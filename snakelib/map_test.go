package snakelib

import "testing"

func testCell( t *testing.T, m *Map, x, y int, expected rune ) {
	actual := m.GetCell( x, y )
	if( actual != expected ) {
		t.Errorf( "Map cell %d, %d was '%c' but should have been '%c'\n",
			x, y, actual, expected )
	}
}

func TestNewEmptyMap( t* testing.T ) {
	m := NewEmptyMap( 4, 3 )
	testCell( t, m, 0, 0, '#' )
	testCell( t, m, 0, 1, '#' )
	testCell( t, m, 2, 2, '#' )
	testCell( t, m, 3, 2, '#' )
	testCell( t, m, 1, 0, '#' )
	testCell( t, m, 1, 1, ' ' )
	testCell( t, m, 2, 1, ' ' )
}
