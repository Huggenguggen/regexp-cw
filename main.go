package main

import (
	"os"
	gm "regexp-cw/gamemodel"

	tea "github.com/charmbracelet/bubbletea"
)

func initialModel() gm.Model {
	Rows, Cols, err := csvUtils.parseProblem("default.csv")
	if err != nil {
		os.Exit(1)
	}
	return gm.Model{
		// Our to-do list is a grocery list
		choices: []string{"Buy carrots", "Buy celery", "Buy kohlrabi"},

		// A map which indicates which choices are selected. We're using
		// the  map like a mathematical set. The keys refer to the indexes
		// of the `choices` slice, above.
		selected: make(map[int]struct{}),
	}
}

func (m gm.Model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}
