package cmd

import (
	"github.com/spf13/cobra"
)

// keysCmd represents the keys command
var keysCmd = &cobra.Command{
	Use: "keys",
}

func init() {
	searchCmd.AddCommand(keysCmd)
}
