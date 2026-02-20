package tint

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestLoadFrom_YAML(t *testing.T) {
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "colors.yml")

	yamlContent := `background: "#1e1e2e"
foreground: "#cdd6f4"

colors:
  black: "#45475a"
  red: "#f38ba8"
  green: "#a6e3a1"
  yellow: "#f9e2af"
  blue: "#89b4fa"
  magenta: "#cba4f7"
  cyan: "#94e2d5"
  white: "#bac2de"

accent: "#89b4fa"
selection: "#313244"
`

	if err := os.WriteFile(configPath, []byte(yamlContent), 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	scheme, err := LoadFrom(configPath)
	if err != nil {
		t.Fatalf("LoadFrom() error = %v", err)
	}

	// Check core colors
	if scheme.Background != "#1e1e2e" {
		t.Errorf("Background = %q, want %q", scheme.Background, "#1e1e2e")
	}
	if scheme.Foreground != "#cdd6f4" {
		t.Errorf("Foreground = %q, want %q", scheme.Foreground, "#cdd6f4")
	}

	// Check ANSI colors
	if scheme.Colors.Red != "#f38ba8" {
		t.Errorf("Colors.Red = %q, want %q", scheme.Colors.Red, "#f38ba8")
	}
	if scheme.Colors.Blue != "#89b4fa" {
		t.Errorf("Colors.Blue = %q, want %q", scheme.Colors.Blue, "#89b4fa")
	}

	// Check extended colors
	if scheme.Accent != "#89b4fa" {
		t.Errorf("Accent = %q, want %q", scheme.Accent, "#89b4fa")
	}
	if scheme.Selection != "#313244" {
		t.Errorf("Selection = %q, want %q", scheme.Selection, "#313244")
	}
}

func TestLoadFrom_TOML(t *testing.T) {
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "colors.toml")

	tomlContent := `background = "#1e1e2e"
foreground = "#cdd6f4"
accent = "#89b4fa"
selection = "#313244"

[colors]
black = "#45475a"
red = "#f38ba8"
green = "#a6e3a1"
yellow = "#f9e2af"
blue = "#89b4fa"
magenta = "#cba4f7"
cyan = "#94e2d5"
white = "#bac2de"
`

	if err := os.WriteFile(configPath, []byte(tomlContent), 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	scheme, err := LoadFrom(configPath)
	if err != nil {
		t.Fatalf("LoadFrom() error = %v", err)
	}

	if scheme.Background != "#1e1e2e" {
		t.Errorf("Background = %q, want %q", scheme.Background, "#1e1e2e")
	}
	if scheme.Colors.Cyan != "#94e2d5" {
		t.Errorf("Colors.Cyan = %q, want %q", scheme.Colors.Cyan, "#94e2d5")
	}
}

func TestLoadFrom_JSON(t *testing.T) {
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "colors.json")

	jsonContent := `{
  "background": "#1e1e2e",
  "foreground": "#cdd6f4",
  "accent": "#89b4fa",
  "selection": "#313244",
  "colors": {
    "black": "#45475a",
    "red": "#f38ba8",
    "green": "#a6e3a1",
    "yellow": "#f9e2af",
    "blue": "#89b4fa",
    "magenta": "#cba4f7",
    "cyan": "#94e2d5",
    "white": "#bac2de"
  }
}`

	if err := os.WriteFile(configPath, []byte(jsonContent), 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	scheme, err := LoadFrom(configPath)
	if err != nil {
		t.Fatalf("LoadFrom() error = %v", err)
	}

	if scheme.Background != "#1e1e2e" {
		t.Errorf("Background = %q, want %q", scheme.Background, "#1e1e2e")
	}
	if scheme.Colors.Green != "#a6e3a1" {
		t.Errorf("Colors.Green = %q, want %q", scheme.Colors.Green, "#a6e3a1")
	}
}

func TestLoadFrom_InvalidFormat(t *testing.T) {
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "colors.txt")

	if err := os.WriteFile(configPath, []byte("not a valid format"), 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	_, err := LoadFrom(configPath)
	if err == nil {
		t.Error("LoadFrom() expected error for unsupported format, got nil")
	}
}

func TestLoadFrom_InvalidYAML(t *testing.T) {
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "colors.yml")

	invalidYAML := `background: "#1e1e2e"
foreground: [invalid: syntax
`

	if err := os.WriteFile(configPath, []byte(invalidYAML), 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	_, err := LoadFrom(configPath)
	if err == nil {
		t.Error("LoadFrom() expected error for invalid YAML, got nil")
	}
}

