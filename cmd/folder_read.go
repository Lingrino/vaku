package cmd

import (
	"fmt"

	"github.com/Lingrino/vaku/vaku"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var folderReadCmd = &cobra.Command{
	Use:   "read [path]",
	Short: "Recursively read a vault folder",

	Args: cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		input := vaku.NewPathInput(args[0])
		input.TrimPathPrefix = !noTrimPathPrefix

		output, err := vgc.FolderRead(input)
		if err != nil {
			fmt.Printf("%s", errors.Wrapf(err, "Failed to read folder %s", args[0]))
		} else {
			print(map[string]interface{}{
				args[0]: output,
			})
		}
	},
}

func init() {
	folderCmd.AddCommand(folderReadCmd)
	folderReadCmd.PersistentFlags().BoolVarP(&noTrimPathPrefix, "no-trim-path-prefix", "T", true, "Output full paths instead of paths with the input path trimmed")
}
