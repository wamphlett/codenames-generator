package generator

import (
	"fmt"
	"os"
	"strings"

	"github.com/wamphlett/codenames-generator/pkg/codenames"
)

func GenerateNameGridCSV(cards []codenames.WordCard) (grid string) {
	// You must provide 25 cards to generate a valid grid
	if len(cards) != 25 {
		fmt.Println("You must provide 25 cards to generate a valid grid!")
		os.Exit(1)
	}

	values := []string{}
	for _, c := range cards {
		values = append(values, c.String())
	}

	for i := 0; i < 25; i += 5 {
		grid = grid + strings.Join(values[i:i+5], ",") + "\n"
	}

	return
}
