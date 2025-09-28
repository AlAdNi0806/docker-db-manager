package form

import (
	"fmt"

	"github.com/eiannone/keyboard"
)

func NewSwitch(switchPrompt SwitchPrompt) (bool, error) {
	selected := switchPrompt.DefaultValue

	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer keyboard.Close()

	for {
		fmt.Println("")
		fmt.Println(addLinePrefix(formatTitle(switchPrompt.Question)))
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
			backgroundColorize(fmt.Sprintf("  %s  ", switchPrompt.Options[0]), leftBg),
			backgroundColorize(fmt.Sprintf("  %s  ", switchPrompt.Options[1]), rightBg),
		)

		fmt.Println(addLinePrefix(switchLine))

		renderSwitchInstructions()

		_, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}

		linesToClear := 6 // 2 for the instructions + 2 for the title plus line before + the switch line and the line before

		if key == keyboard.KeyArrowLeft && selected == false {
			selected = true
			clearLastLines(linesToClear)
		} else if key == keyboard.KeyArrowRight && selected == true {
			selected = false
			clearLastLines(linesToClear)
		} else if key == keyboard.KeyEnter {
			clearLastLines(linesToClear)
			return selected, nil
		} else if key == keyboard.KeyCtrlC {
			keyboard.Close()
			clearLastLines(linesToClear)
			Cleanup()
		} else {
			// we can have a bool that will indicate if the state changed if not we skip the rendering and go to the waiting again
			clearLastLines(linesToClear)
		}
	}
}