func TestValidate_InvalidColors(t *testing.T) {
	tests := []struct {
		name   string
		scheme ColorScheme
		want   error
	}{
		{
			name: "invalid background color",
			scheme: ColorScheme{
				Background: "not-a-color",
				Foreground: "#ffffff",
			},
			want: ErrInvalidColor,
		},
		{
			name: "missing hash",
			scheme: ColorScheme{
				Background: "1e1e2e",
				Foreground: "#ffffff",
			},
			want: ErrInvalidColor,
		},
		{
			name: "invalid hex characters",
			scheme: ColorScheme{
				Background: "#gggggg",
				Foreground: "#ffffff",
			},
			want: ErrInvalidColor,
		},
		{
			name: "valid 3-char hex",
			scheme: ColorScheme{
				Background: "#fff",
				Foreground: "#000",
			},
			want: nil,
		},
		{
			name: "valid 6-char hex",
			scheme: ColorScheme{
				Background: "#ffffff",
				Foreground: "#000000",
			},
			want: nil,
		},
		{
			name: "valid 8-char hex with alpha",
			scheme: ColorScheme{
				Background: "#ffffffff",
				Foreground: "#000000ff",
			},
			want: nil,
		},
		{
			name: "uppercase hex",
			scheme: ColorScheme{
				Background: "#FFFFFF",
				Foreground: "#000000",
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.scheme.Validate()
			if !errors.Is(err, tt.want) {
				t.Errorf("Validate() error = %v, want %v", err, tt.want)
			}

			// If validation passed, check that colors were normalized to lowercase
			if err == nil {
				if tt.scheme.Background != "" && tt.scheme.Background[0] == '#' {
					if tt.scheme.Background != strings.ToLower(tt.scheme.Background) {
						t.Errorf("Background not normalized to lowercase: %q", tt.scheme.Background)
					}
				}
			}
		})
	}
}

func TestFindConfigFile(t *testing.T) {
	// Test that it returns ErrNotFound when no file exists
	t.Setenv("XDG_CONFIG_HOME", t.TempDir())
	t.Setenv("TINT_CONFIG_PATH", "")

	_, err := findConfigFile()
	if !errors.Is(err, ErrNotFound) {
		t.Errorf("findConfigFile() error = %v, want %v", err, ErrNotFound)
	}
}

func TestFindConfigFile_WithOverride(t *testing.T) {
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "custom-colors.yml")

	yamlContent := `background: "#000000"
foreground: "#ffffff"
`
	if err := os.WriteFile(configPath, []byte(yamlContent), 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	t.Setenv("TINT_CONFIG_PATH", configPath)

	found, err := findConfigFile()
	if err != nil {
		t.Fatalf("findConfigFile() error = %v", err)
	}

	if found != configPath {
		t.Errorf("findConfigFile() = %q, want %q", found, configPath)
	}
}

func TestFindConfigFile_PreferYML(t *testing.T) {
	tmpDir := t.TempDir()
	t.Setenv("XDG_CONFIG_HOME", tmpDir)
	t.Setenv("TINT_CONFIG_PATH", "")

	// Create both .yml and .toml files
	ymlPath := filepath.Join(tmpDir, "colors.yml")
	tomlPath := filepath.Join(tmpDir, "colors.toml")

	if err := os.WriteFile(ymlPath, []byte("background: '#000000'\nforeground: '#ffffff'"), 0644); err != nil {
		t.Fatalf("Failed to create yml file: %v", err)
	}
	if err := os.WriteFile(tomlPath, []byte("background = '#000000'\nforeground = '#ffffff'"), 0644); err != nil {
		t.Fatalf("Failed to create toml file: %v", err)
	}

	found, err := findConfigFile()
	if err != nil {
		t.Fatalf("findConfigFile() error = %v", err)
	}

	// Should prefer .yml over .toml
	if found != ymlPath {
		t.Errorf("findConfigFile() = %q, want %q (should prefer .yml)", found, ymlPath)
	}
}

func TestLoadFrom_FileSizeLimit(t *testing.T) {
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "colors.yml")

	// Create a file larger than 1 MB
	largeContent := make([]byte, 1024*1024+1)
	for i := range largeContent {
		largeContent[i] = 'a'
	}

	if err := os.WriteFile(configPath, largeContent, 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	_, err := LoadFrom(configPath)
	if err == nil || err.Error() != "tint: config file too large (max 1 MB)" {
		t.Errorf("LoadFrom() expected file size error, got %v", err)
	}
}

func TestBrightColors(t *testing.T) {
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "colors.yml")

	yamlContent := `background: "#1e1e2e"
foreground: "#cdd6f4"

colors:
  black: "#45475a"
  red: "#f38ba8"
  green: "#a6e3a1"
  yellow: "#f9e2af"
  blue: "#89b4fa"
  magenta: "#cba4f7"
  cyan: "#94e2d5"
  white: "#bac2de"

bright_colors:
  bright_black: "#585b70"
  bright_red: "#f38ba8"
  bright_green: "#a6e3a1"
  bright_yellow: "#f9e2af"
  bright_blue: "#89b4fa"
  bright_magenta: "#f5c2e7"
  bright_cyan: "#94e2d5"
  bright_white: "#cdd6f4"
`

	if err := os.WriteFile(configPath, []byte(yamlContent), 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	scheme, err := LoadFrom(configPath)
	if err != nil {
		t.Fatalf("LoadFrom() error = %v", err)
	}

	if scheme.BrightColors.BrightBlack != "#585b70" {
		t.Errorf("BrightColors.BrightBlack = %q, want %q", scheme.BrightColors.BrightBlack, "#585b70")
	}
	if scheme.BrightColors.BrightMagenta != "#f5c2e7" {
		t.Errorf("BrightColors.BrightMagenta = %q, want %q", scheme.BrightColors.BrightMagenta, "#f5c2e7")
	}
}
