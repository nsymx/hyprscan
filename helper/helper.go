package helper

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/nsymx/hyprscan/checks"
	"github.com/nsymx/hyprscan/styles"

	"charm.land/lipgloss/v2"
)

func CoreRendering() {

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println()
		lipgloss.Println(styles.ErrorStyle.Render("Error:"), "Cannot define user home")
		os.Exit(1)
	}

	hyprPath := filepath.Join(homeDir, ".config", "hypr")
	sourcePath := filepath.Join(homeDir, "src", "Scripts")

	time.Sleep(700 * time.Millisecond)
	checks.Dependencies()

	time.Sleep(1 * time.Second)
	checks.Configuration(hyprPath)

	time.Sleep(1 * time.Second)
	checks.Bashrc(homeDir)

	time.Sleep(700 * time.Millisecond)
	checks.SourceScripts(sourcePath)

	time.Sleep(700 * time.Millisecond)
	checks.Monitors()
}
