package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gabey14/tasker-gocli/taskercli"
	"github.com/spf13/cobra"
)

func init() {

	var (
		editCmdDesc    = "Edits todos"
		editCmdExample = `To edit todo with id 35
		tasker e 35
		tasker edit 35 Change subject
		
		Edit due date:
		tasker e 35 due:tomorrow

		Remove due date:
		tasker e 35 due:none

		Edit status:
		tasker e 35 status:college

		Remove status:
		tasker e 35 status:none
		`
		editCmdLongDesc = `Edits todos with the given id. For more information, see https://tasker.io/docs/`
	)

	var editCmd = &cobra.Command{
		Use:     "edit [id] [key:value]",
		Aliases: []string{"e"},
		Short:   editCmdDesc,
		Long:    editCmdLongDesc,
		Example: editCmdExample,
		Run: func(cmd *cobra.Command, args []string) {
			todoID, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Printf("Could not parse todo ID: '%s'\n", args[0])
				return
			}
			taskercli.NewApp().EditTodo(todoID, strings.Join(args[1:], " "))
		},
	}

	rootCmd.AddCommand(editCmd)
}
