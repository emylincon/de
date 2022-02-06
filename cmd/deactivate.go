/*
Copyright © 2022 Emeka Ugwuanyi <emylincon@gmail.com>

*/

package cmd

import (
	"github.com/spf13/cobra"
)

// deactivateCmd represents the deactivate command
var deactivateCmd = &cobra.Command{
	Use:   "deactivate",
	Short: "Deactivate python environment",
	Long:  `Deactivate python venv environment in the current directory`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return pythonVenvMgr.Deactivate()
	},
}

func init() {
	pythonCmd.AddCommand(deactivateCmd)
}
