package cmd

import (
	"strconv"

	"github.com/gabey14/tasker-gocli/taskercli"
	"github.com/spf13/cobra"
)

func init() {

	var (
		deleteNoteCmdDesc     = "Deletes a note from a todo"
		deleteNoteCmdLongDesc = deleteNoteCmdDesc + "\n For more information, see https://sites.google.com/view/tasker-cli/manualfeatures"
		deleteNoteCmdExample  = `To see todos with notes:
		tasker list --notes
		To delete note 0 from todo 1:
		tasker dn 1 0`
	)

	var deleteNoteCmd = &cobra.Command{
		Use:     "deleteNote <todoId> <noteId>",
		Aliases: []string{"dn"},
		Example: deleteNoteCmdExample,
		Short:   deleteNoteCmdDesc,
		Long:    deleteNoteCmdLongDesc,
		Run: func(cmd *cobra.Command, args []string) {
			todoId, _ := strconv.Atoi(args[0])
			noteId, _ := strconv.Atoi(args[1])
			taskercli.NewApp().DeleteNote(todoId, noteId)
		},
	}

	rootCmd.AddCommand(deleteNoteCmd)
}
