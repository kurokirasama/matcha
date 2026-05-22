# Implementation Plan: Search by Folder Syntax

## Phase 1: Parser Update
- [ ] Task: Modify `ParseSearchQuery` in `backend/backend.go` to handle folder prefixes.
- [ ] Task: Update `SearchQuery` struct to hold the target folder.
- [ ] Task: Conductor - User Manual Verification 'Parser' (Protocol in workflow.md)

## Phase 2: Backend Routing
- [ ] Task: Update search orchestration in `main.go` to respect the parsed folder.
- [ ] Task: Conductor - User Manual Verification 'Routing' (Protocol in workflow.md)

## Phase 3: Verification
- [ ] Task: Add unit tests for the parser.
- [ ] Task: Verify cross-folder search results.
