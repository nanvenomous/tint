// Package tint provides a simple API for loading terminal colorschemes from
// the tint config file (~/.config/colors.{yml,yaml,toml,json}).
package tint

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v3"
)

var (
	// ErrNotFound is returned when no tint config file is found
	ErrNotFound = errors.New("tint: config file not found")

	// ErrInvalidColor is returned when a color value is not in valid hex format
	ErrInvalidColor = errors.New("tint: invalid color format")

	// hexColorRegex validates hex color formats: #RGB, #RRGGBB, #RRGGBBAA
	hexColorRegex = regexp.MustCompile(`^#([0-9a-fA-F]{3}|[0-9a-fA-F]{6}|[0-9a-fA-F]{8})$`)
)

// ColorScheme represents a complete tint colorscheme.
type ColorScheme struct {
	// Core colors
	Background string `yaml:"background" toml:"background" json:"background"`
	Foreground string `yaml:"foreground" toml:"foreground" json:"foreground"`

	// Standard ANSI colors
	Colors struct {
		Black   string `yaml:"black" toml:"black" json:"black"`
		Red     string `yaml:"red" toml:"red" json:"red"`
		Green   string `yaml:"green" toml:"green" json:"green"`
		Yellow  string `yaml:"yellow" toml:"yellow" json:"yellow"`
		Blue    string `yaml:"blue" toml:"blue" json:"blue"`
		Magenta string `yaml:"magenta" toml:"magenta" json:"magenta"`
		Cyan    string `yaml:"cyan" toml:"cyan" json:"cyan"`
		White   string `yaml:"white" toml:"white" json:"white"`
	} `yaml:"colors" toml:"colors" json:"colors"`

	// Extended colors
	Accent    string `yaml:"accent,omitempty" toml:"accent,omitempty" json:"accent,omitempty"`
	Selection string `yaml:"selection,omitempty" toml:"selection,omitempty" json:"selection,omitempty"`
	Cursor    string `yaml:"cursor,omitempty" toml:"cursor,omitempty" json:"cursor,omitempty"`
	Comment   string `yaml:"comment,omitempty" toml:"comment,omitempty" json:"comment,omitempty"`
	Border    string `yaml:"border,omitempty" toml:"border,omitempty" json:"border,omitempty"`

	// Bright colors (optional)
	BrightColors struct {
		BrightBlack   string `yaml:"bright_black,omitempty" toml:"bright_black,omitempty" json:"bright_black,omitempty"`
		BrightRed     string `yaml:"bright_red,omitempty" toml:"bright_red,omitempty" json:"bright_red,omitempty"`
		BrightGreen   string `yaml:"bright_green,omitempty" toml:"bright_green,omitempty" json:"bright_green,omitempty"`
		BrightYellow  string `yaml:"bright_yellow,omitempty" toml:"bright_yellow,omitempty" json:"bright_yellow,omitempty"`
		BrightBlue    string `yaml:"bright_blue,omitempty" toml:"bright_blue,omitempty" json:"bright_blue,omitempty"`
		BrightMagenta string `yaml:"bright_magenta,omitempty" toml:"bright_magenta,omitempty" json:"bright_magenta,omitempty"`
		BrightCyan    string `yaml:"bright_cyan,omitempty" toml:"bright_cyan,omitempty" json:"bright_cyan,omitempty"`
		BrightWhite   string `yaml:"bright_white,omitempty" toml:"bright_white,omitempty" json:"bright_white,omitempty"`
	} `yaml:"bright_colors,omitempty" toml:"bright_colors,omitempty" json:"bright_colors,omitempty"`
}

// Load discovers and loads the tint colorscheme from the standard location.
// It checks $XDG_CONFIG_HOME/colors.{yml,yaml,toml,json} first, then falls
// back to ~/.config/colors.{yml,yaml,toml,json}.
//
// Returns ErrNotFound if no config file is found, allowing callers to fall
// back to application defaults.
func Load() (*ColorScheme, error) {
	path, err := findConfigFile()
	if err != nil {
		return nil, err
	}

	return LoadFrom(path)
}

// LoadFrom loads a colorscheme from the specified file path.
// The format is detected from the file extension.
func LoadFrom(path string) (*ColorScheme, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("tint: failed to read config: %w", err)
	}

	// Check file size (1 MB limit for security)
	if len(data) > 1024*1024 {
		return nil, errors.New("tint: config file too large (max 1 MB)")
	}

	scheme := &ColorScheme{}
	ext := strings.ToLower(filepath.Ext(path))

	switch ext {
	case ".yml", ".yaml":
		if err := yaml.Unmarshal(data, scheme); err != nil {
			return nil, fmt.Errorf("tint: failed to parse YAML: %w", err)
		}
	case ".toml":
		if err := toml.Unmarshal(data, scheme); err != nil {
			return nil, fmt.Errorf("tint: failed to parse TOML: %w", err)
		}
	case ".json":
		if err := json.Unmarshal(data, scheme); err != nil {
			return nil, fmt.Errorf("tint: failed to parse JSON: %w", err)
		}
	default:
		return nil, fmt.Errorf("tint: unsupported file format: %s", ext)
	}

	// Validate colors
	if err := scheme.Validate(); err != nil {
		return nil, err
	}

	return scheme, nil
}

