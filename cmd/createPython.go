/*
Copyright © 2022 Emeka Ugwuanyi <emylincon@gmail.com>

*/

package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// createPythonCmd represents the create command
var createPythonCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a basic python app template dev",
	Long:  `Creates a basic python app template dev`,
	RunE: func(cmd *cobra.Command, args []string) error {
		dir, err := cmd.Flags().GetString("directory")
		if err != nil {
			return err
		}
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}
		email, err := cmd.Flags().GetString("email")
		if err != nil {
			return err
		}
		license, err := cmd.Flags().GetString("license")
		if err != nil {
			return err
		}

		return python(dir, name, email, license)

	},
}

func init() {
	pythonCmd.AddCommand(createPythonCmd)
	createPythonCmd.Flags().StringP("directory", "d", "", "directory for go app")
	createPythonCmd.Flags().StringP("name", "n", "name", "name of developer")
	createPythonCmd.Flags().StringP("email", "e", "email.gmail.com", "email address of developer")
	createPythonCmd.Flags().StringP("license", "l", "MIT", "license for code")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func python(directory, name, email, license string) error {
	if err := os.Mkdir(directory, os.ModePerm); err != nil {
		return err
	}
	// create virtual environment
	err := pythonVenvMgr.Create(directory)
	if err != nil {
		return err
	}
	err = createEnvironment(directory, name, email, "python")
	if err != nil {
		return err
	}
	return setUpLicense(directory, license, name)
}
