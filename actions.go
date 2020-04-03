package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

const MessageLevelWarning = "warning"
const MessageLevelError = "error"

func Message(w io.Writer, level, file, line, column, message string) {
	if w == nil {
		w = os.Stdout
	}

	opts := []string{}

	if len(file) > 0 {
		opts = append(opts, "file=" + file)
	}

	if len(line) > 0 {
		opts = append(opts, "line=" + line)
	}

	if len(column) > 0 {
		opts = append(opts, "col=" + column)
	}

	optsStr := strings.Join(opts, ",")
	fmt.Fprintf(w, "::%s %s::%s\n", level, optsStr, message)
}

func Warning(w io.Writer, file, line, column, message string) {
	Message(w, MessageLevelWarning, file, line, column, message)
}

func Error(w io.Writer, file, line, column, message string) {
	Message(w, MessageLevelError, file, line, column, message)
}
