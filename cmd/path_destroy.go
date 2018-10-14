package cmd

import (
	"fmt"

	"github.com/Lingrino/vaku/vaku"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var pathDestroyCmd = &cobra.Command{
	Use:   "destroy [path]",
	Short: "Destroy a vault path (V2 mounts only)",

	Args: cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		input := vaku.NewPathInput(args[0])

		err := vgc.PathDestroy(input)
		if err != nil {
			fmt.Printf("%s", errors.Wrapf(err, "Failed to destroy path %s", args[0]))
		} else {
			print(map[string]interface{}{
				args[0]: "Successfully destroyed path, if it existed",
			})
		}
	},
}

func init() {
	pathCmd.AddCommand(pathDestroyCmd)
}
