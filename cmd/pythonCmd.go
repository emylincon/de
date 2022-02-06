/*
Copyright © 2022 Emeka Ugwuanyi <emylincon@gmail.com>

*/

package cmd

import (
	"github.com/emylincon/dec/cmd/environment"
	"github.com/spf13/cobra"
)

// pythonCmd represents the python command
var pythonCmd = &cobra.Command{
	Use:   "python",
	Short: "dec subcommand",
	Long:  `dec subcommand`,
}

var pythonVenvMgr = environment.NewPythonVenvMgr()

func init() {
	rootCmd.AddCommand(pythonCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pythonCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pythonCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
