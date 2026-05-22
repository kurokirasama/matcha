# Specification: Google Contacts Sync (Native)

## Overview
Implement a native CLI command to synchronize and merge contacts from the Google People API into Matcha's local contact cache.

## Functional Requirements
- **OAuth Scope**: Update OAuth configuration to include `contacts.readonly`.
- **CLI Command**: Add `matcha contacts sync-google`.
- **Sync Engine**: Fetch contacts from Google, merge with local `contacts.json`, and handle duplicates.

## Acceptance Criteria
- [ ] `matcha contacts sync-google` successfully fetches Google contacts.
- [ ] Contacts are correctly merged into the local cache.
- [ ] Autocomplete in the composer uses the synced contacts.
