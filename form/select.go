package form

import (
	"fmt"

	"github.com/eiannone/keyboard"
)

func NewSelect(selectPrompt SelectPrompt) (SelectOption, error) {
	selected := 0
	initialized := false
	linesToClear := 4 + len(selectPrompt.Options)

	if selectPrompt.Description != "" {
		linesToClear += 2
	}

	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer keyboard.Close()

	for {
		if initialized {
			clearLastLines(linesToClear)
		} else {
			initialized = true
		}

		fmt.Println("")
		fmt.Println(addLinePrefix(formatTitle(selectPrompt.Question)))

		if selectPrompt.Description != "" {
			fmt.Println(addLinePrefix(formatDescription(selectPrompt.Description)))
			fmt.Println(addLinePrefix(""))
		}

		for i, option := range selectPrompt.Options {
			cursor := "  "
			renderOption := option.Label
			if i == selected {
				cursor = colorize("> ", 6)
				renderOption = colorize(option.Label, 107)
			}
			line := fmt.Sprintf("%s%s", cursor, renderOption)
			fmt.Println(addLinePrefix(line))
		}

		renderSelectInstructions()

		_, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}

		if key == keyboard.KeyArrowUp && selected > 0 {
			selected--
		} else if key == keyboard.KeyArrowDown && selected < len(selectPrompt.Options)-1 {
			selected++
		} else if key == keyboard.KeyEnter {
			clearLastLines(linesToClear)
			return selectPrompt.Options[selected], nil
		} else if key == keyboard.KeyCtrlC {
			clearLastLines(linesToClear)
			keyboard.Close()
			Cleanup()
		}
	}
}
