package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var trendingRanking int

// trendingCmd represents the trending command
var trendingCmd = &cobra.Command{
	Use:   "trending",
	Short: "Retrieves trending gifs",
	Long: `Stores trending gifs. Uses --ranking flag to determine which gif to store based on the current Giphy trending gifs list`,
	Run: func(cmd *cobra.Command, args []string) {
		initializeClients()

		gif, err := giphyClient.TrendingGif(trendingRanking)
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
	rootCmd.AddCommand(trendingCmd)

	trendingCmd.Flags().IntVarP(&trendingRanking, "ranking", "r", 0, "defines rank to store from trending gifs")
}
