package cmd

import (
	"github.com/markuszm/go-gif-search/lib"
	"github.com/spf13/cobra"

	"github.com/manifoldco/promptui"
)

// selectCmd represents the select command
var selectCmd = &cobra.Command{
	Use:   "select",
	Short: "Select trending gif from list",
	Long: `Retrieves trending gifs from Giphy, shows you the list. You can select one of the gifs and the gif is stored in the folder specified with --folder flag`,
	Run: func(cmd *cobra.Command, args []string) {
		gifs, err := giphyClient.TrendingGifs()
		if err != nil {
			cmd.PrintErrln(  "Error: could not retrieve gifs \n", err)
			return
		}

		var gifNames []string
		for _, g := range gifs {
			gifNames = append(gifNames, g.Name)
		}

		prompt := promptui.Select{
			Label: "Select Gif from list",
			Items: gifNames,
		}

		_, name, err := prompt.Run()
		if err != nil {
			cmd.PrintErrln("Prompt failed", err)
			return
		}

		gif := searchGifByName(gifs, name)
		file, err := downloader.StoreFile(gif.Url, gif.Id)
		if err != nil {
			cmd.PrintErrln(  "Error: could not store gif \n", err)
			return
		}
		cmd.PrintErrln(  "stored gif in file:", file)
	},
}

func searchGifByName(gifs []lib.Gif, name string) lib.Gif {
	for _, gif := range gifs {
		if name == gif.Name {
			return gif
		}
	}
	return lib.Gif{}
}

func init() {
	trendingCmd.AddCommand(selectCmd)
}
