package form

import (
	"fmt"
	"sync"
	"time"
	"unicode"

	"github.com/eiannone/keyboard"
)

func NewInput(inputPrompt InputPrompt) (string, error) {
	initialized := false
	value := ""
	cursorPosition := 0
	var mu sync.Mutex

	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer keyboard.Close()

	keyCh := make(chan keyboard.KeyEvent, 10)

	go func() {
		for {
			char, key, err := keyboard.GetKey()
			if err != nil {
				close(keyCh)
				return
			}
			keyCh <- keyboard.KeyEvent{Key: key, Rune: char}
		}
	}()

	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	showCursor := true

	for {
		mu.Lock()
		currentValue := value
		currentPos := cursorPosition
		mu.Unlock()

		linesToClear := 5
		if initialized {
			clearLastLines(linesToClear)
		} else {
			initialized = true
		}

		fmt.Println("")
		fmt.Println(addLinePrefix(formatTitle(inputPrompt.Question)))

		displayValue := currentValue
		if displayValue == "" {
			if inputPrompt.Placeholder != "" {
				if showCursor {
					displayValue = fmt.Sprintf("%s%s",
						backgroundColorize(colorize(string(inputPrompt.Placeholder[0]), 235), 230),
						colorize(inputPrompt.Placeholder[1:], 242),
					)
				} else {
					displayValue = colorize(inputPrompt.Placeholder, 242)
				}
			} else {
				if showCursor {
					displayValue = backgroundColorize(" ", 230)
				} else {
					displayValue = " "
				}
			}
		} else {
			partBeforeCursor := displayValue[:currentPos]
			var partAtCursor string
			var partAfterCursor string
			if len(currentValue) == currentPos {
				if showCursor {
					partAtCursor = backgroundColorize(" ", 230)
				} else {
					partAtCursor = " "
				}
				partAfterCursor = ""
			} else {
				char := string(displayValue[currentPos])
				if showCursor {
					partAtCursor = backgroundColorize(colorize(char, 235), 230)
				} else {
					partAtCursor = char
				}
				partAfterCursor = displayValue[currentPos+1:]
			}

			displayValue = partBeforeCursor + partAtCursor + partAfterCursor
		}
		fmt.Println(addLinePrefix(displayValue))

		renderInputInstructions()

		select {
		case event, ok := <-keyCh:
			if !ok {
				return currentValue, nil
			}

			mu.Lock()
			showCursor = true
			ticker.Reset(500 * time.Millisecond)

			key, char := event.Key, event.Rune

			if key == keyboard.KeyEnter {
				mu.Unlock()
				clearLastLines(linesToClear)
				return value, nil
			} else if key == keyboard.KeyArrowLeft && cursorPosition > 0 {
				cursorPosition--
			} else if key == keyboard.KeyArrowRight && cursorPosition < len(value) {
				cursorPosition++
			} else if (key == keyboard.KeyBackspace || key == keyboard.KeyBackspace2) && len(value) > 0 {
				value = value[:cursorPosition-1] + value[cursorPosition:]
				cursorPosition--
			} else if key == keyboard.KeyCtrlC {
				mu.Unlock()
				keyboard.Close()
				clearLastLines(linesToClear)
				Cleanup()
			} else if unicode.IsPrint(char) {
				value = value[:cursorPosition] + string(char) + value[cursorPosition:]
				cursorPosition++
			} else if key == keyboard.KeySpace {
				value = value[:cursorPosition] + " " + value[cursorPosition:]
				cursorPosition++
			}

			mu.Unlock()
		case <-ticker.C:
			showCursor = !showCursor
		}
	}
}
