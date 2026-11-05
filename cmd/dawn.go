package cmd

import (
	"fmt"

	"github.com/dsmith022/hithlain/backend/library"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var dawnCmd = &cobra.Command{
	Use:   "startup hithlain",
	Short: "startup hithlain",
	Long:  "load hithlain and all library data present at the current directory",
	Run:   dawn,
}

func dawn(cmd *cobra.Command, args []string) {
	libPath := viper.GetString("libPath")
	db, err := library.Open(libPath)
	if err != nil {
		fmt.Fprintln(cmd.ErrOrStderr(), "Error loading library:", err)
	}

}
