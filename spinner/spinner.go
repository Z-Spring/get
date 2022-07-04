package spinner

import (
	"fmt"
	"time"
)

// A simple program demonstrating the spinner component from the Bubbles
// component library.

func Spinner(delay time.Duration) {
	Frames := []string{"⣾ ", "⣽ ", "⣻ ", "⢿ ", "⡿ ", "⣟ ", "⣯ ", "⣷ "}
	for {
		for _, r := range Frames {
			fmt.Printf("\r%s Loading ...", r)
			time.Sleep(delay)
		}
	}
}
