package main

import (
	"encoding/xml"
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/pflag"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	noRestorePtr := pflag.Bool("no-restore", false, "")
	pflag.Parse()

	args := pflag.Args()
	if len(args) != 1 {
		usage()
	}

	path, err := filepath.Abs(args[0])
	if err != nil {
		panic(err)
	}

	ext := filepath.Ext(path)

	if ext == ".sln" {
		if !*noRestorePtr {
			dotnetRestore(path)
		}

		path = generateReport(path)
		ext = filepath.Ext(path)
	}

	if ext != ".xml" {
		usage()
	}

	report, err := getReport(path)
	if err != nil {
		panic(err)
	}

	printOutput(os.Stdout, *report)
}

func dotnetRestore(slnPath string) {
	cmd := exec.Command("dotnet", "restore", slnPath)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		os.Exit(1)
	}
}

func generateReport(slnPath string) string {
	outPath := filepath.Join(os.TempDir(), "report.xml")

	cmd := exec.Command("inspectcode.sh", "--output="+outPath, slnPath)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		os.Exit(1)
	}

	return outPath
}

func getReport(path string) (*Report, error) {
	path, err := filepath.Abs(path)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	report := Report{}
	err = xml.Unmarshal(bytes, &report)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &report, nil
}

func printOutput(w io.Writer, report Report) {
	typeMap := map[string]IssueType{}
	for _, issueType := range report.IssueTypes {
		typeMap[issueType.Id] = issueType
	}

	for _, project := range report.Issues {
		for _, issue := range project.Issues {
			issueType := typeMap[issue.TypeId]
			level := severityToLevel(issueType.Severity)
			column := strings.Split(issue.Offset, "-")[0]
			file := strings.ReplaceAll(issue.File, `\`, "/")
			Message(w, level, file, issue.Line, column, issue.Message)
		}
	}
}

func severityToLevel(severity string) string {
	switch severity {
	case "WARNING", "ERROR":
		return MessageLevelError
	default:
		return MessageLevelWarning
	}
}

func usage() {
	exe, _ := os.Executable()
	fmt.Fprintf(os.Stderr, `
Usage: %s [--no-restore] (solution.sln|results.xml)

There are two modes of operation:

* Pass a path to an .sln file: Runs code inspections and emits output in GitHub format.
* Pass a path to an .xml file: Converts existing inspection report to GitHub format.

You can optionally specify --no-restore if you want to skip running 'dotnet restore'.
`, exe)
	os.Exit(1)
}
