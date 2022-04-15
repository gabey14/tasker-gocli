package cmd

import (
	"github.com/gabey14/tasker-gocli/taskercli"
	"github.com/spf13/cobra"
)

func init() {

	var (
		initCmdDesc     = "Initialize a new todo list in the current directory"
		initCmdLongDesc = `Initializes a new todo list in the current directory. For more information, see https://sites.google.com/view/tasker-cli/manualfeatures`
	)

	var initCmd = &cobra.Command{
		Use:     "init",
		Aliases: []string{"i"},
		Short:   initCmdDesc,
		Long:    initCmdLongDesc,
		Run: func(cmd *cobra.Command, args []string) {
			taskercli.NewApp().InitializeRepo()
		},
	}

	rootCmd.AddCommand(initCmd)
}
