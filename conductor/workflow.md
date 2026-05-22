# Project Workflow

## Guiding Principles

1. **The Plan is the Source of Truth:** All work must be tracked in `plan.md`.
2. **The Tech Stack is Deliberate:** Changes to the tech stack must be documented in `tech-stack.md` *before* implementation.
3. **Dual-Branch Architecture (Public Repo):**
    - **Public Branch (`master`/`main`):** For PRs to upstream. No Conductor/private artifacts.
    - **Private Branch (`private`):** High-hierarchy branch. Contains `conductor/`, `tests/`, `GEMINI.md`, `.gitignore`, and `todos.md`.
4. **Nushell-First Mandate:** Priority is given to Nushell pipelines for system interactions and data tasks. Activate `nushell-expert` before first use.
5. **Context Engineering Protocol:** Always follow the structured Discovery -> Synthesis -> Planning -> Execution workflow. Never code without a complete mental model.
6. **Test-Driven Development:** Write unit tests before implementing functionality.
7. **High Code Coverage:** Aim for >80% code coverage for all modules.
8. **User Experience First:** Every decision should prioritize user experience.
9. **Non-Interactive & CI-Aware:** Prefer non-interactive commands. Use `CI=true` for watch-mode tools.
10. **Mandatory Discord Notification (CRITICAL Sequence):**
    - `to-discord` MUST be executed and COMPLETED before `ask_user`.
    - `to-discord` is a Nushell command.
    - All reports for long tasks (>5min) follow this protocol.

## Task Workflow

All tasks follow a strict lifecycle:

### Standard Task Workflow

1. **Select Task:** Choose the next available task from `plan.md` in sequential order

2. **Mark In Progress:** Before beginning work, edit `plan.md` and change the task from `[ ]` to `[~]`

3. **Write Failing Tests (Red Phase):**
   - Create a new test file for the feature or bug fix.
   - Write one or more unit tests that clearly define the expected behavior and acceptance criteria for the task.
   - **CRITICAL:** Run the tests and confirm that they fail as expected. This is the "Red" phase of TDD. Do not proceed until you have failing tests.

4. **Implement to Pass Tests (Green Phase):**
   - Write the minimum amount of application code necessary to make the failing tests pass.
   - Run the test suite again and confirm that all tests now pass. This is the "Green" phase.

5. **Refactor (Optional but Recommended):**
   - With the safety of passing tests, refactor the implementation code and the test code to improve clarity, remove duplication, and enhance performance without changing the external behavior.
   - Rerun tests to ensure they still pass after refactoring.

6. **Verify Coverage:** Run coverage reports using the project's chosen tools. For example, in a Python project, this might look like:
   ```bash
   pytest --cov=app --cov-report=html
   ```
   Target: >80% coverage for new code. The specific tools and commands will vary by language and framework.

7. **Document Deviations:** If implementation differs from tech stack:
   - **STOP** implementation
   - Update `tech-stack.md` with new design
   - Add dated note explaining the change
   - Resume implementation

8. **Commit Code Changes:**
   - Stage all code changes related to the task.
   - Propose a clear, concise commit message e.g, `feat(ui): Create basic HTML structure for calculator`.
   - Perform the commit.

9. **Attach Task Summary with Git Notes:**
   - **Step 9.1: Get Commit Hash:** Obtain the hash of the *just-completed commit* (`git log -1 --format="%H"`).
   - **Step 9.2: Draft Note Content:** Create a detailed summary for the completed task. This should include the task name, a summary of changes, a list of all created/modified files, and the core "why" for the change.
   - **Step 9.3: Attach Note:** Use the `git notes` command to attach the summary to the commit.
     ```bash
     # The note content from the previous step is passed via the -m flag.
     git notes add -m "<note content>" <commit_hash>
     ```

10. **Get and Record Task Commit SHA:**
    - **Step 10.1: Update Plan:** Read `plan.md`, find the line for the completed task, update its status from `[~]` to `[x]`, and append the first 7 characters of the *just-completed commit's* commit hash.
    - **Step 10.2: Write Plan:** Write the updated content back to `plan.md`.

11. **Commit Plan Update:**
    - **Action:** Stage the modified `plan.md` file.
    - **Action:** Commit this change with a descriptive message (e.g., `conductor(plan): Mark task 'Create user model' as complete`).

### Phase Completion Verification and Checkpointing Protocol

**Trigger:** This protocol is executed immediately after a task is completed that also concludes a phase in `plan.md`.

1.  **Announce Protocol Start:** Inform the user that the phase is complete and the verification and checkpointing protocol has begun.

2.  **Ensure Test Coverage for Phase Changes:**
    -   **Step 2.1: Determine Phase Scope:** Identify the files changed in this phase via `git diff --name-only <previous_checkpoint_sha> HEAD`.
    -   **Step 2.2: Verify and Create Tests:** Ensure corresponding test files exist for all modified code files.

3.  **Execute Automated Tests with Proactive Debugging:**
    -   Announce and execute the test command (e.g., `make test` or `go test ./...`).
    -   Attempt a maximum of two fixes for failing tests before asking for guidance.

4.  **Autonomous Manual Verification (AMV) Protocol:**
    -   **Heuristic:** If verification steps are Read-Only, Non-Destructive, and Fast, execute them autonomously.
    -   **Execution:** Close the feedback loop silently and log outcomes in Git notes.
    -   **Fallback:** If AMV fails, attempt one silent fix before requesting user intervention.

