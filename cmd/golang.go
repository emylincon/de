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

// golangCmd represents the golang command
var golangCmd = &cobra.Command{
	Use:   "golang",
	Short: "Create a basic go app template",
	Long:  `Creates a go dev environment`,
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
	golangCmd.Flags().StringP("directory", "d", "", "directory for go app")
	golangCmd.Flags().StringP("name", "n", "name", "name of developer")
	golangCmd.Flags().StringP("email", "e", "email.gmail.com", "email address of developer")
	rootCmd.AddCommand(golangCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// golangCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:

}

func golang(directory, name, email string) error {
	log.Println("directory: ", directory)
	if err := os.Mkdir(directory, os.ModePerm); err != nil {
		return err
	}

	return createEnvironment(directory, name, email, "golang")
}
