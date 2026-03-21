# Archiving and Compressing

- Archiving: Bundling multiple files and directories into a single file (e.g., tar). It does not reduce the file size.
- Compressing: Using mathematical algorithms to reduce the size of a single file (e.g., gzip, bzip2).

## Compressing

### gzip and gunzip
Creates .gz

- **Algorithm**: DEFLATE
- **Pros**: The undisputed standard in the Unix/Linux world. It is incredibly fast and universally available on almost every server or OS by default.
- **Cons**: The compression ratio is only average compared to modern tools.
- **Best for**: Web server compression, log rotation, and sharing standard files where speed and compatibility matter most.

```shell
gzip salam.txt
gunzip salam.txt
```

### bzip2 and bunzip2

- **Algorithm**: Burrows-Wheeler
- **Pros**: Achieves a significantly better compression ratio than gzip.
- **Cons**: It is heavily CPU-intensive and much slower than gzip at both compressing and decompressing.
- **Best for**: Historically used for source code distribution, but it is largely becoming obsolete, replaced by xz and zstd.

### xz and unxz

- **Algorithm**: LZMA2
- **Pros**: Extremely high compression ratio (creates very small files). Decompression is relatively fast.
- **Cons**: Archiving/compressing takes a very long time and uses a lot of RAM.
- **Best for**: Distributing software packages (like Linux kernels or OS images) where the file is compressed once but downloaded and decompressed millions of times.

> `unxz` is just `xz --decompress`

### zst (zStandard)

- **Algorithm**: Zstandard (developed by Facebook)
- **Pros**: The modern king of compression. It offers a compression ratio similar to gzip but at lightning-fast speeds, and can be tuned to match xz’s compression ratio if needed.
- **Cons**: Not installed by default on older legacy servers (though standard on modern Linux distributions).
- **Best for**: Backups, real-time data compression, and modern system packages (Arch Linux and Ubuntu have migrated to zstd for package management).

## Archiving

### tar

- `cf`: create file
- `xf`: extract file
- `z`: compress with gzip after creating the archive
- `b`: compress with bzip2 after creating the archive

```shell
tar cf mytasks.tar task1.txt task2.txt salam.txt

# Compress (Archive + gzip)
tar -czvf archive.tar.gz /path/to/directory

# Compress (Archive + bzip2)
tar -cjvf archive.tar.bz2 /path/to/directory

# Compress (Archive + xz)
tar -cJvf archive.tar.xz /path/to/directory

# Compress (Archive + zstd)
tar -I zstd -cvf archive.tar.zst /path/to/directory

# Extract any of the above (tar automatically detects the compression)
tar -xvf archive.tar.gz
```

## Archiving + Compressing

### zip (.zip)
- **Pros**: Universal compatibility. It works natively on Windows, macOS, and Linux without installing third-party software.
- **Cons**: Mediocre compression ratio. More importantly, it does not perfectly preserve Linux file permissions (like ownership or executable flags).
- **Best for**: Sending files to non-Linux users (Windows/macOS) or general-purpose sharing.
### rar (.rar)
- **Pros**: Excellent at creating “solid archives”, splitting massive files into smaller chunks (e.g., file.part1.rar, file.part2.rar), and adding “Recovery Records” to fix corrupted downloads.
- **Cons**: It is proprietary software. While extracting (unrar) is often free and available on Linux, creating RAR files requires a paid license and the proprietary rar command-line tool.
- **Best for**: File sharing over unreliable networks (due to recovery records) and splitting files.
### 7z (7-Zip) (.7z)
- **Algorithm**: LZMA / LZMA2
- **Pros**: Open-source and offers one of the highest compression ratios available, outperforming zip and rar in file size reduction.
- **Cons**: Like zip, it does not preserve Linux file permissions or symlinks perfectly natively (unless you tar the files first and then 7z the tarball).
- **Best for**: Long-term cold storage of files on Windows, or maximizing storage space when permissions don’t matter.

> Ration is `uncompressed size / compressed size` so higher is smaller size.