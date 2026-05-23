# Specification: Extended Composer Exit Actions

## Overview
This track refactors the email composer's exit confirmation dialog (triggered by `Esc`) to provide more robust options for managing the current draft. Instead of a simple yes/no for exit, the dialog will now offer shortcuts for sending, deleting, saving, or returning to the composition.

## Functional Requirements
- **Enhanced Exit Dialog**:
    - Modify the existing confirmation overlay triggered by the `Esc` key in the email composer.
    - The dialog must now present and support the following actions:
        - **`s`**: **Send** the email immediately.
        - **`a`**: **Abort** and permanently delete the current draft.
        - **`d`**: **Discard/Exit** but save the draft (equivalent to the previous "Yes").
        - **`c`**: **Cancel** and return to the composition view (equivalent to the previous "No").
- **Visual Presentation**:
    - Update the overlay UI to clearly list the new available keys and their actions.
- **Optional Toggle**:
    - Add a configuration option labeled "Enhanced Composer Exit" in the General Settings menu.
    - If disabled, the composer retains the legacy "y/n" exit dialog.
    - Default value: `false` (OFF).

## Acceptance Criteria
- [ ] New `enable_enhanced_composer_exit` boolean added to configuration and persisted correctly.
- [ ] Setting is available in the General Settings UI with proper localizations.
- [ ] When enabled, pressing `Esc` in the composer shows the new dialog with `s`, `a`, `d`, and `c` options.
- [ ] Pressing `s` correctly triggers the email send flow and exits the composer.
- [ ] Pressing `a` exits the composer without saving a draft.
- [ ] Pressing `d` exits the composer and saves the current draft.
- [ ] Pressing `c` closes the overlay and returns focus to the editor.
- [ ] When disabled, the original "y/n" prompt is displayed.

## Out of Scope
- Adding these keybinds as global shortcuts outside of the `Esc` dialog.
- Modifying the actual sending or saving logic (reusing existing functions).
