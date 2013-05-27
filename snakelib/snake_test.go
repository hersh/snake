package snakelib

import "testing"

func testPos( t* testing.T, message string, expected, actual IntPos ) {
	if( expected.X != actual.X || expected.Y != actual.Y ) {
		t.Errorf( message + " should be %d,%d, but is %d, %d\n",
			expected.X, expected.Y, actual.X, actual.Y )
	}
}

func TestUpdate( t* testing.T ) {
	s := NewSnake( IntPos{1, 1}, 4 )
	if s.head_dir != Right {
		t.Errorf( "s.head_dir should default to Right, but is %d.\n", s.head_dir );
	}

	s.body = []IntPos {{1, 1}, {2,2}}
	s.Update()
	testPos( t, "t1 body[0]", IntPos{2,2}, s.body[0] );
	testPos( t, "t1 body[1]", IntPos{3,2}, s.body[1] );
	s.Update()
	testPos( t, "t2 body[0]", IntPos{3,2}, s.body[0] );
	testPos( t, "t2 body[1]", IntPos{4,2}, s.body[1] );

	if( motions[ s.head_dir ] != IntPos{1,0} ) {
		t.Errorf( "motion table is wrong: should be 1, 0 but is %d,%d\n",
			motions[ s.head_dir ].X, motions[ s.head_dir ].Y )
	}
}

func TestTurn( t* testing.T ) {
	s := NewSnake( IntPos{1, 1}, 4 )
	if int( Right ) != 0 {
		t.Errorf( "Right should be 0, but is %d.\n", int( Right ))
	}
	s.Turn( -1 )
	if( s.head_dir != Down ) {
		t.Errorf( "s.head_dir should be Down but is %d.\n", int( s.head_dir ))
	}
	s.Turn( -1 )
	if( s.head_dir != Left ) {
		t.Errorf( "s.head_dir should be Left but is %d.\n", int( s.head_dir ))
	}
	s.Turn( -1 )
	if( s.head_dir != Up ) {
		t.Errorf( "s.head_dir should be Up but is %d.\n", int( s.head_dir ))
	}
	s.Turn( -1 )
	if( s.head_dir != Right ) {
		t.Errorf( "s.head_dir should be Right but is %d.\n", int( s.head_dir ))
	}

	s.Turn( 1 )
	if( s.head_dir != Up ) {
		t.Errorf( "s.head_dir should be Up but is %d.\n", int( s.head_dir ))
	}
	s.Turn( 1 )
	if( s.head_dir != Left ) {
		t.Errorf( "s.head_dir should be Left but is %d.\n", int( s.head_dir ))
	}
	s.Turn( 1 )
	if( s.head_dir != Down ) {
		t.Errorf( "s.head_dir should be Down but is %d.\n", int( s.head_dir ))
	}
	s.Turn( 1 )
	if( s.head_dir != Right ) {
		t.Errorf( "s.head_dir should be Right but is %d.\n", int( s.head_dir ))
	}
}
