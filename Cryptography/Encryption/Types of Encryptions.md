# Types of Encryptions

## Symmetric

Uses the same password for encryption and decryption.

- DES was adopted as a standard in 1977 and uses a 56-bit key. With the advancement in computing power, in 1999, a DES key was successfully broken in less than 24 hours, motivating the shift to 3DES.
- 3DES is DES applied three times; consequently, the key size is 168 bits, though the effective security is 112 bits. 3DES was more of an ad-hoc solution when DES was no longer considered secure. 3DES was deprecated in 2019 and should be replaced by AES; however, it may still be found in some legacy systems.
- AES was adopted as a standard in 2001. Its key size can be 128, 192, or 256 bits.

## Asymmetric 

Uses private and public keys.

Examples are RSA, Diffie-Hellman, and Elliptic Curve cryptography (ECC). 

Asymmetric encryption is based on a particular group of mathematical problems that are easy to compute in one direction but extremely difficult to reverse. In this context, extremely difficult means practically infeasible. For example, we can rely on a mathematical problem that would take a very long time, for example, millions of years, to solve using today’s technology.