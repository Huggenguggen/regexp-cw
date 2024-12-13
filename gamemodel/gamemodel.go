package gamemodel

type HeaderType string

const (
	ROW HeaderType = "ROW"
	COL HeaderType = "COL"
)

type Header struct {
	Expression   string
	HeaderTyping HeaderType
}

type Board struct {
	Rows       []Header
	Columns    []Header
	Row_length int
	Col_length int
	Entries    []rune
}

type Model struct {
	Board      Board  // board struct which holds the actual game
	Cursor     [2]int // which part of the board the cursor is pointing at
	IsComplete bool   // whether or not the game is complete
}
