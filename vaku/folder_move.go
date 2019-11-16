package vaku

import (
	"fmt"
)

// folderMoveWorkerInput takes input/output channels for input to the job
type folderMoveWorkerInput struct {
	inputsC  <-chan map[string]*PathInput
	resultsC chan<- error
}

// FolderMove calls FolderCopy() with the same inputs followed by FolderDelete()
// on the source path if the copy was successful.
func (c *Client) FolderMove(s *PathInput, t *PathInput) error {
	var err error

	// Init both paths to get mount info
	s.opType = "readwrite"
	err = c.InitPathInput(s)
	if err != nil {
		return fmt.Errorf("failed to init path %s: %w", s.Path, err)
	}
	t.opType = "readwrite"
	err = c.InitPathInput(t)
	if err != nil {
		return fmt.Errorf("failed to init path %s: %w", t.Path, err)
	}

	// Get the keys to move
	list, err := c.FolderList(&PathInput{
		Path:           s.Path,
		TrimPathPrefix: true,
	})
	if err != nil {
		return fmt.Errorf("failed to list %s: %w", s.Path, err)
	}

	// Concurrency channels for workers
	inputsC := make(chan map[string]*PathInput, len(list))
	resultsC := make(chan error, len(list))

	// Spawn workers equal to MaxConcurrency
	for w := 1; w <= MaxConcurrency; w++ {
		go c.folderMoveWorker(&folderMoveWorkerInput{
			inputsC:  inputsC,
			resultsC: resultsC,
		})
	}

	// Add all paths to move to the inputs channel
	for _, p := range list {
		inputsC <- map[string]*PathInput{
			"source": {
				Path:          c.PathJoin(s.Path, p),
				mountPath:     s.mountPath,
				mountlessPath: s.mountlessPath,
				mountVersion:  s.mountVersion,
			},
			"target": {
				Path:          c.PathJoin(t.Path, p),
				mountPath:     t.mountPath,
				mountlessPath: t.mountlessPath,
				mountVersion:  t.mountVersion,
			},
		}
	}
	close(inputsC)

	// Empty the results channel into output
	for j := 0; j < len(list); j++ {
		o := <-resultsC
		if o != nil {
			err = fmt.Errorf("failed to move path: %w", o)
		}
	}

	return err
}

// folderMoveWorker does the work of copying a single path to a new destination
func (c *Client) folderMoveWorker(i *folderMoveWorkerInput) {
	var err error
	for {
		inputs, more := <-i.inputsC
		if more {
			err = c.PathMove(inputs["source"], inputs["target"])
			if err != nil {
				i.resultsC <- fmt.Errorf("failed to move path %s to %s: %w", inputs["source"].Path, inputs["target"].Path, err)
				continue
			}
			i.resultsC <- nil
		} else {
			return
		}
	}
}
