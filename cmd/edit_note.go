package cmd

import (
	"strconv"
	"strings"

	"github.com/gabey14/tasker-gocli/taskercli"
	"github.com/spf13/cobra"
)

func init() {

	var (
		editNoteCmdDesc = "Edits a note from a todo"
		editNoteCmdLong = editNoteCmdDesc + "\n For more information, see https://tasker.io/docs/"

		editNoteCmdExample = `To see todos with notes:
		tasker list --notes
		To edit note 0 from todo 1:
		tasker en 1 0 new note`
	)

	var editNoteCmd = &cobra.Command{
		Use:     "editNote <todoId> <noteId> <newNote>",
		Aliases: []string{"en"},
		Example: editNoteCmdExample,
		Short:   editNoteCmdDesc,
		Long:    editNoteCmdLong,
		Run: func(cmd *cobra.Command, args []string) {
			todoId, _ := strconv.Atoi(args[0])
			noteId, _ := strconv.Atoi(args[1])
			taskercli.NewApp().EditNote(todoId, noteId, strings.Join(args[2:], " "))
		},
	}

	rootCmd.AddCommand(editNoteCmd)

}
