/*
Copyright © 2022 Emeka Ugwuanyi <emylincon@gmail.com>

*/

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// activateCmd represents the activate command
var activateCmd = &cobra.Command{
	Use:   "activate",
	Short: "Activate python environment",
	Long:  `Activate python venv environment in the current directory`,
	RunE: func(cmd *cobra.Command, args []string) error {
		response, err := pythonVenvMgr.Activate()
		if err == nil {
			fmt.Println(response)
		}
		return err
	},
}

func init() {
	pythonCmd.AddCommand(activateCmd)
}
