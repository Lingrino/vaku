package vaku

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPathWrite(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		give           string
		giveData       map[string]interface{}
		giveLogical    logical
		giveOptions    []Option
		wantErr        []error
		wantNoReadback bool
	}{
		{
			name: "new path",
			give: "write/bar",
			giveData: map[string]interface{}{
				"Eg5ljS7t": "6F1B5nBg",
				"quqr32S5": "81iY4HAN",
				"r6R0JUzX": "rs1mCRB5",
			},
			wantErr: nil,
		},
		{
			name: "overwrite",
			give: "test/foo",
			giveData: map[string]interface{}{
				"foo": "bar",
			},
			wantErr: nil,
		},
		{
			name:     "nil data",
			give:     "write/foo",
			giveData: nil,
			wantErr:  []error{ErrPathWrite, ErrNilData},
		},
		{
			name: "no mount",
			give: noMountPrefix,
			giveData: map[string]interface{}{
				"foo": "bar",
			},
			wantErr:        []error{ErrPathWrite, ErrVaultWrite},
			wantNoReadback: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			client := testClient(t, tt.giveOptions...)
			readbackClient := cloneCLient(t, client)
			updateLogical(t, client, tt.giveLogical, tt.giveLogical)

			funcs := []func(string, map[string]interface{}) error{
				client.PathWrite,
				client.PathWriteDst,
			}

			for _, ver := range kvMountVersions {
				for _, f := range funcs {
					path := addMountToPath(t, tt.give, ver)

					err := f(path, tt.giveData)
					compareErrors(t, err, tt.wantErr)

					if !tt.wantNoReadback {
						readBack, err := readbackClient.PathRead(path)
						assert.NoError(t, err)
						assert.Equal(t, tt.giveData, readBack)
					}
				}
			}
		})
	}
}
