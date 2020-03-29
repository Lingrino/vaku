package vaku

import (
	"testing"

	"github.com/hashicorp/vault/api"
	"github.com/stretchr/testify/assert"
)

func TestPathList(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		give        string
		giveLogical logical
		giveOptions []Option
		want        []string
		wantErr     []error
		skipMount   bool
	}{
		{
			name:    "list test",
			give:    "test",
			want:    []string{"HToOeKKD", "fizz", "foo", "inner/", "value"},
			wantErr: nil,
		},
		{
			name:        "full path prefix",
			give:        "test/inner/again/",
			giveOptions: []Option{WithFullPath(true)},
			want:        []string{"test/inner/again/inner/"},
			wantErr:     nil,
		},
		{
			name:    "single secret",
			give:    "test/foo",
			want:    nil,
			wantErr: nil,
		},
		{
			name:    "list bad path",
			give:    "doesnotexist",
			want:    nil,
			wantErr: nil,
		},
		{
			name:    "no mount",
			give:    noMountPrefix,
			want:    nil,
			wantErr: nil,
		},
		{
			name: "list error",
			give: "test",
			giveLogical: &errLogical{
				err: errInject,
			},
			want:    nil,
			wantErr: []error{ErrPathList, ErrVaultList},
		},
		{
			name: "nil secret",
			give: "test",
			giveLogical: &errLogical{
				secret: nil,
			},
			want:    nil,
			wantErr: nil,
		},
		{
			name: "nil data",
			give: "test",
			giveLogical: &errLogical{
				secret: &api.Secret{
					Data: nil,
				},
			},
			want:    nil,
			wantErr: nil,
		},
		{
			name: "no keys",
			give: "test",
			giveLogical: &errLogical{
				secret: &api.Secret{
					Data: map[string]interface{}{
						"notkeys": "notkeys",
					},
				},
			},
			want:    nil,
			wantErr: []error{ErrPathList, ErrDecodeSecret},
		},
		{
			name: "keys not []interface{}",
			give: "test",
			giveLogical: &errLogical{
				secret: &api.Secret{
					Data: map[string]interface{}{
						"keys": 1,
					},
				},
			},
			want:    nil,
			wantErr: []error{ErrPathList, ErrDecodeSecret},
		},
		{
			name: "keys not string",
			give: "test",
			giveLogical: &errLogical{
				secret: &api.Secret{
					Data: map[string]interface{}{
						"keys": []interface{}{
							1,
						},
					},
				},
			},
			want:    nil,
			wantErr: []error{ErrPathList, ErrDecodeSecret},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ln, client := testClient(t, tt.giveOptions...)
			defer ln.Close()
			updateLogical(t, client, tt.giveLogical, tt.giveLogical)

			funcs := []func(string) ([]string, error){
				client.PathList,
				client.PathListDst,
			}

			for _, ver := range kvMountVersions {
				for _, f := range funcs {
					path := addMountToPath(t, tt.give, ver)

					list, err := f(path)
					TrimListPrefix(list, ver)

					compareErrors(t, err, tt.wantErr)
					assert.Equal(t, tt.want, list)
				}
			}
		})
	}
}
