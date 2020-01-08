package cmd

import (
	"fmt"
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
		gif, err := giphyClient.SearchGif(keyword, searchRanking)
		if err != nil {
			cmd.PrintErrln( fmt.Sprintf("Error: did not find a gif for your keyword %s \n %s", keyword, err))
			return
		}
		file, err := downloader.StoreFile(gif.Url, gif.Id)
		if err != nil {
			cmd.PrintErrln( "Error: could not store gif \n", err)
			return
		}
		cmd.PrintErrln( "stored gif in file:", file)
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)

	searchCmd.Flags().StringVarP(&keyword, "keyword", "k", "", "keyword to use in search")
	searchCmd.Flags().IntVarP(&searchRanking, "ranking", "r", 0, "defines which rank of a gif list to store")
}
