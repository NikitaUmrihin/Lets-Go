package structs

// score keeps track of wins and losses while playing the game.
type score struct {
	Wins  int
	Loses int
}

// NewScore initializes a new instance of the score struct.
func NewScore() *score {
	return &score{
		Wins:  0,
		Loses: 0,
	}
}
