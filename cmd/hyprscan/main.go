package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"charm.land/lipgloss/v2"
	"charm.land/lipgloss/v2/tree"
	"github.com/charmbracelet/log"
	"github.com/nsymx/hyprscan/internal/helper"
	"github.com/nsymx/hyprscan/internal/styles"
)

var (
	name    = "hyprscan"
	version = "dev"
)

var logger = log.NewWithOptions(os.Stderr, log.Options{
	ReportTimestamp: true,
	Prefix:          ":",
})

func ClearScreen() {

	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("\033[2J\033[H")
		return
	}
}

func renderHeaderTree() {

	t := tree.Root(styles.HeaderTreeRootStyle.Render("•", "version")).
		Child(
			tree.New().
				Root(styles.HeaderTreeChildStyle.Render(version)),
		)

	lipgloss.Println(t)
}

func renderHeader() {

	lipgloss.Println(styles.HeaderStyle.Render("", "", name, ""))
	renderHeaderTree()

	fmt.Println()
	lipgloss.Println(styles.ChecksTitleStyle.Render("•", "Main Checks"))
}

func main() {

	if runtime.GOOS != "linux" {
		fmt.Println()
		logger.Fatal("Unsupported operating system, aborting...")
	}
	versionFlag := flag.Bool("version", false, "Print current version")
	flag.Parse()

	if *versionFlag {
		lipgloss.Println(styles.InfoStyle.Render(name), "-", strings.TrimSpace(version))
		os.Exit(0)
	}
	ClearScreen()

	renderHeader()
	helper.CoreRendering()
}
