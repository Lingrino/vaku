package vaku

import (
	"errors"
)

var (
	// ErrPathMove when PathMove fails.
	ErrPathMove = errors.New("path move")
)

// PathMove moves data at a source path to a destination path (copy + delete).
func (c *Client) PathMove(src, dst string) error {
	err := c.PathCopy(src, dst)
	if err != nil {
		return newWrapErr("", ErrPathMove, err)
	}

	err = c.dc.PathDelete(src)
	if err != nil {
		return newWrapErr(dst, ErrPathMove, err)
	}

	return nil
}
