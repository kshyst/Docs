# Access and Refresh Tokens

In modern web applications, particularly those using OAuth 2.0 and OpenID Connect, tokens are the primary mechanism for authorizing requests and managing user sessions without requiring constant re-authentication.

## Access Tokens

An **Access Token** is a credential that can be used by an application to access an API. It informs the API that the bearer of the token has been authorized to access specific resources and perform specific actions.

### Characteristics
- **Short-Lived**: Typically expires in minutes or hours (e.g., 15 minutes to 1 hour).
- **Bearer Token**: The most common type; anyone who possesses the token can use it.
- **Self-Contained (JWT)**: Often implemented as a JSON Web Token (JWT), containing user info and permissions (claims) within the token itself.

### Usage
The client includes the access token in the `Authorization` header of HTTP requests:
```http
Authorization: Bearer <access_token>
```

## Refresh Tokens

A **Refresh Token** is a special type of token used to obtain a new access token without requiring the user to log in again.

### Characteristics
- **Long-Lived**: Typically lasts for days, weeks, or even months.
- **Single-Purpose**: Its only job is to request new access tokens from the Authorization Server.
- **Opaque or JWT**: Can be a simple random string (opaque) or a JWT.

### The Refresh Flow
1. **Initial Login**: User authenticates and receives both an Access Token and a Refresh Token.
2. **Accessing Resources**: Client uses the short-lived Access Token for API calls.
3. **Token Expiration**: The API returns a `401 Unauthorized` because the Access Token has expired.
4. **Refreshing**: The Client sends the Refresh Token to the Authorization Server.
5. **New Tokens**: If valid, the server issues a **new** Access Token (and optionally a new Refresh Token).

## Security Considerations

### 1. Storage Strategies
How you store tokens determines your vulnerability to different attacks:

| Storage | Pros | Cons |
| :--- | :--- | :--- |
| **LocalStorage** | Survives page refreshes, easy to use. | Vulnerable to **XSS** (Cross-Site Scripting). |
| **Memory** (JS Variable) | Secure from XSS (mostly). | Lost on page refresh. |
| **HTTP-Only Cookie** | Secure from XSS; automatically sent by browser. | Vulnerable to **CSRF** (Cross-Site Request Forgery). |

**Best Practice**: Store Refresh Tokens in an `HttpOnly`, `Secure`, `SameSite=Strict` cookie to prevent JavaScript from accessing them.

### 2. Token Rotation
To mitigate the risk of a stolen Refresh Token, implement **Refresh Token Rotation**:
- Every time a Refresh Token is used, the server invalidates the old one and issues a new one.
- If a leaked Refresh Token is used by an attacker, the legitimate user's subsequent attempt to refresh will fail, alerting the server to a potential breach.

### 3. Revocation
The Authorization Server should provide an endpoint to revoke tokens:
- **Logout**: Invalidate the Refresh Token on the server side.
- **Security Breach**: Allow users or admins to revoke all active sessions.

## Comparison Table

| Feature | Access Token | Refresh Token |
| :--- | :--- | :--- |
| **Lifetime** | Short (minutes/hours) | Long (days/weeks) |
| **Purpose** | Access Resources | Get New Access Tokens |
| **Visibility** | Sent to Resource Server (API) | Sent only to Auth Server |
| **Risk** | High (if intercepted) | Extremely High (allows long-term access) |

!!! warning
    Never expose Refresh Tokens in your application's UI or logs. They are sensitive credentials equivalent to a password for the duration of their validity.
