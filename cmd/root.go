package cmd

import "github.com/spf13/cobra"

var Root = &cobra.Command{
	Use:   "nosce [command]",
	Short: "nosce - Get metadata information of your EC2 instances",
}
