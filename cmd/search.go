package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var searchRanking int
var keyword string

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Searches a gif using keyword",
	Long: `Searches a gif on Giphy using the keyword and stores it in the folder specified with --folder flag
	Returns first one or the one specified by --ranking flag`,
	Run: func(cmd *cobra.Command, args []string) {
		initializeClients()

		gif, err := giphyClient.SearchGif(keyword, searchRanking)
		if err != nil {
			fmt.Fprintf(os.Stderr, "did not find a gif for your keyword")
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
	rootCmd.AddCommand(searchCmd)

	searchCmd.Flags().StringVarP(&keyword, "keyword", "k", "", "keyword to use in search")
	searchCmd.Flags().IntVarP(&searchRanking, "ranking", "r", 0, "defines which rank of a gif list to store")
}
