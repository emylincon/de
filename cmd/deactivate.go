/*
Copyright © 2022 Emeka Ugwuanyi <emylincon@gmail.com>

*/

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// deactivateCmd represents the deactivate command
var deactivateCmd = &cobra.Command{
	Use:   "deactivate",
	Short: "Deactivate python environment",
	Long:  `Deactivate python venv environment in the current directory`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(pythonVenvMgr.Deactivate())
	},
}

func init() {
	pythonCmd.AddCommand(deactivateCmd)
}
