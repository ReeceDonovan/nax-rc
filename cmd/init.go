package cmd

import (
	"fmt"

	"github.com/ReeceDonovan/nax-rc/pkg/ioutil"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new nax-rc project",
	Long:  `Initialize a new nax-rc project. This will create the necessary files to get started with nax-rc.`,
	RunE:  initCommand,
}

func init() {
	// TODO: Add flags to specify the project name and the project directory
	// TODO: Add constants for nax-rc file directory
}

func initCommand(cmd *cobra.Command, args []string) error {
	// Check if the current directory is a nax-rc project
	// If it is, return an error

	projectExists, err := ioutil.CheckPathExists(".nax")
	if err != nil {
		return fmt.Errorf("error checking if current directory is a nax-rc project: %w", err)
	}

	if projectExists {
		return fmt.Errorf("error creating new project: project already exists")
	}

	// Create the .nax-rc directories
	err = ioutil.CreateDirectories([]string{".nax", ".nax/objects", ".nax/refs"})
	if err != nil {
		return fmt.Errorf("error creating directories: %w", err)
	}

	fmt.Println("Initialized empty nax-rc repository in .nax")
	return nil
}
