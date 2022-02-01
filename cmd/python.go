/*
Copyright © 2022 Emeka Ugwuanyi <emylincon@gmail.com>

*/

package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// pythonCmd represents the python command
var pythonCmd = &cobra.Command{
	Use:   "python",
	Short: "Creates a basic python app template dev",
	Long:  `Creates a python dev environment`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("python called")
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

		python(dir, name, email)
		return nil
	},
}

func init() {
	pythonCmd.Flags().StringP("directory", "d", "", "directory for go app")
	pythonCmd.Flags().StringP("name", "n", "name", "name of developer")
	pythonCmd.Flags().StringP("email", "e", "email.gmail.com", "email address of developer")
	rootCmd.AddCommand(pythonCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pythonCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pythonCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func python(directory, name, email string) error {
	log.Println("directory: ", directory)
	if err := os.Mkdir(directory, os.ModePerm); err != nil {
		return err
	}
	return createEnvironment(directory, name, email, "python")
}
