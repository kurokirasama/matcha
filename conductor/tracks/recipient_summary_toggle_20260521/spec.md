# Specification: Recipient Summary & Toggle

## Overview
Improve the email view header by summarizing long recipient lists. A toggle (`v`) allows users to expand the full list when needed.

## Functional Requirements
- **Smart Summary**: If an email has more than 3 recipients, show "Recip 1, Recip 2, Recip 3 ... and X more".
- **Expand Toggle**: Pressing `v` in the email view toggles between summarized and full recipient lists.
- **Dynamic Resizing**: The message viewport should adjust its height when the header expands/contracts.

## Acceptance Criteria
- [ ] Recipient list is summarized by default for large groups.
- [ ] `v` toggles the full list.
- [ ] Viewport height updates correctly.
