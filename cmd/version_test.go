package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVersion(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		giveVersion string
		giveArgs    []string
		wantOut     string
		wantErr     string
	}{
		{
			name:        "test",
			giveVersion: "test",
			wantOut:     "CLI: test\nAPI: 2.0.0\n",
		},
		{
			name:        "version",
			giveVersion: "version",
			wantOut:     "CLI: version\nAPI: 2.0.0\n",
		},
		{
			name:        "args",
			giveVersion: "version",
			giveArgs:    []string{"arg1", "arg2"},
			wantOut:     "",
			wantErr:     "unknown command",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			args := append([]string{"version"}, tt.giveArgs...)
			cli, outW, errW := newTestCLI(t, args)
			cli.setVersion(tt.giveVersion)
			assert.Equal(t, "", errW.String())

			err := cli.cmd.Execute()

			assertError(t, err, tt.wantErr)
			if tt.wantErr == "" {
				assert.Equal(t, tt.wantOut, outW.String())
			}
		})
	}
}
