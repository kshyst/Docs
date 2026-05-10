# Documentation Style Guide

## Tone and Voice
- **Professional and Direct**: Use an instructional, senior-engineer tone. Avoid conversational filler.
- **Active Voice**: Use "Deploy the container" instead of "The container can be deployed."
- **Imperative Form**: Use commands for instructions (e.g., "Run this command", "Configure this file").

## Markdown Formatting
- **Headers**: Use ATX-style headers (`#`, `##`, `###`). Avoid skipping levels.
- **Code Blocks**: ALWAYS include a language tag (e.g., ````bash`, ````yaml`, ````dockerfile`).
- **Lists**: Use consistent bullet points (`-`) or numbered lists (`1.`).
- **Bold/Italics**: Use **bold** for UI elements, technical terms, or emphasis. Use *italics* sparingly.

## MkDocs Specifics
- **Admonitions**: Use `!!! note`, `!!! warning`, `!!! info` for callouts.
- **Links**: Use relative paths for internal documentation links.
- **Images**: Reference images in `img/` or `assets/` folders using standard Markdown syntax.

## Technical Accuracy
- **Verify Commands**: Ensure all shell commands and configuration snippets are technically correct for the target environment.
- **Clarify Concepts**: Explain complex mechanisms (e.g., "Watchtower uses image digests, not just tags").
- **Prerequisites**: Always list prerequisites for deployment or installation guides.
