package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var phrase string

// translateCmd represents the translate command
var translateCmd = &cobra.Command{
	Use:   "translate",
	Short: "Translates a phrase to a gif",
	Long: "Translates phrases to gifs using a weirdness factor." +
		"Uses Giphy API and stores the gif in the folder specified by the --folder flag.",
	Run: func(cmd *cobra.Command, args [] string) {
		initializeClients()

		gif, err := giphyClient.TranslateGif(phrase)
		if err != nil {
			fmt.Fprintf(os.Stderr, "did not find a gif for your phrase")
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
	rootCmd.AddCommand(translateCmd)

	translateCmd.Flags().StringVarP(&phrase, "phrase", "p", "", "defines phrase to search for in giphy translate")
}
