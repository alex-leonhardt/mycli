package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "generate a new name, take stdin as prefix",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var input []byte
		var err error
		var in os.FileInfo

		if in, _ = os.Stdin.Stat(); in.Mode()&os.ModeCharDevice == 0 {
			input, err = ioutil.ReadAll(os.Stdin)
			if err != nil {
				log.Fatal("booyaa")
			}
		}
		fmt.Printf("%s%s%s\n", strings.TrimSuffix(string(input), "\n"), func() string {
			if input != nil {
				return "-"
			}
			return ""
		}(), uuid.New())
	},
}

func init() {
	nameCmd.AddCommand(newCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
