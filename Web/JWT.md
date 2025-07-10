### JWT (JSON Web Token)

> The server generates a token that the client stores and presents with each request. It's a stateless method, meaning the server doesn't need to keep a record of the token.

![JWT](https://media2.dev.to/dynamic/image/width=800%2Cheight=%2Cfit=scale-down%2Cgravity=auto%2Cformat=auto/https%3A%2F%2Fdev-to-uploads.s3.amazonaws.com%2Fuploads%2Farticles%2Ffydm4uic65eaiv956hyp.png)

#### JWT Structure

> - **Header**: Specifies the token type (JWT) and the signing algorithm (e.g., HMAC SHA256).
> - **Payload**: Contains the claims, which are statements about an entity (e.g., user data).
> - **Signature**: Ensures the token hasn't been tampered with and can be verified using the secret key.

HMAC stands for hash based message authentication code

![JWT Structure](https://media2.dev.to/dynamic/image/width=800%2Cheight=%2Cfit=scale-down%2Cgravity=auto%2Cformat=auto/https%3A%2F%2Fdev-to-uploads.s3.amazonaws.com%2Fuploads%2Farticles%2Fxceigflhs68dw1yzm26e.jpg)

#### JWT Claims

[To learn about jwt claims in RFC](https://datatracker.ietf.org/doc/html/rfc7519#section-4)

JWT claims are pieces of information asserted about a subject. For example, an ID token (which is always a JWT) can contain a claim called name that asserts that the name of the user authenticating is "John Doe". In a JWT, a claim appears as a name/value pair where the name is always a string and the value can be any JSON value.

Registered Claims is  a structure of JWT claims registered by IANA and defined by RFC. It ensures interoperability
.

### Header

The header typically consists of two parts: the type of the token, which is JWT, and the signing algorithm being used, such as HMAC SHA256 or RSA.

```json
{
  "alg": "HS256",
  "typ": "JWT"
}
```

### Payload

The second part of the token is the payload, which contains the claims. Claims are statements about an entity (typically, the user) and additional data. There are three types of claims: registered, public, and private claims.

```json
{
  "sub": "1234567890",
  "name": "John Doe",
  "admin": true
}
```

### Signature

To create the signature part you have to take the encoded header, the encoded payload, a secret, the algorithm specified in the header, and sign that.

For example, if you want to use the HMAC SHA256 algorithm, the signature will be created in the following way:

```json
HMACSHA256(
  base64UrlEncode(header) + "." +
  base64UrlEncode(payload),
  secret)
```

The signature is used to verify the message wasn't changed along the way, and, in the case of tokens signed with a private key, it can also verify that the sender of the JWT is who it says it is.