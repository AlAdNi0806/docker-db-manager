package form

import (
	"fmt"
	"os"
)

func colorize(text string, colorCode int) string {
	if colorCode >= 0 && colorCode <= 255 {
		return fmt.Sprintf("\033[38;5;%dm%s\033[0m", colorCode, text)
	}
	// Fallback to 16-color mode if out of range
	return fmt.Sprintf("\033[%dm%s\033[0m", colorCode, text)
}

func backgroundColorize(text string, colorCode int) string {
	if colorCode >= 0 && colorCode <= 255 {
		return fmt.Sprintf("\033[48;5;%dm%s\033[0m", colorCode, text)
	}
	return fmt.Sprintf("\033[%dm%s\033[0m", colorCode, text)
}

func formatBold(text string) string {
	return fmt.Sprintf("\033[1m%s\033[0m", text)
}

func formatTitle(text string) string {
	return formatBold(colorize(text, 5))
}

func formatDescription(text string) string {
	return colorize(text, 244)
}

func clearLastLines(n int) {
	fmt.Printf("\033[%dA", n)
	fmt.Print("\033[J")
}

func addLinePrefix(str string) string {
	return fmt.Sprintf("%s %s", colorize("\u2502", 240), str)
}

func renderSelectInstructions() {
	fmt.Println()
	line := fmt.Sprintf(" %s %s %s %s %s %s",
		colorize("↑", 244),
		colorize("up ⋅", 242),
		colorize("↓", 244),
		colorize("down ⋅", 242),
		colorize("enter", 244),
		colorize("select", 242),
	)
	fmt.Println(line)
}

func renderSwitchInstructions() {
	fmt.Println()
	line := fmt.Sprintf(" %s %s %s %s",
		colorize("←/→", 244),
		colorize("toggle ⋅", 242),
		colorize("enter", 244),
		colorize("next", 242),
	)
	fmt.Println(line)
}

func renderInputInstructions() {
	fmt.Println()
	line := fmt.Sprintf(" %s %s",
		colorize("enter", 244),
		colorize("next", 242),
	)
	fmt.Println(line)
}

func Cleanup() {
	fmt.Print("\033[?25h")
	fmt.Println("\nThe Process Was Interupted!")
	os.Exit(0)
}
