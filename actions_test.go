package main

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWarning(t *testing.T) {
	t.Run("all opts set", func(t *testing.T) {
		buf := &bytes.Buffer{}
		Warning(buf, "file", "line", "col", "msg")
		assert.Equal(t, "::warning file=file,line=line,col=col::msg", buf.String())
	})

	t.Run("one opt set", func(t *testing.T) {
		buf := &bytes.Buffer{}
		Warning(buf, "file", "", "", "msg")
		assert.Equal(t, "::warning file=file::msg", buf.String())
	})

	t.Run("no opts set", func(t *testing.T) {
		buf := &bytes.Buffer{}
		Warning(buf, "", "", "", "msg")
		assert.Equal(t, "::warning ::msg", buf.String())
	})
}

func TestError(t *testing.T) {
	t.Run("all opts set", func(t *testing.T) {
		buf := &bytes.Buffer{}
		Error(buf, "file", "line", "col", "msg")
		assert.Equal(t, "::error file=file,line=line,col=col::msg", buf.String())
	})

	t.Run("one opt set", func(t *testing.T) {
		buf := &bytes.Buffer{}
		Error(buf, "file", "", "", "msg")
		assert.Equal(t, "::error file=file::msg", buf.String())
	})

	t.Run("no opts set", func(t *testing.T) {
		buf := &bytes.Buffer{}
		Error(buf, "", "", "", "msg")
		assert.Equal(t, "::error ::msg", buf.String())
	})
}
