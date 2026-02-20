package main

import (
	"fmt"
	"log"

	"github.com/tint-colors/tint-go"
)

func main() {
	// Load the colorscheme from the standard tint config location
	// (~/.config/colors.{yml,yaml,toml,json})
	scheme, err := tint.Load()
	if err != nil {
		if err == tint.ErrNotFound {
			// No tint config found, use application defaults
			fmt.Println("No tint config found, using defaults")
			return
		}
		log.Fatalf("Failed to load tint config: %v", err)
	}

	// Use the colors in your application
	fmt.Printf("Background: %s\n", scheme.Background)
	fmt.Printf("Foreground: %s\n", scheme.Foreground)
	fmt.Println()

	// Access ANSI colors
	fmt.Println("ANSI Colors:")
	fmt.Printf("  Red:     %s\n", scheme.Colors.Red)
	fmt.Printf("  Green:   %s\n", scheme.Colors.Green)
	fmt.Printf("  Blue:    %s\n", scheme.Colors.Blue)
	fmt.Printf("  Yellow:  %s\n", scheme.Colors.Yellow)
	fmt.Printf("  Magenta: %s\n", scheme.Colors.Magenta)
	fmt.Printf("  Cyan:    %s\n", scheme.Colors.Cyan)
	fmt.Println()

	// Access extended colors
	if scheme.Accent != "" {
		fmt.Printf("Accent:    %s\n", scheme.Accent)
	}
	if scheme.Selection != "" {
		fmt.Printf("Selection: %s\n", scheme.Selection)
	}
	if scheme.Cursor != "" {
		fmt.Printf("Cursor:    %s\n", scheme.Cursor)
	}
}
