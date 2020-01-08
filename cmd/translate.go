package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var phrase string

// translateCmd represents the translate command
var translateCmd = &cobra.Command{
	Use:   "translate",
	Short: "Translates a phrase to a gif",
	Long: `Translates phrases to gifs using a weirdness factor.
		Uses Giphy API and stores the gif in the folder specified by the --folder flag.`,
	Run: func(cmd *cobra.Command, args [] string) {
		gif, err := giphyClient.TranslateGif(phrase)
		if err != nil {
			cmd.PrintErrln(  fmt.Sprintf("Error: did not find a gif for your phrase: %s \n %s", phrase, err))
			return
		}
		file, err := downloader.StoreFile(gif.Url, gif.Id)
		if err != nil {
			cmd.PrintErrln( "Error: could not store gif \n", err)
			return
		}
		cmd.PrintErrln(  "stored gif in file:", file)
	},
}

func init() {
	rootCmd.AddCommand(translateCmd)

	translateCmd.Flags().StringVarP(&phrase, "phrase", "p", "", "defines phrase to search for in giphy translate")
}
