package checks

import (
	"os"

	"charm.land/lipgloss/v2"
	"github.com/nsymx/hyprscan/internal/styles"
)

func Monitors() {

	if os.Getenv("PRIMARY_MONITOR") != "" {
		lipgloss.Println(styles.ChecksStyle.Render("  ◦", "Primary monitor is set:"), "YES ✓")
	} else {
		lipgloss.Println(styles.ErrorStyle.Render("Primary monitor not set in environment...\n"))
	}

	if os.Getenv("SECONDARY_MONITOR") != "" {
		lipgloss.Println(styles.ChecksStyle.Render("  ◦", "Secondary monitor is set:"), "YES ✓")
	} else {
		lipgloss.Println(styles.ErrorStyle.Render("Secondary monitor not set in environment...\n"))
	}
}
