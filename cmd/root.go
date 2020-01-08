package cmd

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"strings"

	"github.com/markuszm/go-gif-search/lib"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

var giphyClient *lib.GiphyClient

var downloader *lib.Downloader

var apiKey string

var limit int

var folder string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-gif-search",
	Short: "Search gifs based on a keyword",
	Long: "Searches gifs based on a keyword and stores the first one to folder specified with --folder flag. Either uses first argument or piped input",
	Run: func(cmd *cobra.Command, args []string) {
		initializeClients()

		if len(args) == 1 {
			keyword := args[0]
			err := downloadGifForKeyword(keyword)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: %s", err)
			}
			return
		}

		input, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			return
		}
		keyword := strings.TrimSpace(string(input))
		err = downloadGifForKeyword(keyword)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s", err)
		}
	},
	Version: "0.0.1",
}

func downloadGifForKeyword(keyword string) error {
	gif, err := giphyClient.SearchGif(keyword, 0)
	if err != nil {
		return errors.Wrap(err, "error searching gif")
	}
	if gif.Id == "" {
		return errors.Wrap(err, "no gif found")
	}
	filePath, err := downloader.StoreFile(gif.Url, gif.Id)
	if err != nil {
		return errors.Wrap(err, "error downloading gif")
	}
	fmt.Fprintf(os.Stderr, "downloaded gif to %s", filePath)
	return nil
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr,  err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-gif-search.yaml)")

	rootCmd.PersistentFlags().IntVar(&limit, "limits", 20, "limit amount of gifs to retrieve")

	rootCmd.PersistentFlags().StringVar(&folder, "folder", os.TempDir(), "folder to store gifs in")

	rootCmd.PersistentFlags().StringVar(&apiKey, "apiKey", "", "giphy api key")

	rootCmd.SetVersionTemplate(`{{with .Name}}{{printf "The Cool Gif Search - %s " .}}{{end}}{{printf "Version: %s" .Version}}`)
}

func initializeClients() {
	if apiKey == "" {
		fmt.Fprint(os.Stderr, "you need to specify a API Key")
		os.Exit(1)
	}

	giphyClient = lib.NewGiphyClient(apiKey, limit)
	downloader = &lib.Downloader{Folder: folder}
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		// Search config in home directory with name ".go-gif-search" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".go-gif-search")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
