package cmd

import (
	"strings"

	"github.com/gabey14/tasker-gocli/taskercli"
	"github.com/spf13/cobra"
)

func init() {

	var (
		archiveCompletedTodo bool
		CompleteTodosCmdDesc = "completes todos"

		CompleteTodosCmdLongDesc = `completes todos. For more information, see https://tasker.io/docs/`

		CompleteTodosCmdExample = `To complete todo with id 35
		tasker complete 35
		tasker c 35

		Complete a todo with id 35 and archive it
		tasker complete 35 --archive
		
		Uncomplete a todo with id 35
		tasker uncomplete 35
		tasker uc 35 `

		UncompleteToDoCmdDesc      = "uncompletes todos"
		UncompleteTodosCmdLongDesc = `uncompletes todos. For more information, see https://tasker.io/docs/`
		UncompleteTodosCmdExample  = `To uncomplete todo with id 35
		tasker uncomplete 35
		tasker uc 35 `
	)

	var completeCmd = &cobra.Command{
		Use:     "complete [id]",
		Aliases: []string{"c"},
		Short:   CompleteTodosCmdDesc,
		Long:    CompleteTodosCmdLongDesc,
		Example: CompleteTodosCmdExample,
		Run: func(cmd *cobra.Command, args []string) {
			taskercli.NewApp().CompleteTodo(strings.Join(args, " "), archiveCompletedTodo)
		},
	}

	var uncompleteCmd = &cobra.Command{
		Use:     "uncomplete [id]",
		Aliases: []string{"uc"},
		Short:   UncompleteToDoCmdDesc,
		Long:    UncompleteTodosCmdLongDesc,
		Example: UncompleteTodosCmdExample,
		Run: func(cmd *cobra.Command, args []string) {
			taskercli.NewApp().UncompleteTodo(strings.Join(args, " "))
		},
	}

	rootCmd.AddCommand(completeCmd)
	completeCmd.Flags().BoolVarP(&archiveCompletedTodo, "archive", "", false, "Archive completed todo automatically")
	rootCmd.AddCommand(uncompleteCmd)

}
