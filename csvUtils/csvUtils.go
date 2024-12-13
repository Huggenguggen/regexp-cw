package csvUtils

import (
	"encoding/csv"
	"fmt"
	"os"
	gm "regexp-cw/gamemodel"
)

func parseProblem(input string) ([]gm.Header, []gm.Header, error) {
	f, err := os.Open(input)
	if err != nil {
		return nil, nil, fmt.Errorf("Unable to read input file "+input, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to parse CSV: %v", err)
	}

	if len(records) != 2 {
		return nil, nil, fmt.Errorf("CSV must contain exactly 2 rows")
	}

	rows := records[0]
	cols := records[1]

	Rows := make([]gm.Header, len(rows))
	Cols := make([]gm.Header, len(cols))

	return Rows, Cols, nil
}
