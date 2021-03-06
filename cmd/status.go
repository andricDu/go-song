package cmd

import (
	"fmt"
	"net/url"

	"github.com/andricdu/go-song/song"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	RootCmd.AddCommand(statusCmd)
}

func getStatus(uploadID string) {
	studyID, accessToken := viper.GetString("study"), viper.GetString("accessToken")
	songURL, err := url.Parse(viper.GetString("songURL"))
	if err != nil {
		panic(err)
	}
	client := song.CreateClient(accessToken, songURL)
	responseBody := client.GetStatus(studyID, uploadID)
	fmt.Println(string(responseBody))
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Get status of uploaded analysis",
	Long:  `Get status of uploaded analysis`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		getStatus(args[0])
	},
}
