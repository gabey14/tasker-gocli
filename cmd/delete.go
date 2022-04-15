package cmd

import (
	"strings"

	"github.com/gabey14/tasker-gocli/taskercli"
	"github.com/spf13/cobra"
)

func init() {

	var (
		deleteCmdDesc    = "Deletes todos"
		deleteCmdExample = `To delete todo with id 35
		tasker d 35
		tasker delete 35
		
		Note, this will also free up the id 35`

		deleteCmdLongDesc = `Deletes todos with the given id. For more information, see https://sites.google.com/view/tasker-cli/manualfeatures`
	)

	var deleteCmd = &cobra.Command{
		Use:     "delete [id]",
		Aliases: []string{"d", "rm"},
		Short:   deleteCmdDesc,
		Long:    deleteCmdLongDesc,
		Example: deleteCmdExample,
		Run: func(cmd *cobra.Command, args []string) {
			taskercli.NewApp().DeleteTodo(strings.Join(args, " "))
		},
	}

	rootCmd.AddCommand(deleteCmd)
}
