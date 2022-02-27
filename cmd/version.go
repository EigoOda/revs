package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	version = "0.1"

	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version",
		Long:  `Print the version`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("version: %s\n", version)
		},
	}
)

func init() {
	rootCmd.AddCommand(versionCmd)
}
