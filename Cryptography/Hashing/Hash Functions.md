# Hash Functions
Hash functions are different from encryption. There is no key, and it’s meant to be impossible (or computationally impractical) to go from the output back to the input.

A hash function takes some input data of any size and creates a summary or digest of that data. The output has a fixed size. It’s hard to predict the output for any input and vice versa.

## Good Hash

Good hashing algorithms will be relatively fast to compute and prohibitively slow to reverse, i.e., go from the output and determine the input. Any slight change in the input data, even a single bit, should cause a significant change in the output.

This example there are 2 files which are one alphabet and consecutive. Their MD5, SHA1 and SHA256 hashes are so different.

```shell

strategos@g5000 ~> cat file1.txt 
T⏎
strategos@g5000 ~> cat file2.txt 
U⏎   
strategos@g5000 ~> hexdump -C file1.txt 
00000000  54                                                |T|
00000001
strategos@g5000 ~> hexdump -C file2.txt 
00000000  55                                                |U|
00000001
strategos@g5000 ~> md5sum *.txt
b9ece18c950afbfa6b0fdbfa4ff731d3  file1.txt
4c614360da93c0a041b22e537de151eb  file2.txt
strategos@g5000 ~> sha1sum *.txt
c2c53d66948214258a26ca9ca845d7ac0c17f8e7  file1.txt
b2c7c0caa10a0cca5ea7d69e54018ae0c0389dd6  file2.txt
strategos@g5000 ~> sha256sum *.txt
e632b7095b0bf32c260fa4c539e9fd7b852d0de454e9be26f24d0d6f91d069d3  file1.txt
a25513c7e0f6eaa80a3337ee18081b9e2ed09e00af8531c8f7bb2542764027e7  file2.txt
```

## Why Hashing

Ensures data integrity and password confidentiality

## Hash Collision

Hash functions are designed to be uncollisionable and it should prevent an attacker to engineer
a collision. But pigeon-hole effect says otherwise.

MD5 hash function is 16 Bytes