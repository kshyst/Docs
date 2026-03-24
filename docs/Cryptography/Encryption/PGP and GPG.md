# PGP(Pretty Good Privacy)

It’s software that implements encryption for encrypting files, performing digital signing, and more. GnuPG or GPG is an open-source implementation of the OpenPGP standard.

# GPG(Gnu Privacy Guard)
It is a free and open-source encryption software that uses public-key cryptography. GPG can be used to encrypt files and messages, and to sign files and messages. Encryption makes it so that only the intended recipient can decrypt the file or message while signing makes it so that the recipient can verify that the file or message was sent by the person it claims to be from.

## GPG usage

Commonly used in emails to prove confidentiality. Furthermore can be used to sign emails and protect integrity

To generate a new GPG key:

```shell
  gpg --full-gen-key
```

## GPG in CTF
You may need to use GPG to decrypt files in CTFs. With PGP/GPG, private keys can be protected with passphrases in a similar way that we protect SSH private keys. If the key is passphrase protected, you can attempt to crack it using John the Ripper and gpg2john. The key provided in this task is not protected with a passphrase. The man page for GPG can be found online here.

## GPG Backup

When using new computer, you need to import GPG backup to decrypt your messages and datas:

```shell
    gpg --import backup.key
```

And then to decrypt them use:

```shell
    gpg --decrypt confidential_message.gpg
```