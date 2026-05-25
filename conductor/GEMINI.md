# Conductor Standards (Matcha)

## Source Control & Privacy
1.  **Branch Isolation**: The `private` branch is **strictly local**. It contains project tracking data, design drafts, and user-specific memory.
2.  **MANDATORY: NO PUSHING**: Never push the `private` branch to any remote (`origin`, `upstream`, etc.).
3.  **Synchronization Rules**: The `git-sync` protocol or any automated sync loop MUST only be applied to public/shared branches (e.g., `master`, `main`).
4.  **Local History**: If you are working on the `private` branch, perform your commits but **HALT** before any push command.

## Workflow Integration
- All implementation tracks must follow the `conductor/workflow.md` protocol.
- Documentation synchronization is required for every completed track.
- Use `git notes` to attach technical summaries to commits for better auditability.
