package cmd

import (
	"fmt"
	"os"

	"github.com/gkwa/galaxygoat"
	"github.com/spf13/cobra"
)

var elementsToRemove string

// rootCmd represents the base command
var rootCmd = &cobra.Command{
	Use:   "galaxygoat",
	Short: "A tool to remove specific HTML elements",
	Long: `GalaxyGoat is a utility that reads HTML from stdin, 
removes specified elements, and outputs the modified HTML to stdout.

Examples:
  cat input.html | galaxygoat > output.html
  cat input.html | galaxygoat --remove="path,svg,script" > output.html`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Process HTML directly using the public API
		return galaxygoat.RemoveElementsFromReader(os.Stdin, elementsToRemove, os.Stdout)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	// Define flags
	rootCmd.Flags().StringVarP(&elementsToRemove, "remove", "r", "path", "Comma-separated list of HTML elements to remove")
}
