/*
Fork-only file (main-ex branch of hadi-susanto/gocatcli-ex).
Kept in its own file so upstream changes in root.go never conflict.
*/

package commands

// ApplyEx turns the shared commands into the gocatcli-ex flavour:
// --version shows the upstream x.y.z with the commit date and commit
// ("1.3.4-2026.11.17 (8a800ae)") and the command is named after the fork binary.
// Only the displayed version changes: catalogs keep the plain upstream version.
// It must be called before Execute.
func ApplyEx(binName string, commit string, date string) {
	rootCmd.Version = versionEx(version, commit, date)
	rootCmd.Use = binName
}

// versionEx builds "x.y.z-YYYY.MM.DD (commit)", leaving out the missing parts
func versionEx(base string, commit string, date string) string {
	var v = base
	if date != "" {
		v += "-" + date
	}
	if commit != "" {
		v += " (" + commit + ")"
	}
	return v
}
