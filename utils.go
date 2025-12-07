package adventOfCode

import (
	"fmt"
	"os"
	"path/filepath"
)

func OpenInput(elem ...string) (*os.File, error) {
	if f, err := os.Open(filepath.Join(elem...)); err == nil {

		return f, nil
	}
	return nil, fmt.Errorf("could not find input file")
}
