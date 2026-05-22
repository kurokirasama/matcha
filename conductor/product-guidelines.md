# Product Guidelines: Matcha

**Prose Style & Tone**
*   **Professional & Technical:** Use clear, direct, and efficient language. Documentation and in-app messages should be precise, technical, and free of unnecessary fluff, respecting the user's expertise as a CLI power user.

**User Experience (UX) Principles**
*   **Keyboard-First Design:** Every action must be accessible via keyboard shortcuts. The interface should be optimized for speed, allowing users to navigate and manage emails without reaching for the mouse.
*   **Discoverability:** While keyboard-driven, the client must provide clear visual cues and easily accessible help menus (e.g., a '?' shortcut) to ensure features are discoverable by new and experienced users alike.
*   **Responsiveness & Async UI:** The interface must remain responsive at all times. Network-heavy tasks like fetching emails or sending messages must happen asynchronously in the background, ensuring the UI never freezes or blocks user input.

**Visual Identity & Branding**
*   **Matcha Green & Modern TUI:** The default visual identity centers on a modern "Matcha" green theme, utilizing the full capabilities of the Bubble Tea and Lipgloss frameworks to create a vibrant yet professional terminal experience.

**Error Handling**
*   **Concise & Unobtrusive:** Errors should be presented briefly and clearly, avoiding clutter in the main interface. Use a dedicated status line or subtle toast notifications for error reporting, and only interrupt the workflow for critical failures that require immediate user action.
