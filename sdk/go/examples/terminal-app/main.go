package main

import (
	"fmt"
	"log"

	"github.com/tint-colors/tint-go"
)

// TerminalApp represents a hypothetical terminal application
type TerminalApp struct {
	BackgroundColor string
	ForegroundColor string
	ErrorColor      string
	SuccessColor    string
	WarningColor    string
	AccentColor     string
}

// DefaultColors returns a TerminalApp with default colors
func DefaultColors() *TerminalApp {
	return &TerminalApp{
		BackgroundColor: "#000000",
		ForegroundColor: "#ffffff",
		ErrorColor:      "#ff0000",
		SuccessColor:    "#00ff00",
		WarningColor:    "#ffff00",
		AccentColor:     "#0000ff",
	}
}

// LoadFromTint configures the app with tint colors, falling back to defaults
func (app *TerminalApp) LoadFromTint() error {
	scheme, err := tint.Load()
	if err != nil {
		if err == tint.ErrNotFound {
			// No config found, keep defaults
			fmt.Println("Using default colors (no tint config found)")
			return nil
		}
		return fmt.Errorf("failed to load tint: %w", err)
	}

	// Apply tint colors
	app.BackgroundColor = scheme.Background
	app.ForegroundColor = scheme.Foreground
	app.ErrorColor = scheme.Colors.Red
	app.SuccessColor = scheme.Colors.Green
	app.WarningColor = scheme.Colors.Yellow

	// Use accent if available, otherwise fall back to blue
	if scheme.Accent != "" {
		app.AccentColor = scheme.Accent
	} else {
		app.AccentColor = scheme.Colors.Blue
	}

	fmt.Println("Loaded colors from tint config")
	return nil
}

func (app *TerminalApp) PrintColors() {
	fmt.Println("\nTerminal App Colors:")
	fmt.Printf("  Background: %s\n", app.BackgroundColor)
	fmt.Printf("  Foreground: %s\n", app.ForegroundColor)
	fmt.Printf("  Error:      %s\n", app.ErrorColor)
	fmt.Printf("  Success:    %s\n", app.SuccessColor)
	fmt.Printf("  Warning:    %s\n", app.WarningColor)
	fmt.Printf("  Accent:     %s\n", app.AccentColor)
}

func main() {
	// Create app with defaults
	app := DefaultColors()

	// Try to load tint config
	if err := app.LoadFromTint(); err != nil {
		log.Fatalf("Error: %v", err)
	}

	// Print the final colors
	app.PrintColors()
}
