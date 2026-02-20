# tint

**One colorscheme file. Every terminal app. No build step.**

---

The Unix terminal ecosystem has a theming problem. Not a lack of themes — there are thousands.
The problem is that every app has its own format, its own config location, and its own idea of what a "colorscheme" is.
Getting a consistent look across your terminal emulator, shell prompt, fzf, bat, delta, lazygit, and tmux requires either a templating pipeline, a lot of patience, or giving up.

`tint` is a proposed open standard that tries to fix this the Unix way: a single, simple file that apps can opt into reading directly.

---

## The File

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

That's it. YAML, TOML, and JSON are all supported — use whichever you prefer. The file lives at `~/.config/colors.yml` (or `.toml` / `.json`). Apps that support `tint` read it at startup. Change the file, restart your apps, done.

---

## Why Not Base16 / pywal / Flavours?

These are good tools and inspired parts of this spec. But they all share the same model: a build step that generates app-specific configs from a source palette. That means intermediate files, config stomping, and a pipeline you have to remember to re-run.

`tint` takes a different approach. There is no build step. There is no intermediate output. The file *is* the source of truth, and apps read it directly. The goal is that "theming your terminal" eventually means editing one file.

---

## For App Maintainers

Supporting `tint` in your app should take less than an afternoon. The spec is intentionally minimal — if `~/.config/colors.yml` exists, read it and use the values. If it doesn't, fall back to your existing defaults. Nothing breaks. Nothing is required.

SDKs are available to make this trivial:

| Language | Package |
|----------|---------|
| Go       | `github.com/tint-colors/tint-go` |
| Rust     | `tint` (crates.io) |
| C        | `libtint` |
| Python   | `tint-colors` (PyPI) |
| Lua      | `tint.nvim` |
| Node     | `@tint-colors/tint` (npm) |

A typical integration looks like this in Go:

```go
scheme, err := tint.Load()
if err != nil {
    // no colors.yml found, use defaults
    return
}
myApp.Background = scheme.Background
myApp.Accent = scheme.Accent
```

The SDKs handle format detection, validation, and fallback behavior. You don't need to think about YAML vs TOML.

---

## The Compiled Format

For apps written in C or other languages where adding a YAML/TOML parser is a non-starter, `tint` provides a CLI that compiles `colors.yml` to a dead-simple flat format that's trivially parseable with no dependencies:

```sh
tint compile
# writes ~/.cache/tint/colors
```

```
background=#1e1e2e
foreground=#cdd6f4
red=#f38ba8
green=#a6e3a1
...
```

`libtint` reads this compiled format only — zero dependencies, a single header file, works anywhere.

---

## Neovim

Neovim's highlight group system is more expressive than a flat palette can describe, and `tint` doesn't try to replace it. Instead, `tint.nvim` is a community plugin that reads `~/.config/colors.yml` and generates a full highlight group mapping from it — giving you a complete, cohesive neovim theme that matches the rest of your system automatically.

```lua
require('tint').setup()
-- that's it
```

---

## Scope

`tint` is intentionally focused on **terminal applications**. GUI toolkit theming (GTK, Qt), desktop environments, and compositor aesthetics are out of scope. Keeping the scope narrow is what makes the spec tractable and the SDKs small.

---

## Status

This is an early-stage proposal. The spec is stable. The SDKs are in active development.
If you maintain a terminal application and want to add support, open an issue or PR — early adopters will be listed here.

---

## Contributing

The most valuable contributions right now are:

- Native integrations in popular TUI apps
- SDK implementations in additional languages
- Feedback on the spec before it stabilizes at 1.0

See [contributing.md](./contributing.md) for details.

---

## Spec

The full specification lives in [spec.md](./spec.md). It is intentionally short.

---

*`tint` is opt-in, dependency-light, and doesn't touch your existing configs. It just opens a door.*
