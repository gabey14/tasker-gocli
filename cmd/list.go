package cmd

import (
	"strings"

	"github.com/gabey14/tasker-gocli/taskercli"
	"github.com/spf13/cobra"
)

func init() {

	var (
		unicodeSupport bool
		colorSupport   bool
		listNotes      bool
		showStatus     bool
		listCmdDesc    = "List todos"
		listCmdExample = `
		Filtering by date:
 		------------------

		tasker list due:<date>
  	tasker list duebefore:<date>
  	tasker list dueafter:<date>

		where <date> is one of:
		(tod|today|tom|tomorrow|thisweek|nextweek|lastweek|mon|tue|wed|thu|fri|sat|sun|none|<specific date>)

		List all todos due today:
			tasker list due:tod

		Lists all todos due tomorrow:
			tasker list due:tom

		Lists all todos due monday:
			tasker list due:mon

		Lists all todos with no due date:
			tasker list due:none

		Lists all overdue todos:
			tasker list duebefore:today

		Lists all todos in due in the future:
			tasker list dueafter:today

		When using a specific date, it needs to be in the format of jun23 or 23jun:
			tasker list due:jun23

		Filtering by status:
		--------------------

		List all todos with a status of "started"
			tasker list status:started

		List all todos without a status of "started"
			tasker list status:-started

		List all todos without a status of "started" or "finished"
			tasker list status:-started,-finished

		Filtering by projects or contexts:
  	----------------------------------

  	Project and context filtering are very similar:
			tasker list project:<project>
			tasker list context:<context>

  	List all todos with a project of "mobile"
    	tasker list project:mobile

  	List all todos with a project of "mobile" and "devops"
    	tasker list project:mobile,devops

  	List all todos with a project of "mobile" but not "devops"
    	tasker list project:mobile,-devops

  	List all todos without a project of "devops"
    	tasker list project:-devops

	 	Filtering by priority, completed, etc:
  	--------------------------------------

  	You can filter todos on their priority or completed status:
    	tasker list is:priority
    	tasker list not:priority

    	tasker list is:completed
    	tasker list not:completed

  	There are additional filters for showing completed todos:
    	tasker list completed:today
    	tasker list completed:thisweek

  	By default, tasker will not show archived todos. To show archived todos:
    	tasker list is:archived

		Grouping:
    ---------

  	List all todos grouped by context:
    	tasker list group:c

  	List all todos grouped by project:
    	tasker list group:p

  	List all todos grouped by status:
 	  	tasker list group:s

  	Combining filters:
 	  ------------------

  	Of course, you can combine grouping and filtering to get a nice formatted list.

  	Lists all todos due today grouped by context:
    	tasker list group:c due:today

  	Lists all todos due today for +mobile, grouped by context:
    	tasker list project:mobile group:c due:thisweek

  	Lists all prioritized todos that are not completed and are overdue.  Include a todo's notes when listing:
    	tasker list --notes is:priority not:completed duebefore:tod

  	Lists all todos due tomorrow concerning @frank for +project, grouped by project:
    	tasker list context:frank group:p due:tom

  	Indicator flags
  	---------------

  	If you pass --status=true as a flag, you'll see an extra column when listing todos.

  	* = Todo is prioritized
  	N = Todo has notes attached
  	A = Todo is archived
		`

		listCmdLongDesc = `List todos, optionally providing a filter.
		When listing todos, you can apply powerful filters, and perform grouping.
		
		See the full docs at https://tasker.io/docs/`
	)

	var listCmd = &cobra.Command{
		Use:     "list",
		Aliases: []string{"l", "ls"},
		Short:   listCmdDesc,
		Long:    listCmdLongDesc,
		Example: listCmdExample,
		Run: func(cmd *cobra.Command, args []string) {
			taskercli.NewAppWithPrintOptions(unicodeSupport, colorSupport).ListTodos(strings.Join(args, " "), listNotes, showStatus)
		},
	}

	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVarP(&unicodeSupport, "unicode", "", true, "Allows unicode support in Tasker output")

	listCmd.Flags().BoolVarP(&colorSupport, "color", "", true, "Allows color support in Tasker output")

	listCmd.Flags().BoolVarP(&listNotes, "notes", "", false, "Show notes when listing todos")

	listCmd.Flags().BoolVarP(&showStatus, "status", "", false, "Show a todo's status when listing todos")

}
