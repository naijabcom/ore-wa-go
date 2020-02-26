package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var helloCmd = &cobra.Command {
		Use: "hello",
		Short: "A hello world",
		Long: "A hello world long!",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Hello => %s", args)
		},
	}

func init() {
	rootCmd.AddCommand(helloCmd)
}