5.  **Propose Detailed Manual Verification Plan (if non-AMV):**
    -   Generate actionable steps for the user if verification is complex or destructive.
    -   **CRITICAL:** Send a Discord notification *before* asking for confirmation.

6.  **Await Explicit User Feedback:**
    -   Pause and wait for the user's response for non-AMV verification.

7.  **Create Checkpoint Commit:**
    -   Perform a checkpoint commit (e.g., `feat: Checkpoint end of Phase X`).

8.  **Attach Auditable Verification Report using Git Notes:**
    -   Attach test results, AMV outcomes, and user confirmations to the checkpoint commit.

9.  **Get and Record Phase Checkpoint SHA:**
    -   Record the SHA in `plan.md` using the format `[checkpoint: <sha>]`.

10. **Track Cleanup & Synchronization:**
    - **CRITICAL:** Once a track is archived, activate the `git-sync` skill.
    - **Multi-Branch Sync:** Ensure `git-sync` is executed properly in both Public and Private branches.


11. **Announce Completion:** Confirm the phase is complete and the checkpoint is recorded.

### Quality Gates

Before marking any task complete, verify:

- [ ] All tests pass
- [ ] Code coverage meets requirements (>80%)
- [ ] Code follows project's code style guidelines (in `conductor/code_styleguides/`)
- [ ] All public functions/methods are documented (GoDoc style)
- [ ] Type safety is enforced (Go types)
- [ ] No linting or static analysis errors (`go vet`, `golangci-lint`)
- [ ] Documentation updated if needed
- [ ] No security vulnerabilities introduced

## Development Commands

### Setup
```bash
go mod tidy
make build
```

### Daily Development
```bash
go run .
make test
go fmt ./...
```

### Before Committing
```bash
make lint
make test-verbose
```

## Testing Requirements

### Unit Testing
- Every module must have corresponding tests.
- Use appropriate test setup/teardown mechanisms (e.g., fixtures, beforeEach/afterEach).
- Mock external dependencies.
- Test both success and failure cases.

### Integration Testing
- Test complete user flows
- Verify database transactions
- Test authentication and authorization
- Check form submissions

### Mobile Testing
- Test on actual iPhone when possible
- Use Safari developer tools
- Test touch interactions
- Verify responsive layouts
- Check performance on 3G/4G

## Code Review Process

### Self-Review Checklist
Before requesting review:

1. **Functionality**
   - Feature works as specified
   - Edge cases handled
   - Error messages are user-friendly

2. **Code Quality**
   - Follows style guide
   - DRY principle applied
   - Clear variable/function names
   - Appropriate comments

3. **Testing**
   - Unit tests comprehensive
   - Integration tests pass
   - Coverage adequate (>80%)

4. **Security**
   - No hardcoded secrets
   - Input validation present
   - SQL injection prevented
   - XSS protection in place

5. **Performance**
   - Database queries optimized
   - Images optimized
   - Caching implemented where needed

6. **Mobile Experience**
   - Touch targets adequate (44x44px)
   - Text readable without zooming
   - Performance acceptable on mobile
   - Interactions feel native

## Commit Guidelines

### Message Format
```
<type>(<scope>): <description>

[optional body]

[optional footer]
```

### Types
- `feat`: New feature
- `fix`: Bug fix
- `docs`: Documentation only
- `style`: Formatting, missing semicolons, etc.
- `refactor`: Code change that neither fixes a bug nor adds a feature
- `test`: Adding missing tests
- `chore`: Maintenance tasks

### Examples
```bash
git commit -m "feat(auth): Add remember me functionality"
git commit -m "fix(posts): Correct excerpt generation for short posts"
git commit -m "test(comments): Add tests for emoji reaction limits"
git commit -m "style(mobile): Improve button touch targets"
```

## Definition of Done

A task is complete when:

1. All code implemented to specification
2. Unit tests written and passing
3. Code coverage meets project requirements
4. Documentation complete (if applicable)
5. Code passes all configured linting and static analysis checks
6. Works beautifully on mobile (if applicable)
7. Implementation notes added to `plan.md`
8. Changes committed with proper message
9. Git note with task summary attached to the commit

## Emergency Procedures

### Critical Bug in Production
1. Create hotfix branch from main
2. Write failing test for bug
3. Implement minimal fix
4. Test thoroughly including mobile
5. Deploy immediately
6. Document in plan.md

### Data Loss
1. Stop all write operations
2. Restore from latest backup
3. Verify data integrity
4. Document incident
5. Update backup procedures

### Security Breach
1. Rotate all secrets immediately
2. Review access logs
3. Patch vulnerability
4. Notify affected users (if any)
5. Document and update security procedures

## Deployment Workflow

### Pre-Deployment Checklist
- [ ] All tests passing
- [ ] Coverage >80%
- [ ] No linting errors
- [ ] Mobile testing complete
- [ ] Environment variables configured
- [ ] Database migrations ready
- [ ] Backup created

### Deployment Steps
1. Merge feature branch to main
2. Tag release with version
3. Push to deployment service
4. Run database migrations
5. Verify deployment
6. Test critical paths
7. Monitor for errors

### Post-Deployment
1. Monitor analytics
2. Check error logs
3. Gather user feedback
4. Plan next iteration

## Continuous Improvement

- Review workflow weekly
- Update based on pain points
- Document lessons learned
- Optimize for user happiness
- Keep things simple and maintainable
