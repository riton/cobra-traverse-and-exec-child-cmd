package cmd

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// allCmd represents the all command
var allCmd = &cobra.Command{
	Use:   "all",
	Short: "Build all software",
	RunE:  buildAllRun,
}

func init() {
	buildCmd.AddCommand(allCmd)
}

// buildAllRun should discover all 'build' child commands
// other than 'all' and execute them.
func buildAllRun(cmd *cobra.Command, args []string) error {
	// Get all build commands, except the 'all' command to avoid infinite recursion
	var buildSoftwareCmds []*cobra.Command
	for _, builderCmd := range buildCmd.Commands() {
		if builderCmd != cmd {
			fmt.Printf("[buildAllRun] discovered cmd.Use = %q\n", builderCmd.Use)
			buildSoftwareCmds = append(buildSoftwareCmds, builderCmd)
		}
	}

	for _, builderCmd := range buildSoftwareCmds {
		fmt.Printf("[buildAllRun] executing command = %q\n", builderCmd.Use)
		//
		// Problem: This 'Execute()' call will cause infinite recursion :-(
		//
		if err := builderCmd.Execute(); err != nil {
			return errors.Wrapf(err, "running command %s", builderCmd.Use)
		}
	}

	return nil
}
