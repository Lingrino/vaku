package vaku

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsFolder(t *testing.T) {
	t.Parallel()

	tests := []struct {
		give string
		want bool
	}{
		{
			give: "/",
			want: true,
		},
		{
			give: "a/",
			want: true,
		},
		{
			give: "a/b/",
			want: true,
		},
		{
			give: "",
			want: false,
		},
		{
			give: "a",
			want: false,
		},
		{
			give: "a/b",
			want: false,
		},
		{
			give: "123/456",
			want: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.give, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, IsFolder(tt.give))
		})
	}
}

func TestMakeFolder(t *testing.T) {
	t.Parallel()

	tests := []struct {
		give string
		want string
	}{
		{
			give: "",
			want: "/",
		},
		{
			give: "a",
			want: "a/",
		},
		{
			give: "a/",
			want: "a/",
		},
		{
			give: "a/b",
			want: "a/b/",
		},
		{
			give: "a/",
			want: "a/",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.give, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, MakeFolder(tt.give))
		})
	}
}

func TestEnsurePrefix(t *testing.T) {
	t.Parallel()

	tests := []struct {
		give       string
		givePrefix string
		want       string
	}{
		{
			give:       "",
			givePrefix: "",
			want:       "",
		},
		{
			give:       "a",
			givePrefix: "",
			want:       "a",
		},
		{
			give:       "",
			givePrefix: "a",
			want:       "a",
		},
		{
			give:       "a/",
			givePrefix: "a",
			want:       "a/",
		},
		{
			give:       "a",
			givePrefix: "a/",
			want:       "a/a",
		},
		{
			give:       "a/b/c/d",
			givePrefix: "a/b/",
			want:       "a/b/c/d",
		},
		{
			give:       "a/b/c/d",
			givePrefix: "b",
			want:       "b/a/b/c/d",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.give, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, EnsurePrefix(tt.give, tt.givePrefix))
		})
	}
}

func TestKeyJoin(t *testing.T) {
	t.Parallel()

	tests := []struct {
		give []string
		want string
	}{
		{
			give: []string{"/"},
			want: "/",
		},
		{
			give: []string{"a/"},
			want: "a/",
		},
		{
			give: []string{"b", ""},
			want: "b",
		},
		{
			give: []string{"a/b", "c"},
			want: "a/b/c",
		},
		{
			give: []string{"d/e/", "/f"},
			want: "d/e/f",
		},
		{
			give: []string{"/g/h/", "/i/"},
			want: "g/h/i/",
		},
		{
			give: []string{"/j/", "/k/l", "m"},
			want: "j/k/l/m",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.want, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, KeyJoin(tt.give...))
		})
	}
}

func TestPathJoin(t *testing.T) {
	t.Parallel()

	tests := []struct {
		give []string
		want string
	}{
		{
			give: []string{"/"},
			want: "",
		},
		{
			give: []string{"a/"},
			want: "a",
		},
		{
			give: []string{"b", ""},
			want: "b",
		},
		{
			give: []string{"a/b", "c"},
			want: "a/b/c",
		},
		{
			give: []string{"d/e/", "/f"},
			want: "d/e/f",
		},
		{
			give: []string{"/g/h/", "/i/"},
			want: "g/h/i",
		},
		{
			give: []string{"/j/", "/k/l", "m"},
			want: "j/k/l/m",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.want, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, PathJoin(tt.give...))
		})
	}
}

func TestPrefixList(t *testing.T) {
	t.Parallel()

	tests := []struct {
		giveList   []string
		givePrefix string
		want       []string
	}{
		{
			giveList:   []string{"a"},
			givePrefix: "b",
			want:       []string{"b/a"},
		},
		{
			giveList:   []string{"/c/d/e/"},
			givePrefix: "/f/",
			want:       []string{"f/c/d/e/"},
		},
		{
			giveList:   []string{"/g/"},
			givePrefix: "h",
			want:       []string{"h/g/"},
		},
		{
			giveList:   []string{"i/j"},
			givePrefix: "i",
			want:       []string{"i/i/j"},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.givePrefix, func(t *testing.T) {
			t.Parallel()

			PrefixList(tt.giveList, tt.givePrefix)

			assert.Equal(t, tt.want, tt.giveList)
		})
	}
}

