# Glossary

## Context Window

The maximum amount of tokens that an LLM's input, output, remembrance and process at any single given time.

**Token** is not equal to **Word**.

Token ~= 0.75 Word

**Short-Term memory** is an analogy to this. For example humans can only remember the last 3 pages they read.

When exceeding the Context Window:

- **Truncation(Forgetting)**: Dropping the initial instructions from the context window or the beginning of the document.
- Or just simple error