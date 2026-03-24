# YAML files

### Indentations

Should use 2 space or 4 spaces indentations. Using **Tab** is prohibited

### Comments

Using #

### Data types

- strings: Using " " or ' '
- numbers: any `float` or `int`
- boolean: `true` or `false`
- null: `null` or `~`
- lists: using `-` just like in md. Nested lists are available too.
```yml
shopping_list:
  - fruits:
      - apple
      - banana
  - vegetables:
      - carrot
      - broccoli 
```
- dictionary: writing multiple key-value pairs under a key. This is used a fucking lot. Nested dicts are also available.
```yml
user:
  name: Bob
  details:
    age: 30
    skills:
      - Python
      - Docker 
```
- multiline strings: using `|` or `>` .
```yml
description: |
  This is a multiline
  text block.
  The new lines are preserved. 

summary: >
  This is a folded multiline string.
  The newlines will be replaced with spaces
  when displayed as a single line.
```
- anchors and aliases: using << and * to point to a setting and & to define a setting. & is anchor and * is alias.
```yml
defaults: &default_settings
  timeout: 30
  retries: 3

production:
  <<: *default_settings
  timeout: 60 
```