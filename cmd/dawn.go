package cmd

import (
	"github.com/spf13/cobra"
)

var dawnCmd = &cobra.Command{
	Use:   "startup hithlain",
	Short: "startup hithlain",
	Long:  "load hithlain and all library data present at the current directory",
	Run:   dawn,
}

func dawn(cmd *cobra.Command, args []string) {

}
