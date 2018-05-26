package vaku_test

import (
	"testing"
	"vaku/vaku"

	"github.com/stretchr/testify/assert"
)

type TestPathWriteData struct {
	inputPath *vaku.PathInput
	inputData map[string]interface{}
	outputErr bool
}

func TestPathWrite(t *testing.T) {
	c := clientInitForTests(t)

	tests := map[int]TestPathWriteData{
		1: {
			inputPath: vaku.NewPathInput("secretv1/writetest/foo"),
			inputData: map[string]interface{}{
				"value": "bar",
			},
			outputErr: false,
		},
		2: {
			inputPath: vaku.NewPathInput("secretv2/writetest/foo"),
			inputData: map[string]interface{}{
				"value": "bar",
			},
			outputErr: false,
		},
		3: {
			inputPath: vaku.NewPathInput("secretv1/writetest/bar/"),
			inputData: map[string]interface{}{
				"Eg5ljS7t": "6F1B5nBg",
				"quqr32S5": "81iY4HAN",
				"r6R0JUzX": "rs1mCRB5",
			},
			outputErr: false,
		},
		4: {
			inputPath: vaku.NewPathInput("secretv2/writetest/bar/"),
			inputData: map[string]interface{}{
				"Eg5ljS7t": "6F1B5nBg",
				"quqr32S5": "81iY4HAN",
				"r6R0JUzX": "rs1mCRB5",
			},
			outputErr: false,
		},
		5: {
			inputPath: vaku.NewPathInput("secretdoesnotexist/writetest/bar"),
			inputData: map[string]interface{}{
				"Eg5ljS7t": "6F1B5nBg",
				"quqr32S5": "81iY4HAN",
				"r6R0JUzX": "rs1mCRB5",
			},
			outputErr: true,
		},
	}

	for _, d := range tests {
		e := c.PathWrite(d.inputPath, d.inputData)
		readBack, re := c.PathRead(d.inputPath)
		if d.outputErr {
			assert.Error(t, e)
		} else {
			assert.Equal(t, readBack, d.inputData)
			assert.NoError(t, e)
			assert.NoError(t, re)
		}
	}
}
