package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"unicode"

	"github.com/eiannone/keyboard"
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

func cleanup() {
	keyboard.Close()
	fmt.Print("\033[?25h")
	fmt.Println("\nThe Process Was Interupted!")
	os.Exit(0)
}

func NewSelect(question string, description string, options []string) string {
	selected := 0

	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer keyboard.Close()

	for {
		fmt.Println("")
		fmt.Println(addLinePrefix(formatTitle(question)))

		for i, option := range options {
			cursor := "  "
			renderOption := option
			if i == selected {
				cursor = colorize("> ", 6)
				renderOption = colorize(option, 107)
			}
			line := fmt.Sprintf("%s%s", cursor, renderOption)
			fmt.Println(addLinePrefix(line))
		}

		renderSelectInstructions()

		_, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}

		linesToClear := 4 + len(options)

		if key == keyboard.KeyArrowUp && selected > 0 {
			selected--
			clearLastLines(linesToClear)
		} else if key == keyboard.KeyArrowDown && selected < len(options)-1 {
			selected++
			clearLastLines(linesToClear)
		} else if key == keyboard.KeyEnter {
			clearLastLines(linesToClear)
			return options[selected]
		} else if key == keyboard.KeyCtrlC {
			clearLastLines(linesToClear)
			cleanup()
		} else {
			clearLastLines(linesToClear)
		}
	}
}

func NewSwitch(question string, description string, options [2]string, defaultValue bool) bool {
	selected := defaultValue

	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer keyboard.Close()

	for {
		fmt.Println("")
		fmt.Println(addLinePrefix(formatTitle(question)))
		fmt.Println(addLinePrefix(""))

		bg1 := 107 // selected background
		bg2 := 238 // unselected background

		var leftBg, rightBg int
		if selected {
			leftBg = bg1
			rightBg = bg2
		} else {
			leftBg = bg2
			rightBg = bg1
		}

		switchLine := fmt.Sprintf("%s %s",
			backgroundColorize(fmt.Sprintf("  %s  ", options[0]), leftBg),
			backgroundColorize(fmt.Sprintf("  %s  ", options[1]), rightBg),
		)

		fmt.Println(addLinePrefix(switchLine))

		renderSwitchInstructions()

		_, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}

		linesToClear := 4 + len(options) // 2 for the instructions + 2 for the title plus line before + the switch line and the line before

		if key == keyboard.KeyArrowLeft && selected == false {
			selected = true
			clearLastLines(linesToClear)
		} else if key == keyboard.KeyArrowRight && selected == true {
			selected = false
			clearLastLines(linesToClear)
		} else if key == keyboard.KeyEnter {
			clearLastLines(linesToClear)
			return selected
		} else if key == keyboard.KeyCtrlC {
			clearLastLines(linesToClear)
			cleanup()
		} else {
			// we can have a bool that will indicate if the state changed if not we skip the rendering and go to the waiting again
			clearLastLines(linesToClear)
		}
	}
}

func NewInput(question string, description string, placeholder string) string {
	value := ""
	cursorPosition := 0

	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer keyboard.Close()

	for {
		fmt.Println("")
		fmt.Println(addLinePrefix(formatTitle(question)))

		displayValue := value
		if displayValue == "" {
			if placeholder != "" {
				displayValue = fmt.Sprintf("%s%s",
					backgroundColorize(colorize(string(placeholder[0]), 235), 230),
					colorize(placeholder[1:], 242),
				)
			} else {
				displayValue = fmt.Sprintf("%s",
					backgroundColorize(" ", 230),
				)
			}
		} else {
			partBeforeCursor := displayValue[:cursorPosition]
			var partAtCursor string
			var partAfterCursor string
			if len(value) == cursorPosition {
				partAtCursor = " "
				partAfterCursor = ""
			} else {
				partAtCursor = string(displayValue[cursorPosition])
				partAfterCursor = displayValue[cursorPosition+1:]
			}

			displayValue = fmt.Sprintf("%s%s%s",
				partBeforeCursor,
				backgroundColorize(colorize(partAtCursor, 235), 230),
				partAfterCursor,
			)
		}
		fmt.Println(addLinePrefix(fmt.Sprintf("%s", displayValue)))

		renderInputInstructions()

		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}

		linesToClear := 5

		if key == keyboard.KeyEnter {
			clearLastLines(linesToClear)
			return value
		} else if key == keyboard.KeyArrowLeft && cursorPosition > 0 {
			cursorPosition--
			clearLastLines(linesToClear)
		} else if key == keyboard.KeyArrowRight && cursorPosition < len(value) {
			cursorPosition++
			clearLastLines(linesToClear)
		} else if (key == keyboard.KeyBackspace || key == keyboard.KeyBackspace2) && len(value) > 0 {
			value = value[:cursorPosition-1] + value[cursorPosition:]
			cursorPosition--
			clearLastLines(linesToClear)
		} else if key == keyboard.KeyCtrlC {
			clearLastLines(linesToClear)
			cleanup()
		} else if unicode.IsPrint(char) {
			value = value[:cursorPosition] + string(char) + value[cursorPosition:]
			cursorPosition++
			clearLastLines(linesToClear)
		} else if key == keyboard.KeySpace {
			value = value[:cursorPosition] + " " + value[cursorPosition:]
			cursorPosition++
			clearLastLines(linesToClear)
		} else {
			clearLastLines(linesToClear)
		}
	}
}

func main() {
	fmt.Print("\033[?25l")       // Hide cursor
	defer fmt.Print("\033[?25h") // Ensure cursor is shown on exit

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sigChan
		cleanup()
	}()

	var answers []string

	ans0 := NewInput("Are you smart", "", "")
	answers = append(answers, fmt.Sprintf("Favorite fruit: %v", ans0))
	fmt.Println("") //So that we have a line before our answers
	fmt.Println(answers[0])

	ans1 := NewSwitch("Are you smart", "", [2]string{"Yes", "No"}, true)
	answers = append(answers, fmt.Sprintf("Favorite fruit: %v", ans1))
	fmt.Println(answers[1])

	// Ask first question
	ans2 := NewSelect("Choose your favorite fruit", "", []string{"Apple", "Banana", "Orange", "Grape"})
	answers = append(answers, fmt.Sprintf("Favorite fruit: %s", ans2))
	fmt.Println(answers[2])
}

// var UnicodeDividers = Dividers{
// 	ALL: "┼",
// 	NES: "├",
// 	NSW: "┤",
// 	NEW: "┴",
// 	ESW: "┬",
// 	NE:  "└",
// 	NW:  "┘",
// 	SW:  "┐",
// 	ES:  "┌",
// 	EW:  "─",
// 	NS:  "│",
// }
// var UnicodeRoundedDividers = Dividers{
// 	ALL: "┼",
// 	NES: "├",
// 	NSW: "┤",
// 	NEW: "┴",
// 	ESW: "┬",
// 	NE:  "╰",
// 	NW:  "╯",
// 	SW:  "╮",
// 	ES:  "╭",
// 	EW:  "─",
// 	NS:  "│",
// }
