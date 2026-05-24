# Initial Concept
A powerful, feature-rich terminal email client built with Go and the Bubble Tea TUI framework.

# Product Definition
**Product Vision**
Matcha aims to become the standard terminal-based email client, providing a beautiful, modern, and keyboard-driven experience for those who live in the command line. It bridges the gap between traditional terminal tools and modern email features like AI assistance and extensibility.

**Target Audience**
*   **Developers & CLI Power Users:** Users who prioritize speed, keyboard shortcuts, and a seamless workflow within their existing terminal environment.
*   **Privacy-Conscious Users:** Individuals who require robust security through PGP/GPG and hardware token integration (like YubiKey).
*   **Terminal Enthusiasts:** Anyone looking for a high-quality, aesthetically pleasing TUI experience for daily communication.

**Primary Goals**
*   **Aesthetics & UX:** Leveraging the Bubble Tea framework to provide a modern, responsive, and intuitive terminal user interface.
*   **Protocol Versatility:** Ensuring reliable and performant support for IMAP, JMAP, POP3, and Maildir, allowing users to consolidate all their accounts in one place.
*   **Extensibility:** Maintaining a robust plugin system that empowers the community to build and share custom functionality.

**Core Features**
*   **Background Synchronization (Daemon Mode):** A dedicated background process that ensures emails are always up-to-date and notifications are timely.
*   **AI Integration:** Built-in support for LLMs to assist with rewriting, summarizing, and drafting emails directly within the client.
*   **AI Agent & CLI Support:** A non-interactive mode (`matcha send`) that allows scripts and autonomous agents to manage email communications programmatically.
*   **Advanced Security:** Deep integration with PGP/GPG for signing and encrypting messages, including support for OpenPGP smart cards.
*   **Flexible Layouts:** Support for multiple viewing modes, including side-by-side (Vertical) and top-bottom (Horizontal) split views. Includes a **Layout Quick Toggle** (Shift+L) for rapid cycling between layouts.
*   **Extended Composer Exit Actions:** A rich, keyboard-driven confirmation dialog for the email composer, allowing users to quickly send, save, abort, or cancel when exiting.
