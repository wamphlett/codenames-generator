package commands

import (
	"fmt"

	"github.com/wamphlett/codenames-generator/pkg/generator"

	"github.com/wamphlett/codenames-generator/pkg/codenames"

	"github.com/spf13/cobra"
)

const (
	DefaultWordFile string = "wordlist.json"
)

var (
	generateCommand = &cobra.Command{
		Use:   "generate",
		Short: "generate files",
		Long:  "generates the files required to play the game",
	}

	generateWordsCommand = &cobra.Command{
		Use:   "words",
		Short: "generate word file",
		Long:  "generates the file with the words for the game",
		Run: func(cmd *cobra.Command, args []string) {
			// Shuffle some cards!
			shuffledCards := codenames.ShuffleCards(
				codenames.GetWordCardsFromFile(DefaultWordFile))

			// Lay them out in a grid
			fmt.Println(generator.GenerateNameGridCSV(*codenames.PickCards(shuffledCards, 25)))
		},
	}

	generateBoardCommand = &cobra.Command{
		Use:   "board",
		Short: "generate the board file",
		Long:  "generates the board file",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Say Whaaaaat. Not written yet. soz.")
		},
	}
)

func init() {
	generateCommand.AddCommand(generateWordsCommand, generateBoardCommand)
	rootCmd.AddCommand(generateCommand)
}
