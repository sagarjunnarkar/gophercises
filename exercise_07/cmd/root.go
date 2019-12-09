package cmd

import "github.com/spf13/cobra"

//RootCmd is var holding Cobra Command struct
var RootCmd = &cobra.Command{
	Use:   "task",
	Short: "Task is a CLI task manager",
}
