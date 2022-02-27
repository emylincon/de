/*
Copyright © 2022 Emeka Ugwuanyi <emylincon@gmail.com>

*/

package cmd

import (
	"fmt"
	"os"

	"github.com/emylincon/dec/cmd/environment"
	"github.com/spf13/cobra"
)

// createGolangCmd represents the create command
var createGolangCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a basic go app template",
	Long:  `Create a basic go app template`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("golang called")
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

		golang(dir, name, email)
		return nil
	},
}

func init() {
	golangCmd.AddCommand(createGolangCmd)
	createGolangCmd.Flags().StringP("directory", "d", "", "directory for go app")
	createGolangCmd.Flags().StringP("name", "n", "name", "name of developer")
	createGolangCmd.Flags().StringP("email", "e", "email.gmail.com", "email address of developer")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func golang(directory, name, email string) error {
	if err := os.Mkdir(directory, os.ModePerm); err != nil {
		return err
	}
	err := createEnvironment(directory, name, email, "golang")
	if err != nil {
		return err
	}
	return environment.GoModTidy(directory)

}
