package cmd

import (
	"fmt"
	"os"

	"github.com/grvcoelho/nosce/metadata"
	"github.com/spf13/cobra"
)

func init() {
	Root.AddCommand(GetCommand)

	Root.PersistentFlags().StringP("endpoint", "e", "http://169.254.169.254", "Endpoint where the information will be fetched")
}

var GetCommand = &cobra.Command{
	Use:   "get [key]",
	Short: "Get a piece of metadata information",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			cmd.Usage()
			os.Exit(1)
		}

		endpoint := Root.PersistentFlags().Lookup("endpoint").Value.String()
		metadata := metadata.New(endpoint)

		key := args[0]
		value, err := metadata.Get(key)

		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(0)
		}

		fmt.Println(value)
	},
}
