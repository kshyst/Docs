# YAML for Docker

YAML (YAML Ain't Markup Language) is a human-readable data serialization language. It is commonly used for configuration files, including Docker Compose and Kubernetes manifests.

## Basic Syntax Rules

- **Indentation**: Use spaces (usually 2 or 4) for indentation. **Tabs are strictly prohibited** and will cause parsing errors.
- **Comments**: Use the `#` symbol for comments.

## Data Types

### Scalars
- **Strings**: Can be unquoted, or enclosed in single (`' '`) or double (`" "`) quotes.
- **Numbers**: Supports integers and floating-point numbers.
- **Booleans**: Represented by `true`/`false` (or `yes`/`no`).
- **Null**: Represented by `null` or the tilde (`~`).

### Collections

#### Lists (Sequences)
Items in a list start with a hyphen (`-`).
```yaml
shopping_list:
  - fruits:
      - apple
      - banana
  - vegetables:
      - carrot
      - broccoli
```

#### Dictionaries (Mappings)
Mappings use a `key: value` format. They are extensively used in Docker Compose files.
```yaml
user:
  name: Bob
  details:
    age: 30
    skills:
      - Python
      - Docker
```

### Advanced Features

#### Multiline Strings
- `|` (Literal): Preserves newlines.
- `>` (Folded): Replaces newlines with spaces.

```yaml
description: |
  This is a multiline
  text block.
  Newlines are preserved.

summary: >
  This is a folded multiline string.
  Newlines will be replaced with spaces
  when parsed as a single line.
```

#### Anchors and Aliases
Anchors (`&`) and Aliases (`*`) allow you to reuse configuration blocks to keep your files DRY (Don't Repeat Yourself). The merge key (`<<`) can be used to import the contents of an anchor.

```yaml
defaults: &default_settings
  timeout: 30
  retries: 3

production:
  <<: *default_settings
  timeout: 60
```