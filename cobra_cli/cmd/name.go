package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// nameCmd represents the name command
var nameCmd = &cobra.Command{
	Use:   "name",
	Short: "name command..",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("name called")
		if len(args) == 0 {
			fmt.Println(cmd.Help())
		}
	},
}

func init() {
	rootCmd.AddCommand(nameCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	nameCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	nameCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
