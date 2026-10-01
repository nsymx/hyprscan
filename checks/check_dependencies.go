package checks

import (
	"fmt"
	"os/exec"

	"charm.land/lipgloss/v2"
	"github.com/nsymx/hyprscan/styles"
)

func Dependencies() {

	var packages = []string{
		"hypridle",
		"hyprland",
		"hyprlock",
		"uwsm",
		"lua",
	}

	var missingPackages []string

	for _, pkg := range packages {
		_, err := exec.LookPath(pkg)

		if err != nil {
			missingPackages = append(missingPackages, pkg)
		}
	}

	if len(missingPackages) == 0 {
		lipgloss.Println(styles.ChecksStyle.Render("  ◦", "Dependencies:"), "Installed ✓")
	} else {
		fmt.Println()
		lipgloss.Println(styles.ErrorStyle.Render("The following packages are missing:\n"))

		for _, pkg := range missingPackages {
			fmt.Println("->", pkg)
		}

		fmt.Println()
	}
}
