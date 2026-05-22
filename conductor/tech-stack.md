# Technology Stack: Matcha

**Core Language & Runtime**
*   **Go (1.26.3):** The primary language for the entire application, chosen for its performance, concurrency model, and strong ecosystem for terminal tools.

**Terminal User Interface (TUI)**
*   **Charm Framework:**
    *   **Bubble Tea (v2):** The Elm-inspired framework for managing application state and rendering the TUI.
    *   **Bubbles (v2):** A library of pre-built UI components (like text inputs, lists, and viewports).
    *   **Lipgloss (v2):** Used for styling and layout definitions in the terminal.

**Email & Communication Protocols**
*   **IMAP (v2):** Supported via `github.com/emersion/go-imap/v2` for modern IMAP interactions.
*   **JMAP:** Supported via `git.sr.ht/~rockorager/go-jmap` for advanced, modern email synchronization.
*   **POP3:** Supported via `github.com/knadh/go-pop3`.
*   **Maildir:** Supported via `github.com/emersion/go-maildir`.

**Security & Encryption**
*   **OpenPGP:** Comprehensive support for signing and encrypting messages using `github.com/ProtonMail/go-crypto`.
*   **Hardware Tokens:** Support for OpenPGP smart cards (like YubiKey) via `cunicu.li/go-openpgp-card`.
*   **PKCS7:** Support for cryptographic message syntax via `go.mozilla.org/pkcs7`.

**Extensibility & Scripting**
*   **Gopher-Lua:** An implementation of the Lua programming language in Go, used to power the application's plugin system.

**Infrastructure & Deployment**
*   **Go Modules:** Dependency management.
*   **Nix (Flakes):** Reproducible builds and development environments.
*   **Goreleaser:** Automated release pipeline for multi-platform binaries.
*   **Snapcraft:** Deployment as a Snap package.
