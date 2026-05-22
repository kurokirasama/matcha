# Implementation Plan: Google Contacts Sync (Native)

## Phase 1: OAuth & API
- [ ] Task: Update OAuth scopes in `config/oauth.go` (or relevant file).
- [ ] Task: Implement Google People API client in a new Go package.
- [ ] Task: Conductor - User Manual Verification 'API Client' (Protocol in workflow.md)

## Phase 2: CLI & Sync
- [ ] Task: Add the `sync-google` subcommand to the CLI.
- [ ] Task: Implement the merge logic in `config/contacts.go`.
- [ ] Task: Conductor - User Manual Verification 'Sync Logic' (Protocol in workflow.md)

## Phase 3: Verification
- [ ] Task: Run sync and check `contacts.json`.
- [ ] Task: Verify autocomplete in TUI.
