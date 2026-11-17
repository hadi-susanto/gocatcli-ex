/*
Fork-only entry point (main-ex branch of hadi-susanto/gocatcli-ex).
Builds the gocatcli-ex binary, see Makefile.ex.
*/

package main

import (
	"os"
	"runtime/debug"
	"time"

	"github.com/deadc0de6/gocatcli/internal/commands"
	"github.com/deadc0de6/gocatcli/internal/log"
)

const binName = "gocatcli-ex"

// commit and commitDate are set at build time with
// -ldflags "-X main.commit=<short-sha> -X main.commitDate=<RFC3339 commit date>"
var commit = ""
var commitDate = ""

func main() {
	var vcsRev, vcsTime = vcsInfo()
	if commit == "" {
		commit = vcsRev
	}
	if commitDate == "" {
		commitDate = vcsTime
	}
	commands.ApplyEx(binName, commit, formatDate(commitDate))

	var err = commands.Execute()
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
}

// formatDate turns an RFC3339 date into "YYYY.MM.DD" (UTC, so every
// build path agrees on the day), or "" when it cannot be parsed.
func formatDate(date string) string {
	var t, err = time.Parse(time.RFC3339, date)
	if err != nil {
		return ""
	}
	return t.UTC().Format("2006.01.02")
}

// vcsInfo is the fallback when built without ldflags: it reads the
// git revision and commit time embedded by `go build` in a git checkout (Go >= 1.18).
func vcsInfo() (string, string) {
	var info, ok = debug.ReadBuildInfo()
	if !ok {
		return "dev", ""
	}
	var rev string
	var vcsTime string
	var dirty bool
	for _, s := range info.Settings {
		switch s.Key {
		case "vcs.revision":
			rev = s.Value
		case "vcs.time":
			vcsTime = s.Value
		case "vcs.modified":
			dirty = s.Value == "true"
		}
	}
	if rev == "" {
		return "dev", ""
	}
	if len(rev) > 7 {
		rev = rev[:7]
	}
	if dirty {
		rev += "-dirty"
	}
	return rev, vcsTime
}
