# Specification: Horizontal Email Preview

## Overview
This track implements a horizontal email preview mode (bottom split) and refactors the existing preview toggle into a consolidated "Split View" setting. Users can choose between standard full-screen viewing (OFF), side-by-side (Vertical), or top-bottom (Horizontal) layouts.

## Functional Requirements
- **Consolidated Layout Setting**:
    - Add a single `layout` configuration option to the settings, labeled as "Split View".
    - Supported values: `off`, `vertical`, `horizontal`.
    - Default value: `off`.
    - **CRITICAL**: Implementation must be strictly additive to ensure no existing configuration fields (SMTP, theme, language, etc.) are lost.
- **Horizontal Preview Pane**:
    - Implement a rendering mode where the email preview appears in the bottom half of the terminal screen.
    - The inbox message list should occupy the top half (50%).
    - Image rendering must support vertical offsets to prevent overlapping.
- **Mode Switching Logic**:
    - **OFF (Official Default)**: Highlighting an email in the inbox list does nothing automatically. Pressing `Enter` explicitly opens the email in a full-screen view.
    - **Vertical (Official Split)**: Highlighting an email automatically fetches and displays it in the preview pane on the right side.
    - **Horizontal**: Highlighting an email automatically fetches and displays it in the preview pane at the bottom.
- **Persistence**: The chosen layout must be saved in the application's configuration file.

## Acceptance Criteria
- [ ] A new "Split View" setting is available in the settings menu with "OFF", "Vertical", and "Horizontal" options.
- [ ] Selecting "Horizontal" correctly splits the inbox view into top (list) and bottom (preview) halves without image overlap.
- [ ] Selecting "Vertical" preserves the official side-by-side split behavior.
- [ ] Selecting "OFF" preserves the official behavior where opening an email uses the full terminal area, and highlighting does nothing.
- [ ] The setting persists across application restarts without affecting other user data.
- [ ] **NO Shift+L quick toggle functionality is implemented.**