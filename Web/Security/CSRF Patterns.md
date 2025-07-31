# CSRF Patterns

## Synchronizer Token Pattern

- The server generates unique token for user and stores it in server's session storage. 
- Token is provided to user commonly via forms.
- For each mutating request user sends the token.

**PROS**
- simple and secure

**CONS**
- Stateful and does not scale well for stateless api

## Double Submit Cookie Pattern

- The server stores the csrf token in user's cookies
- For state changing requests, the client gets the token from cookies and  sends it in request headers
- Server checks if the value in cookies and request header checks out.

**PROS**
- Stateless

**CONS**
- Relies on websites same-origin policy to prevent attackers from reading the cookie
- Not as robust as synchronizer pattern if cookies are accessible via javascript (Should use http-only/secure flags)

## Per Request (Masked) Token Pattern

- A per session base token is stored in client's cookies or session
- For each response the server generates a masked(unique per request) version of the token and provides it to the client via forms or API points.
- The client sends the masked token and server on masks it to verify the token.

**PROS**
- Prevents BREACH attack by never exposing the raw token.
- Multiple tabs or requests don't interfere with each other
- Used in Django, Rails and gorilla/csrf

**CONS**
- Complex to implement

