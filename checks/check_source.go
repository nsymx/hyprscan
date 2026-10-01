package checks

import (
	"os"

	"charm.land/lipgloss/v2"
	"github.com/nsymx/hyprscan/styles"
)

func SourceScripts(sourcePath string) {

	if info, err := os.Stat(sourcePath); err == nil && info.IsDir() {
		lipgloss.Println(styles.ChecksStyle.Render("  ◦", "Source scripts:"), "Present ✓")
	} else {
		lipgloss.Println(styles.ErrorStyle.Render("Source scripts directory not found...\n"))
	}
}
