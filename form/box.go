package form

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"
)

func PrintFullWidthBox(title string) {
	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		width = 80 // fallback
	}
	if width < 10 {
		width = 10
	}

	topBottom := "┌" + strings.Repeat("─", width-2) + "┐"
	emptyLine := "│" + strings.Repeat(" ", width-2) + "│"

	fmt.Println(topBottom)
	if title != "" {
		// Center the title
		titleLen := len(title)
		if titleLen < width-2 {
			left := (width - 2 - titleLen) / 2
			right := width - 2 - titleLen - left
			fmt.Printf("│%s%s%s│\n", strings.Repeat(" ", left), title, strings.Repeat(" ", right))
		} else {
			// Truncate if too long
			fmt.Printf("│%s│\n", title[:width-4]+"..")
		}
	} else {
		fmt.Println(emptyLine)
	}
	fmt.Println("└" + strings.Repeat("─", width-2) + "┘")
}
