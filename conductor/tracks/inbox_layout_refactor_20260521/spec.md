# Specification: Inbox Layout Refactor & Horizontal Preview

## Overview
Refactor the inbox layout to utilize full terminal height and implement a toggleable horizontal preview mode (bottom-split) using the `L` key (since `l` is `next_tab`).

## Functional Requirements
- **Full-Height List**: Remove the hardcoded 50/50 vertical split when no preview is active.
- **Horizontal Preview**: Add a layout mode where the preview pane appears *below* the message list.
- **Layout Toggle**: Use `L` to cycle through layout modes:
    1. Full Height List (No preview)
    2. Vertical Split (Side-by-side)
    3. Horizontal Split (Top-bottom)

## Acceptance Criteria
- [ ] `L` key cycles through the three layout modes.
- [ ] List occupies full height when preview is disabled.
- [ ] Layout state persists across view switches.
