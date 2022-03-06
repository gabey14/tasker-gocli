package cmd

import (
	"strings"

	"github.com/gabey14/tasker-gocli/taskercli"
	"github.com/spf13/cobra"
)

func init() {

	var (
		setStatusCmdDesc    = "Sets the status of a todo"
		setStatusCmdExample = `To add a "college" status to a todo:
    tasker status 33 college
    tasker s 33 college

  You can remove a status by setting a status to "none".  Example:
    tasker s 33 none `

		setStatusCmdLongDesc = `Sets the status of a todo with the given id. For more information, see https://tasker.io/docs/`
	)

	var setStatusCmd = &cobra.Command{
		Use:     "status [id] <status>",
		Aliases: []string{"s"},
		Short:   setStatusCmdDesc,
		Long:    setStatusCmdLongDesc,
		Example: setStatusCmdExample,
		Run: func(cmd *cobra.Command, args []string) {
			taskercli.NewApp().SetTodoStatus(strings.Join(args, " "))
		},
	}

	rootCmd.AddCommand(setStatusCmd)

}
