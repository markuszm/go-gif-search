package cmd

import (
	"github.com/spf13/cobra"
)

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "Loads a random gif from Giphy",
	Long: `Downloads and stores a random gif from Giphy to the folder specified with --folder flag`,
	Run: func(cmd *cobra.Command, args []string) {
		gif, err := giphyClient.RandomGif()
		if err != nil {
			cmd.PrintErrln( "Error: could not retrieve a gif \n", err)
			return
		}
		file, err := downloader.StoreFile(gif.Url, gif.Id)
		if err != nil {
			cmd.PrintErrln("Error: could not store gif \n", err)
			return
		}
		cmd.PrintErrln( "stored gif in file:", file)
	},
}

func init() {
	rootCmd.AddCommand(randomCmd)
}
