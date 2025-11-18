# Contributing to ThorVG Go Bindings

Thank you for your interest in contributing to the ThorVG Go bindings!

## Getting Started

1. Fork the repository
2. Clone your fork: `git clone https://github.com/YOUR_USERNAME/thorvg.git`
3. Create a new branch: `git checkout -b feature/your-feature-name`
4. Make your changes
5. Test your changes
6. Commit and push
7. Create a Pull Request

## Development Setup

### Prerequisites

- Go 1.16 or later
- ThorVG C library installed
- pkg-config
- CGo enabled

### Installing ThorVG C Library

```bash
# Ubuntu/Debian
sudo apt-get install meson ninja-build pkg-config
git clone https://github.com/thorvg/thorvg.git
cd thorvg
meson build
ninja -C build install
```

## Code Style

- Follow standard Go formatting (`go fmt`)
- Run `go vet` to check for common mistakes
- Write clear, descriptive commit messages
- Add tests for new functionality
- Update documentation for API changes

## Testing

Run tests with:

```bash
make test
```

For tests that require ThorVG C library:

```bash
go test -v -tags=integration ./...
```

## Documentation

- Add godoc comments for all exported functions and types
- Update README.md for significant changes
- Add examples for new features

## Pull Request Process

1. Ensure your code passes all tests
2. Update the README.md with details of changes if applicable
3. Add any new dependencies to go.mod
4. The PR will be merged once reviewed and approved

## Code of Conduct

- Be respectful and inclusive
- Focus on constructive feedback
- Help others learn and grow

## Questions?

Feel free to open an issue for questions or discussions.

## License

By contributing, you agree that your contributions will be licensed under the MIT License.
