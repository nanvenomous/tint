# tint-go

Go SDK for [tint](https://github.com/tint-colors/tint) — one colorscheme file for all your terminal apps.

## Installation

```bash
go get github.com/tint-colors/tint-go
```

## Quick Start

```go
package main

import (
    "fmt"
    "log"

    "github.com/tint-colors/tint-go"
)

func main() {
    scheme, err := tint.Load()
    if err != nil {
        if err == tint.ErrNotFound {
            // No tint config found, use your app's defaults
            return
        }
        log.Fatal(err)
    }

    // Use the colors
    fmt.Printf("Background: %s\n", scheme.Background)
    fmt.Printf("Foreground: %s\n", scheme.Foreground)
    fmt.Printf("Red: %s\n", scheme.Colors.Red)
}
```

## Usage

### Loading the Config

The SDK looks for a config file at `~/.config/colors.{yml,yaml,toml,json}` (or `$XDG_CONFIG_HOME/colors.*`):

```go
scheme, err := tint.Load()
if err != nil {
    if err == tint.ErrNotFound {
        // No config file found - use your defaults
    } else {
        // Parse error or other issue
    }
}
```

### Loading from a Custom Path

```go
scheme, err := tint.LoadFrom("/path/to/colors.yml")
```

### Environment Variable Override

Users can override the config location with `TINT_CONFIG_PATH`:

```bash
export TINT_CONFIG_PATH=~/.config/my-custom-colors.yml
```

## Available Colors

### Core Colors

```go
scheme.Background  // Background color
scheme.Foreground  // Foreground/text color
```

### ANSI Colors

```go
scheme.Colors.Black
scheme.Colors.Red
scheme.Colors.Green
scheme.Colors.Yellow
scheme.Colors.Blue
scheme.Colors.Magenta
scheme.Colors.Cyan
scheme.Colors.White
```

### Extended Colors (Optional)

```go
scheme.Accent      // Primary accent color
scheme.Selection   // Selection/highlight background
scheme.Cursor      // Cursor color
scheme.Comment     // Comment/dim text color
scheme.Border      // Border/separator color
```

### Bright Colors (Optional)

```go
scheme.BrightColors.BrightBlack
scheme.BrightColors.BrightRed
// ... etc
```

All color values are hex strings (e.g., `"#1e1e2e"`), normalized to lowercase.

## Example: Terminal App Integration

```go
type MyApp struct {
    bg, fg, errorColor string
}

func NewApp() *MyApp {
    app := &MyApp{
        bg: "#000000",
        fg: "#ffffff",
        errorColor: "#ff0000",
    }

    // Try to load tint config
    if scheme, err := tint.Load(); err == nil {
        app.bg = scheme.Background
        app.fg = scheme.Foreground
        app.errorColor = scheme.Colors.Red
    }

    return app
}
```

## Config File Format

The SDK supports YAML, TOML, and JSON. Here's an example YAML config:

```yaml
# ~/.config/colors.yml

background: "#1e1e2e"
foreground: "#cdd6f4"

colors:
  black:   "#45475a"
  red:     "#f38ba8"
  green:   "#a6e3a1"
  yellow:  "#f9e2af"
  blue:    "#89b4fa"
  magenta: "#cba4f7"
  cyan:    "#94e2d5"
  white:   "#bac2de"

accent:    "#89b4fa"
selection: "#313244"
```

See the [tint specification](../../SPEC.md) for full details.

## Error Handling

The SDK returns specific errors you can check:

- `tint.ErrNotFound`: No config file found (use app defaults)
- `tint.ErrInvalidColor`: Color format validation failed
- Other errors: File read or parse errors

```go
scheme, err := tint.Load()
switch {
case errors.Is(err, tint.ErrNotFound):
    // Use defaults
case errors.Is(err, tint.ErrInvalidColor):
    // Invalid color in config
default:
    // Other error
}
```

## Features

- **Zero config required**: Falls back gracefully when no config exists
- **Multi-format**: Supports YAML, TOML, and JSON automatically
- **Validation**: Validates hex color formats
- **XDG compliant**: Respects `XDG_CONFIG_HOME`
- **Minimal dependencies**: Only format parsers (yaml, toml, json from stdlib)
- **Fast**: Config is parsed once at startup

## Examples

See the [examples](./examples) directory for complete examples:

- [basic](./examples/basic) - Simple color loading
- [terminal-app](./examples/terminal-app) - Full application integration

## License

MIT
