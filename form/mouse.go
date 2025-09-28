package form

import (
	"fmt"
	"log"

	"github.com/gdamore/tcell/v2"
)

func main() {
	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatal(err)
	}
	if err := s.Init(); err != nil {
		log.Fatal(err)
	}
	defer s.Fini()

	// Enable mouse tracking
	s.EnableMouse()

	for {
		ev := s.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventMouse:
			x, y := ev.Position()
			button := ev.Buttons()
			fmt.Printf("Mouse at (%d, %d), buttons: %v\n", x, y, button)

			if button == tcell.Button1 {
				fmt.Println("Left click!")
			} else if button == tcell.WheelUp {
				fmt.Println("Scrolled up!")
			} else if button == tcell.WheelDown {
				fmt.Println("Scrolled down!")
			}

		case *tcell.EventResize:
			width, height := ev.Size()
			fmt.Printf("Terminal resized to %dx%d\n", width, height)

		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape || ev.Rune() == 'q' {
				return
			}
		}
	}
}
