package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "vaku",
	Short: "Vaku CLI extends the official Vault CLI with useful high-level functions",
	Long: `Vaku CLI extends the official Vault CLI with useful high-level functions

The Vaku CLI is intended to be used side by side with the official Vault CLI,
and only provides functions to extend the existing functionality.

Vaku does not log you in to vault or help you with getting a token. Like the CLI,
it will look for a token first at the VAULT_TOKEN env var and then in ~/.vault-token

Currently only supports json output.

Built by Sean Lingren <srlingren@gmail.com>
CLI documentation is available using 'vaku help [cmd]'
API documentation is available at https://godoc.org/github.com/Lingrino/vaku/vaku`,

	// Auth to vault on all commands
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		authVGC()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
