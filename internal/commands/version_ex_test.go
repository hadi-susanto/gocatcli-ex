/*
Fork-only file (main-ex branch of hadi-susanto/gocatcli-ex).
author: Hadi Susanto Lesmana (https://github.com/hadi-susanto)
*/

package commands

import "testing"

func TestVersionEx(t *testing.T) {
	var cases = []struct {
		commit string
		date   string
		want   string
	}{
		{"8a800ae", "2026.11.17", "1.3.4-2026.11.17 (8a800ae)"},
		{"8a800ae-dirty", "2026.11.17", "1.3.4-2026.11.17 (8a800ae-dirty)"},
		{"dev", "", "1.3.4 (dev)"},
		{"", "2026.11.17", "1.3.4-2026.11.17"},
		{"", "", "1.3.4"},
	}
	for _, c := range cases {
		var got = versionEx("1.3.4", c.commit, c.date)
		if got != c.want {
			t.Errorf("versionEx(%q, %q) = %q, want %q", c.commit, c.date, got, c.want)
		}
	}
}