// Validate checks that all present color values are in valid hex format.
// It normalizes colors to lowercase for consistency.
func (cs *ColorScheme) Validate() error {
	// Validate and normalize all color fields
	if err := cs.validateColor(&cs.Background, "background"); err != nil {
		return err
	}
	if err := cs.validateColor(&cs.Foreground, "foreground"); err != nil {
		return err
	}

	// Standard colors
	if err := cs.validateColor(&cs.Colors.Black, "colors.black"); err != nil {
		return err
	}
	if err := cs.validateColor(&cs.Colors.Red, "colors.red"); err != nil {
		return err
	}
	if err := cs.validateColor(&cs.Colors.Green, "colors.green"); err != nil {
		return err
	}
	if err := cs.validateColor(&cs.Colors.Yellow, "colors.yellow"); err != nil {
		return err
	}
	if err := cs.validateColor(&cs.Colors.Blue, "colors.blue"); err != nil {
		return err
	}
	if err := cs.validateColor(&cs.Colors.Magenta, "colors.magenta"); err != nil {
		return err
	}
	if err := cs.validateColor(&cs.Colors.Cyan, "colors.cyan"); err != nil {
		return err
	}
	if err := cs.validateColor(&cs.Colors.White, "colors.white"); err != nil {
		return err
	}

	// Extended colors (optional)
	if err := cs.validateColor(&cs.Accent, "accent"); err != nil {
		return err
	}
	if err := cs.validateColor(&cs.Selection, "selection"); err != nil {
		return err
	}
	if err := cs.validateColor(&cs.Cursor, "cursor"); err != nil {
		return err
	}
	if err := cs.validateColor(&cs.Comment, "comment"); err != nil {
		return err
	}
	if err := cs.validateColor(&cs.Border, "border"); err != nil {
		return err
	}

	// Bright colors (optional)
	if err := cs.validateColor(&cs.BrightColors.BrightBlack, "bright_colors.bright_black"); err != nil {
		return err
	}
	if err := cs.validateColor(&cs.BrightColors.BrightRed, "bright_colors.bright_red"); err != nil {
		return err
	}
	if err := cs.validateColor(&cs.BrightColors.BrightGreen, "bright_colors.bright_green"); err != nil {
		return err
	}
	if err := cs.validateColor(&cs.BrightColors.BrightYellow, "bright_colors.bright_yellow"); err != nil {
		return err
	}
	if err := cs.validateColor(&cs.BrightColors.BrightBlue, "bright_colors.bright_blue"); err != nil {
		return err
	}
	if err := cs.validateColor(&cs.BrightColors.BrightMagenta, "bright_colors.bright_magenta"); err != nil {
		return err
	}
	if err := cs.validateColor(&cs.BrightColors.BrightCyan, "bright_colors.bright_cyan"); err != nil {
		return err
	}
	if err := cs.validateColor(&cs.BrightColors.BrightWhite, "bright_colors.bright_white"); err != nil {
		return err
	}

	return nil
}

// validateColor checks if a color is in valid hex format and normalizes it.
// Empty colors are allowed (optional fields).
func (cs *ColorScheme) validateColor(color *string, fieldName string) error {
	if *color == "" {
		return nil // Empty is allowed for optional fields
	}

	if !hexColorRegex.MatchString(*color) {
		return fmt.Errorf("%w: %s = %q (must be #RGB, #RRGGBB, or #RRGGBBAA)", ErrInvalidColor, fieldName, *color)
	}

	// Normalize to lowercase
	*color = strings.ToLower(*color)

	return nil
}

// findConfigFile searches for the tint config file in standard locations.
// It checks $TINT_CONFIG_PATH first, then $XDG_CONFIG_HOME, then ~/.config.
func findConfigFile() (string, error) {
	// Check for explicit override
	if path := os.Getenv("TINT_CONFIG_PATH"); path != "" {
		if _, err := os.Stat(path); err == nil {
			return path, nil
		}
		return "", fmt.Errorf("tint: TINT_CONFIG_PATH set but file not found: %s", path)
	}

	// Determine config directory
	configDir := os.Getenv("XDG_CONFIG_HOME")
	if configDir == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf("tint: failed to get home directory: %w", err)
		}
		configDir = filepath.Join(home, ".config")
	}

	// Try each extension in order
	extensions := []string{".yml", ".yaml", ".toml", ".json"}
	for _, ext := range extensions {
		path := filepath.Join(configDir, "colors"+ext)
		if _, err := os.Stat(path); err == nil {
			return path, nil
		}
	}

	return "", ErrNotFound
}
