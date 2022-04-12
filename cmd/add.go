package cmd

import (
	"strings"

	"github.com/gabey14/tasker-gocli/taskercli"
	"github.com/spf13/cobra"
)

func init() {
	// TASK - Add documentation link once done
	var (
		addCmdDesc    = "Adds todos"
		addCmdExample = `  tasker add Prepare meeting notes about +importantProject for the meeting with @john due:today
  tasker add Meeting with @john about +project due:tod
  tasker a +work +verify did @john fix the build? due:tom
  tasker a here is an important task priority:true recur:weekdays due:tom`

		addCmdLongDesc = `Adds todos.

  You can optionally specify a due date.
  This can be done by by putting 'due:<date>' at the end, where <date> is in (tod|today|tom|tomorrow|mon|tue|wed|thu|fri|sat|sun|thisweek|nextweek).

  Dates can also be explicit, using 3 characters for the month.  They can be written in 2 different formats:
    tasker a buy flowers for mom due:may12
    tasker get halloween candy due:31oct

  Todos can also recur.  Set the 'recur' directive to control recurrence:
    tasker a Daily standup recur:weekdays
    tasker a 1o1 meeting with jim recur:weekly

  For the full documentation on recurrence, see the docs:
  https://taskercli.io/docs/`
	)

	var addCmd = &cobra.Command{
		Use:     "add <todo>",
		Aliases: []string{"a"},
		Example: addCmdExample,
		Short:   addCmdDesc,
		Long:    addCmdLongDesc,
		Run: func(cmd *cobra.Command, args []string) {
			taskercli.NewApp().AddTodo(strings.Join(args, " "))
		},
	}

	rootCmd.AddCommand(addCmd)
}
