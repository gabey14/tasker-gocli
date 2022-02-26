package cmd

import (
	"fmt"

	"github.com/gabey14/tasker-gocli/taskercli"
	"github.com/spf13/cobra"
)

func init() {
	var (
		versionCmdDesc     = "Displays the version of taskercli"
		versionCmdLongDesc = versionCmdDesc + "."
	)

	var versionCmd = &cobra.Command{
		Use:   "version",
		Long:  versionCmdLongDesc,
		Short: versionCmdDesc,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("TaskerCli v%s\n", taskercli.VERSION)
		},
	}

	rootCmd.AddCommand(versionCmd)

}
