/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/dsmith022/hithlain/backend/library"
	"github.com/spf13/cobra"
)

var librarian *library.Librarian

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "hithlain",
	Short: "A branching text story engine, developed in Go.",
	Long: `

Hithlain

/ˈhɪθ.laɪn/ — HITH-line
— noun

1) A strong yet light thread of Elvish craft,
grey as gathered mist and soft as silk beneath the hand; 
woven in Lórien by arts unspoken, whose making was hidden from all save those who wrought it.

2) A lightweight engine for weaving branching RPG stories, choices, and game state into playable text adventures.`,
	PersistentPreRun:  loadLibrary,
	PersistentPostRun: saveLibrary,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() error {
	return rootCmd.Execute()

}

func init() {
	rootCmd.PersistentFlags().StringVar(&librarian.FilePath, "libraryPath", "", "library storage (default is $HOME/.hithlain.json)")

}

func loadLibrary(cmd *cobra.Command, args []string) {
	var err error
	librarian, err = library.LoadLibrary(librarian.GetLibraryPath())
	if err != nil {
		fmt.Fprintln(cmd.ErrOrStderr(), "Error loading library:", err)
	}
}

func saveLibrary(cmd *cobra.Command, args []string) {
	if err := librarian.SaveLibrary(); err != nil {
		fmt.Fprintln(cmd.ErrOrStderr(), "Error saving library:", err)
	}
}
