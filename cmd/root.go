/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
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
	Run: func(cmd *cobra.Command, args []string) {
		initializeClients()

		if len(args) == 1 {
			keyword := args[0]
			err := downloadGifForKeyword(keyword)
			if err != nil {
				log.Fatal(err)
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
			log.Fatal(err)
		}
	},
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
	fmt.Printf("downloaded gif to %s", filePath)
	return nil
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-gif-search.yaml)")

	rootCmd.PersistentFlags().IntVar(&limit, "limits", 20, "limit amount of gifs to retrieve")

	rootCmd.PersistentFlags().StringVar(&folder, "folder", os.TempDir(), "folder to store gifs in")

	rootCmd.PersistentFlags().StringVar(&apiKey, "apiKey", "", "giphy api key")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initializeClients() {
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
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".go-gif-search" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".go-gif-search")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