func TestTrimListPrefix(t *testing.T) {
	t.Parallel()

	tests := []struct {
		giveList   []string
		givePrefix string
		want       []string
	}{
		{
			giveList:   []string{"a"},
			givePrefix: "b",
			want:       []string{"a"},
		},
		{
			giveList:   []string{"/c/d/e/"},
			givePrefix: "/c/",
			want:       []string{"d/e/"},
		},
		{
			giveList:   []string{"f/g"},
			givePrefix: "f",
			want:       []string{"g"},
		},
		{
			giveList:   []string{"i/j"},
			givePrefix: "k",
			want:       []string{"i/j"},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.givePrefix, func(t *testing.T) {
			t.Parallel()

			TrimListPrefix(tt.giveList, tt.givePrefix)

			assert.Equal(t, tt.want, tt.giveList)
		})
	}
}

func TestTrimMapKeyPrefix(t *testing.T) {
	t.Parallel()

	tests := []struct {
		giveMap    map[string]map[string]interface{}
		givePrefix string
		want       map[string]map[string]interface{}
	}{
		{
			giveMap: map[string]map[string]interface{}{
				"foo/bar": {"a": "b"},
			},
			givePrefix: "foo",
			want: map[string]map[string]interface{}{
				"bar": {"a": "b"},
			},
		},
		{
			giveMap: map[string]map[string]interface{}{
				"foo/bar": {"a": "b"},
			},
			givePrefix: "foo/",
			want: map[string]map[string]interface{}{
				"bar": {"a": "b"},
			},
		},
		{
			giveMap: map[string]map[string]interface{}{
				"foo/bar": {"a": "b"},
			},
			givePrefix: "fo",
			want: map[string]map[string]interface{}{
				"o/bar": {"a": "b"},
			},
		},
		{
			giveMap: map[string]map[string]interface{}{
				"foo/bar": {"a": "b"},
			},
			givePrefix: "fooo",
			want: map[string]map[string]interface{}{
				"foo/bar": {"a": "b"},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.givePrefix, func(t *testing.T) {
			t.Parallel()

			TrimMapKeyPrefix(tt.giveMap, tt.givePrefix)

			assert.Equal(t, tt.want, tt.giveMap)
		})
	}
}

func TestEnsureMapKeyPrefix(t *testing.T) {
	t.Parallel()

	tests := []struct {
		giveMap    map[string]map[string]interface{}
		givePrefix string
		want       map[string]map[string]interface{}
	}{
		{
			giveMap: map[string]map[string]interface{}{
				"foo/bar": {"a": "b"},
			},
			givePrefix: "foo",
			want: map[string]map[string]interface{}{
				"foo/bar": {"a": "b"},
			},
		},
		{
			giveMap: map[string]map[string]interface{}{
				"foo/bar": {"a": "b"},
			},
			givePrefix: "foo/",
			want: map[string]map[string]interface{}{
				"foo/bar": {"a": "b"},
			},
		},
		{
			giveMap: map[string]map[string]interface{}{
				"foo/bar": {"a": "b"},
			},
			givePrefix: "fo",
			want: map[string]map[string]interface{}{
				"foo/bar": {"a": "b"},
			},
		},
		{
			giveMap: map[string]map[string]interface{}{
				"foo/bar": {"a": "b"},
			},
			givePrefix: "fooo",
			want: map[string]map[string]interface{}{
				"fooo/foo/bar": {"a": "b"},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.givePrefix, func(t *testing.T) {
			t.Parallel()

			EnsureMapKeyPrefix(tt.giveMap, tt.givePrefix)

			assert.Equal(t, tt.want, tt.giveMap)
		})
	}
}
