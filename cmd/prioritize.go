package cmd

import (
	"strings"

	"github.com/gabey14/tasker-gocli/taskercli"
	"github.com/spf13/cobra"
)

func init() {

	var (
		setPriorityExample = `To prioritize a todo with id 35:
		tasker prioritize 35
    tasker p 35 `

		setPriorityDescLong = `Sets the priority of a todo with the given id. For more information, see https://tasker.io/docs/`

		setPriorityDesc = "Sets the priority of a todo"

		setUnPriorityExample = `To un-prioritize a todo with id 35:
		tasker un-prioritize 35
		tasker up 35 `

		setUnPriorityDescLong = `Unsets the priority of a todo with the given id. For more information, see https://tasker.io/docs/`

		setUnPriorityDesc = "Unsets the priority of a todo"
	)

	var setPriorityCmd = &cobra.Command{
		Use:     "prioritze [id]",
		Aliases: []string{"p"},
		Short:   setPriorityDesc,
		Long:    setPriorityDescLong,
		Example: setPriorityExample,
		Run: func(cmd *cobra.Command, args []string) {
			taskercli.NewApp().PrioritizeTodo(strings.Join(args, " "))
		},
	}

	var setUnPriorityCmd = &cobra.Command{
		Use:     "unprioritize [id]",
		Aliases: []string{"up"},
		Short:   setUnPriorityDesc,
		Long:    setUnPriorityDescLong,
		Example: setUnPriorityExample,
		Run: func(cmd *cobra.Command, args []string) {
			taskercli.NewApp().UnprioritizeTodo(strings.Join(args, " "))
		},
	}

	rootCmd.AddCommand(setPriorityCmd)
	rootCmd.AddCommand(setUnPriorityCmd)
}
