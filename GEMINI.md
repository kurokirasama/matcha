# Project Context: Matcha

Matcha is a powerful, feature-rich terminal email client built with Go and the Bubble Tea TUI framework. It provides a modern, keyboard-driven experience with support for multiple accounts, background synchronization, and a robust plugin system.

## Project Overview

- **Core Technology:** Go (1.26+) using the **Charm** ecosystem (Bubble Tea v2, Bubbles v2, Lipgloss v2).
- **Architecture:**
    - **Backends (`backend/`):** Pluggable support for IMAP, JMAP, POP3, and Maildir.
    - **Daemon (`daemon/`):** A background process for concurrent account synchronization and notifications.
    - **CLI (`cli/`):** Integration logic and subcommands (e.g., `matcha send`, `matcha daemon`).
    - **TUI (`tui/`, `view/`):** Responsive terminal interface components.
    - **Native Core (`clib/`):** Performance-critical tasks (HTML parsing, image conversion, markdown) implemented in C and integrated via CGO.
- **Key Features:** AI-assisted rewriting/summarization, PGP/GPG encryption (including YubiKey support), and 35+ community plugins.

## Building and Running

The project uses a `Makefile` to manage common tasks:

- **Build:** `make build` (outputs to `bin/matcha`)
- **Run:** `make run` (builds and executes the client)
- **Test:** `make test` (runs all unit tests)
- **Lint:** `make lint` (runs `go fmt` and `go vet`)
- **Nix:** Use `nix build` or `nix develop` (via `flake.nix`).

## Development Conventions

- **CGO & Native Code:** High-performance processing (HTML, images, markdown) resides in `clib/`. These components have pure Go fallbacks but prefer the C implementation for speed.
- **Branch Naming:**
    - `feature/description`
    - `bugfix/description`
    - `docs/description`
    - `refactor/description`
- **Commit Messages:** Follow a clear, concise style. Use `type(scope): subject` where applicable (e.g., `feat(backend): add JMAP search support`).
- **AI Policy:** AI-assisted contributions are welcome but must be fully understood and reviewed by the contributor. "Understand what you submit" is the primary rule.
- **Testing:** New features should include unit tests (standard Go `testing` package). For TUI components, use the established patterns in `main_test.go` or package-specific tests.

## Workspace Management (Conductor)

This project uses the **Conductor** methodology for task management and planning.
- **Project Context:** Located in `conductor/index.md`.
- **Tracks:** Active work is organized into "tracks" under `conductor/tracks/`. Each track has its own `spec.md` and `plan.md`.
- **Workflow:** Adheres to the protocols defined in `conductor/workflow.md`.
