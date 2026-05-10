# Documentation Workflow

## Phase 1: Research
1. **Infrastructure Audit**: Read `mkdocs.yml`, `.gitlab-ci.yml`, and `ansible/` files to understand how documentation is built and deployed.
2. **Domain Discovery**: Search the codebase for existing topics or configuration related to the target documentation (e.g., search for `docker-compose.yml` files when writing Docker docs).
3. **Constraint Identification**: Identify network mirrors, specific tool versions, or local configurations that must be documented.

## Phase 2: Planning
1. **Define Scope**: List the specific files to be created (e.g., `Introduction.md`, `Deployment.md`, `Best Practices.md`).
2. **Create a Plan**: Use a dedicated plan file or shared summary to outline the structure and key points for each document.
3. **Solicit Feedback**: Confirm the plan with the user before proceeding to high-volume generation.

## Phase 3: Generation
1. **Bulk Creation**: Use the `generalist` subagent for creating multiple files simultaneously.
2. **Context Injection**: Provide the subagent with clear instructions, style guidelines, and the approved plan.
3. **Iterative Refinement**: For complex files, perform surgical edits using `replace` after the initial generation.

## Phase 4: Verification & Cleanup
1. **Structure Check**: Use `list_directory` to ensure all files are in the correct locations.
2. **Style Review**: Read samples of the generated content to ensure it follows the `style-guide.md`.
3. **Integration**: Ensure new documents are correctly indexed or linked if manual navigation is required.
4. **Refactoring**: Merge sparse files and delete redundant content to keep the documentation lean.
