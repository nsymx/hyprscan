package checks

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/nsymx/hyprscan/internal/styles"

	"charm.land/lipgloss/v2"
)

func Configuration(hyprPath string) {

	configFiles := []string{
		filepath.Join(hyprPath, "hyprland.lua"),
		filepath.Join(hyprPath, "hyprlock.conf"),
		filepath.Join(hyprPath, "hypridle.conf"),
		filepath.Join(hyprPath, "mocha.conf"),
		filepath.Join(hyprPath, "conf/animations-high.lua"),
		filepath.Join(hyprPath, "conf/autostart.lua"),
		filepath.Join(hyprPath, "conf/environment.lua"),
		filepath.Join(hyprPath, "conf/monitor.lua"),
		filepath.Join(hyprPath, "conf/layouts.lua"),
		filepath.Join(hyprPath, "conf/window.lua"),
		filepath.Join(hyprPath, "conf/windowrules.lua"),
		filepath.Join(hyprPath, "conf/keyboard.lua"),
		filepath.Join(hyprPath, "conf/keybindings.lua"),
		filepath.Join(hyprPath, "conf/misc.lua"),
	}

	var missingConfigs []string

	for _, file := range configFiles {
		info, err := os.Stat(file)
		if os.IsNotExist(err) || (err == nil && info.IsDir()) {
			missingConfigs = append(missingConfigs, file)
		}
	}

	if len(missingConfigs) == 0 {
		lipgloss.Println(styles.ChecksStyle.Render("  ◦", "Config files:"), "OK ✓")
	} else {
		lipgloss.Println(styles.ErrorStyle.Render("The following config files are missing:\n"))

		for _, file := range missingConfigs {
			fmt.Println("->", file)
		}

		fmt.Println()
	}
}

func Bashrc(homeDir string) {

	bashrcPath := filepath.Join(homeDir, ".bashrc")
	if info, err := os.Stat(bashrcPath); err == nil && !info.IsDir() {
		lipgloss.Println(styles.ChecksStyle.Render("  ◦", "Bashrc:"), "Present ✓")
	} else {
		fmt.Println()
		lipgloss.Println(styles.ErrorStyle.Render("Cannot find .bashrc file...\n"))
	}
}
