package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "hithlain",
	Short: "Hithlain is a storytelling platform",
	Long: `A nonlinear story composition tool, ran in the CLI, and built with Go by dsmith022.

			*/hɪθ.laɪn/ — HITH-line*  
			*noun*

			1) A strong yet light thread of Elvish craft, grey in color and soft as silk to the touch; 
			woven in Lórien by arts whose making was never revealed.

			2) A lightweight engine for weaving branching stories, choices, and game state into playable text adventures.
				`,
}
