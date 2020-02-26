package cmd 

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command {
		Use: "ore-wa-go",
		Short: "Kotlin transpiler to golang",
		Long: "OreWaGo is a CLI for transpiled from kotlin to golang",
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() { 
	rootCmd.AddCommand(helloCmd)
}