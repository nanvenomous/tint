# Contributing to tint

Thank you for your interest in contributing to `tint`! This document provides guidelines for contributing to the project.

---

## Ways to Contribute

### 1. Native Integrations

The most impactful contribution is adding native `tint` support to existing terminal applications. This means:

- Modifying the app to read `~/.config/colors.{yml,yaml,toml,json}` directly
- Using the appropriate SDK for the app's language
- Maintaining backwards compatibility with existing config

**How to submit:**
1. Add tint support to your app or submit a PR to an existing app
2. Open an issue here to be listed in the README as an official integration
3. Include a link to the PR/release and usage instructions

### 2. SDK Implementations

SDKs make it trivial for app maintainers to add tint support. We need implementations in:

- **High Priority**: C (libtint), Python, Rust, Lua
- **Nice to have**: JavaScript/Node, Ruby, Zig, Swift

**Requirements for SDKs:**
- Follow the [spec](./SPEC.md) exactly
- Support all three formats (YAML, TOML, JSON)
- Graceful error handling (no config = no error)
- Minimal dependencies
- Comprehensive tests
- Clear documentation

**How to submit:**
1. Create a new directory: `sdk/<language>/`
2. Follow the structure of `sdk/go/` as a reference
3. Include: implementation, tests, README, and at least one example
4. Open a PR with your implementation

### 3. Spec Feedback

The spec is intentionally minimal, but it's not set in stone yet. If you have feedback:

- **Before 1.0**: Open an issue to discuss changes
- **After 1.0**: Only backwards-compatible additions will be considered

**Good feedback includes:**
- Real-world use cases that aren't covered
- Ambiguities in the current spec
- Security concerns
- Implementation challenges from SDK development

---

## Development Setup

### Go SDK

```bash
cd sdk/go
go mod download
go test -v
```

### Running Examples

```bash
# Create a test config
mkdir -p ~/.config
cp examples/catppuccin-mocha.yml ~/.config/colors.yml

# Run an example
cd sdk/go
go run examples/basic/main.go
```

---

## Code Guidelines

### General

- Write clear, idiomatic code for the target language
- Include comprehensive tests (aim for >80% coverage)
- Document all public APIs
- Follow existing code style in the repo

### Go SDK

- Follow [Effective Go](https://golang.org/doc/effective_go)
- Use `gofmt` for formatting
- Run `go vet` before submitting
- Add tests for all new functionality

### Other Languages

- Follow community conventions for that language
- Use standard linters and formatters
- Match the structure and style of the Go SDK where applicable

---

## Pull Request Process

1. **Fork and branch**: Create a feature branch from `main`
2. **Implement**: Write your code, tests, and documentation
3. **Test**: Ensure all tests pass and add new tests as needed
4. **Document**: Update relevant README files and add examples
5. **PR**: Open a pull request with a clear description

**PR Title Format:**
- `[sdk/go] Add XYZ feature`
- `[spec] Clarify color format requirements`
- `[examples] Add Nord theme`

**PR Description Should Include:**
- What problem does this solve?
- How was it tested?
- Any breaking changes?
- Related issues (if any)

---

## Testing

All SDKs should include tests for:

- ✅ Loading YAML, TOML, and JSON formats
- ✅ File discovery (XDG_CONFIG_HOME, ~/.config)
- ✅ Environment variable override (TINT_CONFIG_PATH)
- ✅ Graceful fallback when config doesn't exist
- ✅ Color validation (valid/invalid hex formats)
- ✅ Error handling for malformed files
- ✅ Bright colors and extended fields (optional)

---

## Documentation

When adding a new SDK:

1. Create `sdk/<language>/README.md` with:
   - Installation instructions
   - Quick start example
   - API reference
   - At least one complete example
2. Update the main README.md to list your SDK
3. Add at least one working example in `sdk/<language>/examples/`

---

## Themes

We welcome high-quality theme examples in the `examples/` directory.

**Requirements:**
- Include YAML, TOML, OR JSON version (pick one)
- Include a comment header with theme name and credit
- Use a well-known, established theme (not custom)
- Test that it loads correctly with the Go SDK

**Popular themes we'd love to see:**
- Nord
- Solarized (Dark/Light)
- Tokyo Night
- One Dark
- Monokai

---

## Release Process

(For maintainers)

1. Update version in all SDK package files
2. Update CHANGELOG.md
3. Tag release: `git tag -a v1.0.0 -m "Release v1.0.0"`
4. Push: `git push origin v1.0.0`
5. Publish SDKs to package registries (crates.io, npm, PyPI, etc.)

---

## Code of Conduct

- Be respectful and constructive
- Focus on what's best for the project and users
- Assume good intent
- Provide actionable feedback

---

## Questions?

- Open an issue for questions about contributing
- Tag it with `question` label
- We'll do our best to respond within a few days

---

Thank you for contributing to make terminal theming better for everyone!
