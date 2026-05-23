# Specification: Multi-line Email Headers

## Overview
This track introduces a new "Multi-line Headers" setting that changes how email metadata is displayed in the reading view (`EmailView`). When enabled, the From, To, Cc, Bcc, and Subject fields will be displayed on separate lines in that specific order, rather than condensed into a single line.

## Functional Requirements
- **Settings Toggle**:
    - Add a boolean setting `EnableMultilineHeaders` to the application configuration.
    - Default value: `false` (OFF by default).
    - Add a corresponding toggle labeled "Multi-line Headers" in the General Settings menu.
    - Implementation must be strictly additive, preserving all existing config fields.
- **Header Rendering (`EmailView`)**:
    - If `EnableMultilineHeaders` is true, render headers across multiple lines in the order: From, To, Cc, Bcc, Subject.
    - Empty fields (e.g., no Cc or Bcc) must be omitted entirely to save vertical space.
    - If `EnableMultilineHeaders` is false, the application must retain the current single-line rendering behavior without any visual regression.

## Acceptance Criteria
- [ ] New `enable_multiline_headers` boolean added to configuration and persisted correctly.
- [ ] Setting is available in the General Settings UI with proper localizations.
- [ ] When enabled, opening an email displays From, To, Cc, Bcc, Subject on separate lines.
- [ ] Empty fields are completely hidden in multi-line mode.
- [ ] When disabled, the legacy single-line view is perfectly preserved.