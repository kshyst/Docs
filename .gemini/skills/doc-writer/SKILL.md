---
name: doc-writer
description: Specialized workflow for researching, planning, and writing comprehensive technical documentation (MkDocs). Use when Gemini CLI needs to create new documentation suites, clean up existing Markdown files, or integrate infrastructure details (CI/CD, Ansible) into the project documentation.
---

# Doc Writer

## Overview
The `doc-writer` skill provides a structured approach to generating and maintaining high-quality technical documentation. It emphasizes research-driven content, consistent styling, and the efficient use of subagents for bulk generation.

## Workflow Decision Tree
1. **Is the documentation new?** → Follow the [Documentation Workflow](references/workflow.md) from Phase 1.
2. **Is it an update to existing docs?** → Research the specific topic, then use `replace` for surgical edits.
3. **Is it a high-volume request (>3 files)?** → Use the `generalist` subagent with Phase 3 instructions.
4. **Is it a cleanup task?** → Follow Phase 4 of the workflow.

## Guidelines
All documentation must adhere to the [Documentation Style Guide](references/style-guide.md).

### Key Constraints to Research
Before writing, always identify:
- **Build System**: Usually MkDocs (check `mkdocs.yml`).
- **Deployment**: Usually GitLab CI + Ansible (check `.gitlab-ci.yml` and `ansible/`).
- **Network Mirrors**: Identify any local PyPI/APT mirrors to ensure accuracy in deployment guides.

## Example Requests
- "Write comprehensive documentation for Kubernetes in DevOps/k8s/"
- "Clean up my existing Docker documentation and fix formatting"
- "Add a guide for Harbor registry and explain the RBAC model"
- "Explain how Watchtower detects new images in the docs"

## Verification
Always verify the directory structure and file contents using `list_directory` and `read_file` (sampling) after any generation task.
