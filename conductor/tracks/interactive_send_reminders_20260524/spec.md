# Track Specification

## Overview
Integrate attachment and empty subject reminders natively into the composer. These reminders will alert users if they attempt to send an email or exit the composer while an attachment seems missing (based on keyword triggers) or the subject is empty. The warnings will appear directly in the composer view and in the custom composer exit dialog (if enabled).

## Functional Requirements
- **Triggers**:
    - Subject Reminder: Triggered if the subject line is empty.
    - Attachment Reminder: Triggered if the email body contains specific keywords (e.g., "attach", "attached", "enclosed") but no files are attached.
    - Both features must be independently toggled via settings.
- **Composer Screen Integration**:
    - Display warning messages dynamically in the composer screen when triggers are activated.
- **Enhanced Exit Dialog Integration**:
    - If `EnableEnhancedComposerExit` is active, the custom exit dialog must display the active warnings.
- **Pre-Send Warning Dialog**:
    - If the user attempts to send the email (either from the composer or the custom exit dialog) and a trigger is active, intercept the send action.
    - Display a warning dialog summarizing the issues (e.g., "Missing Attachment", "Empty Subject").
    - The dialog must present the same options as the custom composer dialog (`[s]end`, `[a]bort`, `[d]ave`, `[c]ancel`).
    - **Exception**: If the user presses `s` (Send) from *within* this warning dialog, the email sends despite the warnings.
    - If the user presses `c` (Cancel), return to the main composer editing view.

## Visual Presentation
- Warnings must be styled using the application's dynamic accent color (defined by the active theme), avoiding hardcoded colors.

## Acceptance Criteria
- [ ] Users can enable/disable Subject and Attachment reminders in settings.
- [ ] Attempting to send an email with an empty subject (when enabled) prompts the warning dialog.
- [ ] Attempting to send an email with attachment keywords but no attachments (when enabled) prompts the warning dialog.
- [ ] The warning dialog allows sending anyway, saving as draft, aborting, or canceling back to the composer.
- [ ] Warnings are clearly visible in the enhanced exit dialog (if enabled).
- [ ] Warning text respects the active theme's accent color.

## Out of Scope
- NLP-based intent detection for attachments (relying strictly on static keyword matching).