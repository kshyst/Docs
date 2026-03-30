# Encryption and Signing

## Encrypting

As described in previous section, a public and private key pair can be used to encrypt or sign messages. There is an implementation of this method called gpg which can be used on Linux (and other machiens) to perform these tasks. First you need to generate a key:

```shell
gpg --gen-key
```

Now the key is created in the `~/.gnupg` directory and we can share our public key with others:

```shell
gpg --list-keys
```

> the exported key is in binary format, add `-a` to armor it in ascii

Importing it on the other side: 

```shell
gpg --import jadi.pub.key
```

Revoking in case of compromised key:

```shell
gpg --output jadi.revoke.asc --gen-revoke jadijadi@gmail.com
```

To encrypt using gpg:

```shell
echo "I Loved your course! I'll tell all my friends about it." > file.txt
gpg --out file.txt.encrypted --recipient jadijdai@gmail.com --encrypt file.txt
```

Decrypt:

```shell
gpg --out out.txt --decrypt file.txt.encrypted
```

## Signing

Why it works:

- will be able to open the file using my Public Key
- Will be sure that I have signed this because only I has access to the private key of the public key they used to open the file.

```shell
echo "I'm Jadi and I'm glad that you reached to the end of your LPIC study" > message-jadi.txt
gpg --output message-jadi.sig --sign message-jadi.txt
```

Now its enough for me to publish the message-jadi.sig to the Internet or send it to someone. If they want to make sure that it is truely comming from me, its enough for them to check my signature (obviously after importing my public key):

```shell
# Check the signature
gpg --verify message-jadi.sig
# Decrypt it
gpg --output message.jadi --decrypt message-jadi.sig
```

Please also note the `--clearsign` option. This option will create a file ending in `.asc` which contains the original un-encrypted (clear text) message alongside the signature. This way non-tech-savy people will be able to read the message too and only if someone wants, she can `--verify` the signature.

### gpg-agent

Just like the `ssh-agent`, `gpg-agent` is a tool which acts like a password manager for you gpg keys. It keeps the keys in the memory so you wont need to provide password on every single use.