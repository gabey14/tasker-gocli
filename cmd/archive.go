package cmd

import (
	"strings"

	"github.com/gabey14/tasker-gocli/taskercli"
	"github.com/spf13/cobra"
)

func init() {
	var (
		archiveExample = `
	To arvhive a todo with id 35:
    tasker archive 35
    tasker ar 35

	Garbage collection will delete all archived todos, reclaming ids:
    tasker archive gc
    tasker ar gc

	See the full docs here:
  https://tasker.io/docs/ `

		unarchiveExample = `
		To unarchive todo with id 35:
    tasker unarchive 35
    tasker uar 35`

		archiveCompletedExample = `
		To archive all completed todos:
    tasker archive completed
    tasker ar c
		`

		archiveGarbageCollectionExample = `
		Garbage collection will delete all archived todos, reclaming ids:
    tasker archive gc
    tasker ar gc`
	)

	var archivedCmd = &cobra.Command{
		Use:     `archive [id]`,
		Aliases: []string{"ar"},
		Example: archiveExample,
		Short:   "Archive a todo",
		Run: func(cmd *cobra.Command, args []string) {
			taskercli.NewApp().ArchiveTodo(strings.Join(args, " "))
		},
	}

	var unarchiveCmd = &cobra.Command{
		Use:     `unarchive [id]`,
		Aliases: []string{"uar"},
		Example: unarchiveExample,
		Short:   "Un-archives a todo",
		Run: func(cmd *cobra.Command, args []string) {
			taskercli.NewApp().UnarchiveTodo(strings.Join(args, " "))
		},
	}

	var archiveCompletedCmd = &cobra.Command{
		Use:     "c",
		Example: archiveCompletedExample,
		Short:   "Archives all completed todos",
		// TASK - Add docs link
		Long: "For more info, check out https://tasker.io/docs/",
		Run: func(cmd *cobra.Command, args []string) {
			taskercli.NewApp().ArchiveCompleted()
		},
	}

	var archiveGarbageCollectCmd = &cobra.Command{
		Use:     `gc`,
		Example: archiveGarbageCollectionExample,
		Short:   "Delete all archived todos",
		// TASK - Add docs link
		Long: "For more info, check out https://tasker.io/docs/",
		Run: func(cmd *cobra.Command, arsg []string) {
			taskercli.NewApp().GarbageCollect()
		},
	}

	rootCmd.AddCommand(archivedCmd)
	rootCmd.AddCommand(unarchiveCmd)
	rootCmd.AddCommand(archiveCompletedCmd)
	rootCmd.AddCommand(archiveGarbageCollectCmd)

}
