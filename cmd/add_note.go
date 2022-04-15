package cmd

import (
	"strconv"
	"strings"

	"github.com/gabey14/tasker-gocli/taskercli"
	"github.com/spf13/cobra"
)

func init() {
	var (
		addNoteCmdDesc     = "Add a note to a task"
		addNoteCmdLongDesc = addNoteCmdDesc + "\n For more information, see https://sites.google.com/view/tasker-cli/manualfeatures"
		addNoteCmdExample  = " tasker an 35 note for todo with id 35"
	)

	var addNoteCmd = &cobra.Command{
		Use:     "addnote <todoId> <note>",
		Aliases: []string{"an"},
		Example: addNoteCmdExample,
		Short:   addNoteCmdDesc,
		Long:    addNoteCmdLongDesc,
		Args:    cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			todoId, _ := strconv.Atoi(args[0])
			taskercli.NewApp().AddNote(todoId, strings.Join(args[1:], " "))
		},
	}

	rootCmd.AddCommand(addNoteCmd)

}
