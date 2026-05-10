# SSH(Secure Shell)

## Authenticating the Server

```shell
ssh 10.10.244.173
The authenticity of host '10.10.244.173 (10.10.244.173)' can't be established.
ED25519 key fingerprint is SHA256:lLzhZc7YzRBDchm02qTX0qsLqeeiTCJg5ipOT0E/YM8.
This key is not known by any other name.
Are you sure you want to continue connecting (yes/no/[fingerprint])? yes
Warning: Permanently added '10.10.244.173' (ED25519) to the list of known hosts.
```

### ED25519

ED25519 is the public-key algorithm used for digital signature generation and verification in this example. Our SSH client didn’t recognise this key and is asking us to confirm whether we want to continue with the connection. This warning is because a man-in-the-middle attack is probable; a malicious server might have intercepted the connection and replied, pretending to be the target server.

In this case, the user must authenticate the server, i.e., confirm the server’s identity by checking the public key signature. Once you answer with “yes”, the SSH client will record this **public key signature for this host**. In the future, it will connect you silently **unless this host replies with a different public key**.

## SSH Key Fingerprints

A **fingerprint** is a short, unique string of characters used to identify a much longer public key. Because public keys (like RSA 4096-bit keys) are too long for humans to easily compare, SSH uses a cryptographic hash of the public key to create a fingerprint.

- **Format**: Modern SSH clients typically show fingerprints using the **SHA256** hash algorithm (e.g., `SHA256:lLzhZc...`).
- **Purpose**: When you connect to a server for the first time, the client displays the server's public key fingerprint. You should verify this fingerprint against a trusted source (like the server administrator or a web dashboard) to ensure you are connecting to the correct machine.
- **Verification**: If the fingerprint shown by your client doesn't match the expected one, it's a sign of a **Man-in-the-Middle (MITM)** attack. Someone might be intercepting your connection and presenting their own key to steal your credentials.

!!! info "Checking a local key's fingerprint"
    You can check the fingerprint of your own local public key using:
    ```bash
    ssh-keygen -lf ~/.ssh/id_ed25519.pub
    ```

## Authenticating the Client

Normally the user uses username and password. But it's lame.

So we use public and private keys.

Alongside RSA these are other algorithms to encrypt keys:

- `DSA (Digital Signature Algorithm)` is a public-key cryptography algorithm specifically designed for digital signatures.
- `ECDSA (Elliptic Curve Digital Signature Algorithm)` is a variant of DSA that uses elliptic curve cryptography to provide smaller key sizes for equivalent security.
- `ECDSA-SK (ECDSA with Security Key)` is an extension of ECDSA. It incorporates hardware-based security keys for enhanced private key protection.
- `Ed25519` is a public-key signature system using EdDSA (Edwards-curve Digital Signature Algorithm) with Curve25519.
- `Ed25519-SK (Ed25519 with Security Key)` is a variant of Ed25519. Similar to ECDSA-SK, it uses a hardware-based security key for improved private key protection.

```shell
ssh-keygen -t ed25519
```

## Private Keys

It’s very important to mention that the passphrase used to decrypt the private key doesn’t identify you to the server at all; it only decrypts the SSH private key. The passphrase is never transmitted and never leaves your system.

The permissions must be set up correctly to use a private SSH key; otherwise, your SSH client will ignore the file with a warning. Only the owner should be able to read or write to the private key (600 or stricter). ssh -i privateKeyFileName user@host is how you specify a key for the standard Linux OpenSSH client.

### Keys Trusted by the Remote Host
The `~/.ssh` folder is the default place to store these keys for OpenSSH. The authorized_keys (note the US English spelling) file in this directory holds public keys that are allowed access to the server if key authentication is enabled. By default on many Linux distributions, key authentication is enabled as it is more secure than using a password to authenticate. Only key authentication should be accepted if you want to allow SSH access for the root user.

## Case Study: GitHub Authentication

GitHub uses SSH public key authentication to securely identify users during `git push` and `git pull` operations. This removes the need to enter your username and password for every interaction.

### The Authentication Flow

1. **Key Generation**: You generate an SSH key pair (e.g., using `ssh-keygen -t ed25519`) on your local machine.
2. **Public Key Upload**: You copy the content of your **public key** (e.g., `~/.ssh/id_ed25519.pub`) and add it to your GitHub account under **Settings > SSH and GPG keys**.
3. **The Challenge**: When you run `git push`, your Git client connects to `git@github.com`. GitHub looks up the public keys associated with your account.
4. **Signature Verification**: GitHub sends a "challenge" (a piece of random data) to your client. Your SSH agent signs this challenge using your **private key** and sends the signature back.
5. **Access Granted**: GitHub verifies the signature using your stored public key. If it matches, GitHub knows you possess the corresponding private key and authorizes the push/pull operation.

!!! note
    GitHub never sees your private key. The authentication happens entirely through the cryptographic proof that you own the private key corresponding to the public key you uploaded.