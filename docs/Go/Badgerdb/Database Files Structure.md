# Database Files Structure

- LSM tree files (levels on disk: `.sst`, `.idx`, `.blk`)
- Value log files (`000000.vlog`, `000001.vlog`, …)
- Manifest file (`MANIFEST`, sometimes `MANIFEST.tmp`)
- Lock file (`LOCK`)

## LSM Tree Files

Uses an LSM (Log-Structured-Merge) tree for keys and small values.

```text
000001.sst
000001.blk
000001.idx
000002.sst
000002.blk
000002.idx
...
```

### `.sst` main table file

- keys
- value pointers (to the value log or inline values)
- version timestamps
- tombstones (delete markers)

### `.blk` block index

Contains offsets of all blocks inside the SST file.

This allows Badger to jump directly to a block instead of scanning the `.sst` file.

### `.idx` bloom filters and block metadata

- bloom filters for fast key existence checks
- block-level metadata
- index of index structure


## Value Logs `vlog`

These contain actual value bytes.
Only pointers are stored in the LSM tree (unless a value is very small).

```text
000000.vlog
000001.vlog
000002.vlog
...
```
They contain:

- key hash
- value length
- version
- metadata flags
- the value bytes

## MANIFEST

This is the “brain” describing the current database state.

It contains:

- all active table file numbers
- all active vlog file numbers
- level placement in LSM
- next file ID counter
- flags and options used when creating the DB

## LOCK

Ensures that only one Badger instance can open the DB at the same path.

It is removed automatically when Badger closes cleanly.