package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "Loads a random gif from Giphy",
	Long: `Downloads and stores a random gif from Giphy to the folder specified with --folder flag`,
	Run: func(cmd *cobra.Command, args []string) {
		initializeClients()

		gif, err := giphyClient.RandomGif()
		if err != nil {
			fmt.Fprintf(os.Stderr, "could not find a gif")
			return
		}
		file, err := downloader.StoreFile(gif.Url, gif.Id)
		if err != nil {
			fmt.Fprintf(os.Stderr, "could not store gif")
			return
		}
		fmt.Fprintf(os.Stderr, "stored gif in file: %s", file)
	},
}

func init() {
	rootCmd.AddCommand(randomCmd)
}
