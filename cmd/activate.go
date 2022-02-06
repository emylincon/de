/*
Copyright © 2022 Emeka Ugwuanyi <emylincon@gmail.com>

*/

package cmd

import (
	"github.com/spf13/cobra"
)

// activateCmd represents the activate command
var activateCmd = &cobra.Command{
	Use:   "activate",
	Short: "Activate python environment",
	Long:  `Activate python venv environment in the current directory`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return pythonVenvMgr.Activate()
	},
}

func init() {
	pythonCmd.AddCommand(activateCmd)
}
