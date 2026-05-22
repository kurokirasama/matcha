# Specification: Search by Folder Syntax

## Overview
Adds support for scoping searches to specific folders using the `label:folder` or `folder:name` syntax in the search bar.

## Functional Requirements
- **DSL Expansion**: Update the search parser to recognize `label:` and `folder:` prefixes.
- **Search Routing**: Ensure the backend search call uses the specified folder instead of the current one.

## Acceptance Criteria
- [ ] `label:Sent query` searches only the Sent folder.
- [ ] Parser correctly extracts folder names from the query string.
