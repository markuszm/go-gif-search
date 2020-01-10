package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var trendingRanking int

// trendingCmd represents the trending command
var trendingCmd = &cobra.Command{
	Use:   "trending",
	Short: "Retrieves trending gifs",
	Long: `Stores trending gifs. Uses --ranking flag to determine which gif to store based on the current Giphy trending gifs list`,
	Run: func(cmd *cobra.Command, args []string) {
		gif, err := giphyClient.TrendingGif(trendingRanking)
		if err != nil {
			cmd.PrintErrln(  "Error: could not retrieve the gifs \n", err)
			return
		}
		file, err := downloader.StoreFile(gif.Url, gif.Id)
		if err != nil {
			cmd.PrintErrln(  "Error: could not store gif \n", err)
			return
		}
		fmt.Println(file)
	},
}

func init() {
	rootCmd.AddCommand(trendingCmd)

	trendingCmd.Flags().IntVarP(&trendingRanking, "ranking", "r", 0, "defines rank to store from trending gifs")
}
