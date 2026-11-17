/*
Fork-only file (main-ex branch of hadi-susanto/gocatcli-ex).
author: Hadi Susanto Lesmana (https://github.com/hadi-susanto)
*/

package main

import "testing"

func TestFormatDate(t *testing.T) {
	var cases = map[string]string{
		"2026-11-17T03:04:05Z":      "2026.11.17", // go build vcs.time, goreleaser .CommitDate
		"2026-11-17T08:30:00+07:00": "2026.11.17", // git %cI with a local offset
		"2026-11-17T02:00:00+07:00": "2026.11.16", // previous day in UTC
		"":                          "",
		"not-a-date":                "",
	}
	for in, want := range cases {
		var got = formatDate(in)
		if got != want {
			t.Errorf("formatDate(%q) = %q, want %q", in, got, want)
		}
	}
}
