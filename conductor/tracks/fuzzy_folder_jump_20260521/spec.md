# Specification: Fuzzy Folder Jump (Native)

## Overview
Enhance the folder navigation by adding a native "Jump to Folder" feature with fuzzy search. This uses the `j` key to open a searchable menu of all available folders.

## Functional Requirements
- **Fuzzy Picker**: Implement a searchable list component for folders.
- **Jump Keybinding**: Map `j` in the inbox/folder context to open the picker.
- **Navigation**: Selecting a folder in the picker immediately switches the view to that folder.

## Acceptance Criteria
- [ ] Pressing `j` opens a fuzzy search menu of folders.
- [ ] Selecting a folder switches the mailbox.
- [ ] Integration with existing folder discovery logic.
