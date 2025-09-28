package form

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/term"
)

func Resize() {
	// Get initial size
	width, height, _ := term.GetSize(int(os.Stdout.Fd()))
	fmt.Printf("Initial size: %dx%d\n", width, height)

	// Listen for SIGWINCH
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGWINCH)

	go func() {
		for range sigChan {
			width, height, err := term.GetSize(int(os.Stdout.Fd()))
			if err != nil {
				fmt.Println("Resize detected, but failed to get size")
			} else {
				fmt.Printf("Terminal resized to %dx%d\n", width, height)
				// Redraw your UI here!
			}
		}
	}()

	fmt.Println("Press Ctrl+C to exit...")
	select {} // block forever
}
