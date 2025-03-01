# Authentication and Security

## Authentication and Authorization Difference

### Authentication

> Authentication is the process of verifying the identity of a user or system. It is used to ensure that the user or system is who they claim to be. Authentication is typically done by providing a username and password, but it can also involve other methods such as biometric authentication, two-factor authentication, or multi-factor authentication.


### Authorization

> Authorization is the process of determining what actions a user or system is allowed to perform. It is used to control access to resources or services based on the identity of the user or system. Authorization is typically done by assigning roles or permissions to users or systems, which define what actions they are allowed to perform.

## Token-Based Authentication

### Authentication Token

> A Token is a computer-generated code that acts as a digitally encoded signature of a user. They are used to authenticate the identity of a user to access any website or application network.

> - Physical token: A Physical token use a tangible device to store the information of a user. Here, the secret key is a physical device that can be used to prove the user’s identity. Two elements of physical tokens are hard tokens and soft tokens. Hard tokens use smart cards and USB to grant access to the restricted network like the one used in corporate offices to access the employees. Soft tokens use mobile or computer to send the encrypted code (like OTP) via authorized app or SMS.
> 
> - Web token: The authentication via web token is a fully digital process. Here, the server and the client interface interact upon the user’s request. The client sends the user credentials to the server and the server verifies them, generates the digital signature, and sends it back to the client. Web tokens are popularly known as JSON Web Token (JWT), a standard for creating digitally signed tokens.

### Token-based Authentication

> Token-based authentication is a two-step authentication strategy to enhance the security mechanism for users to access a network. The users once register their credentials, receive a unique encrypted token that is valid for a specified session time. During this session, users can directly access the website or application without login requirements. It enhances the user experience by saving time and security by adding a layer to the password system.
>
> A token is stateless as it does not save information about the user in the database. This system is based on cryptography where once the session is complete the token gets destroyed. So, it gets the advantage against hackers to access resources using passwords.  

### Elements of Token-based Authentication

> - **User**: A person who intends to access the network carrying his/her username & password.
> - **Client-Server**: A client is a front-end login interface where the user first interacts to enroll for the restricted resource.
> - **Authorization server**: A backend unit handling the task of verifying the credentials, generating tokens, and send to the user.
> - **Resource server**: It is the entry point where the user enters the access token. If verified, the network greets users with a welcome note.

![Token-based Authentication](https://media.geeksforgeeks.org/wp-content/uploads/20220301152611/TokenbasedAuthentication3-300x191.png)

### How Token-based Authentication Works

> - **Request**: The user intends to enter the service with login credentials on the application or the website interface. The credentials involve a username, password, smartcard, or biometrics
> - **Verification**: The login information from the client-server is sent to the authentication server for verification of valid users trying to enter the restricted resource. If the credentials pass the verification the server generates a secret digital key to the user via HTTP in the form of a code. The token is sent in a JWT open standard format which includes- 
>> - Header: Contains the type of token and the algorithm used to encrypt the token.
>> - Payload: Contains the user’s information and the token’s expiration time.
>> - Signature: It verifies the authenticity of the user and the messages transmitted.
> - **Token**: The user receives the token code and enters it into the resource server to grant access to the network. The access token has a validity of 30-60 seconds and if the user fails to apply it can request the Refresh token from the authentication server. There’s a limit on the number of attempts a user can make to get access. This prevents brute force attacks that are based on trial and error methods.  
> - **Storage**: Once the resource server validated the token and grants access to the user, it stores the token in a database for the session time you define. The session time is different for every website or app. For example, Bank applications have the shortest session time of about a few minutes only.


## CSRF and CORS

### CSRF Attack

> CSRF (Cross-Site Request Forgery) is an attack that impersonates a trusted user and sends a website unwanted commands.
> This can be done, for example, by including malicious parameters in a URL behind a link that purports to go somewhere else:

```html
<img src="https://www.example.com/index.php?action=delete&id=123" />
```

> For users who have modification permissions on https://www.example.com, the \<img> element executes action on https://www.example.com without their noticing, even if the element is not at https://www.example.com.

#### CSRF Prevention

> - **SameSite cookies**: SameSite cookies allow you to specify that you want the browser to only send cookies in response to requests originating from the cookie's origin site, for example. This makes the CSRF attack fail because the malicious commands will not have cookies sent with them and therefore cannot authenticate as the user. The available values are:
>> - Strict: Causes the browser to only send the cookie in response to requests originating from the cookie's origin site.
>> - Lax: Similar to Strict, except the browser also sends the cookie when the user navigates to the cookie's origin site (even if the user is coming from a different site).
>> - Specifies that cookies are sent on both originating and cross-site requests.

> - **anti-CSRF tokens**: Anti-CSRF tokens prevent CSRF attacks by requiring the existence of a secret, unique, and unpredictable token on all destructive changes. These tokens can be set for an entire user session, rotated on a regular basis, or be created uniquely for each request.

### CORS

> CORS (Cross-Origin Resource Sharing) is a system, consisting of transmitting HTTP headers, that determines whether browsers block frontend JavaScript code from accessing responses for cross-origin requests.
> 
> The same-origin security policy forbids cross-origin access to resources. But CORS gives web servers the ability to say they want to opt into allowing cross-origin access to their resources.

>Cross-Origin Resource Sharing (CORS) is an HTTP-header based mechanism that allows a server to specify who can access its resources from a different origin, effectively permitting or restricting web pages to request resources from different domains, schemes, or ports. This is necessary for browsers to safely allow cross-origin requests while maintaining security.
#### CORS Headers

> - **Access-Control-Allow-Origin**: This header specifies which origins are allowed to access the resource. It can be set to a specific origin, *, or null.
> - **Access-Control-Allow-Methods**: This header specifies which HTTP methods are allowed when accessing the resource.
> - **Access-Control-Allow-Headers**: This header specifies which headers are allowed when accessing the resource.
> - **Access-Control-Allow-Credentials**: This header specifies whether credentials are allowed when accessing the resource.
> - **Access-Control-Expose-Headers**: This header specifies which headers are exposed to the browser when accessing the resource.
> - **Access-Control-Max-Age**: This header specifies how long the results of a preflight request can be cached.
> - **Access-Control-Request-Method**: This header specifies the HTTP method that will be used when accessing the resource.
> - **Access-Control-Request-Headers**: This header specifies the headers that will be used when accessing the resource.
> - **Origin**: This header specifies the origin of the request.


## OAuth and OpenID

### OAuth

> OAuth (Open Authorization) is an open standard for access delegation, commonly used as a way for Internet users to grant websites or applications access to their information on other websites but without giving them the passwords.

#### OAuth Flow

> - **User Authorization**: The user authorizes the client to access their resources.
> - **Request Token**: The client requests a token from the authorization server.
> - **Access Token**: The authorization server sends an access token to the client.
> - **Resource Access**: The client uses the access token to access the user's resources.
> - **Token Expiration**: The access token expires after a certain period of time.
> - **Refresh Token**: The client can request a new access token using a refresh token.
> - **Token Revocation**: The user can revoke the access token at any time.
> - **Token Scope**: The access token can have a limited scope.
> - **Token Type**: The access token can be a bearer token or a JWT.

### OpenID

> OpenID is an open standard and decentralized authentication protocol that allows users to be authenticated by co-operating sites (known as relying parties, or RP) using a third-party service, eliminating the need for webmasters to provide their own ad hoc login systems, and allowing users to log in to multiple unrelated websites without having to have a separate identity and password for each.

#### OpenID Flow

> - **User Authentication**: The user authenticates with the OpenID provider.
> - **OpenID Request**: The client requests an OpenID token from the OpenID provider.
> - **OpenID Token**: The OpenID provider sends an OpenID token to the client.
> - **User Authorization**: The user authorizes the client to access their resources.
> - **Resource Access**: The client uses the OpenID token to access the user's resources.
> - **Token Expiration**: The OpenID token expires after a certain period of time.
> - **Refresh Token**: The client can request a new OpenID token using a refresh token.
> - **Token Revocation**: The user can revoke the OpenID token at any time.
> - **Token Scope**: The OpenID token can have a limited scope.
> - **Token Type**: The OpenID token can be a bearer token or a JWT.
> - **User Info**: The OpenID token can contain user information.

## JWT and Session  

### JWT (JSON Web Token)

> The server generates a token that the client stores and presents with each request. It's a stateless method, meaning the server doesn't need to keep a record of the token.

![JWT](https://media2.dev.to/dynamic/image/width=800%2Cheight=%2Cfit=scale-down%2Cgravity=auto%2Cformat=auto/https%3A%2F%2Fdev-to-uploads.s3.amazonaws.com%2Fuploads%2Farticles%2Ffydm4uic65eaiv956hyp.png)

#### JWT Structure

> - **Header**: Specifies the token type (JWT) and the signing algorithm (e.g., HMAC SHA256).
> - **Payload**: Contains the claims, which are statements about an entity (e.g., user data).
> - **Signature**: Ensures the token hasn't been tampered with and can be verified using the secret key.

![JWT Structure](https://media2.dev.to/dynamic/image/width=800%2Cheight=%2Cfit=scale-down%2Cgravity=auto%2Cformat=auto/https%3A%2F%2Fdev-to-uploads.s3.amazonaws.com%2Fuploads%2Farticles%2Fxceigflhs68dw1yzm26e.jpg)

#### JWT Advantages

> - **Scalability**: Due to their stateless nature, JWTs are ideal for distributed systems.
> - **Flexibility**: They can be used across different domains and applications.
> - **Security**: When properly implemented, they provide a secure way to handle user authentication.

### Session-Based Authentication

> it's stateful. The server creates a session for the user and stores session data on the server-side. The client holds only a session identifier, typically in a cookie.

#### How Session-Based Authentication Works

> - **User Authentication**: The user logs in with their credentials.
> - **Session Creation**: Upon successful authentication, the server creates a session record with a unique identifier, user identifier, session start time, expiry, and possibly additional context like IP address and User Agent. Stores that in Database.
> - **Cookie Storage**: This session identifier is sent back and stored as a cookie in the user’s browser.
> - **Session Verification**: Each request from the user’s browser includes this cookie, then server validates the session by querying to Database. If valid, the request is processed.

#### Session-Based Authentication Advantages

> - **Simplicity and Reliability**: The server’s session record acts as a centralized truth source, making it straightforward to manage user sessions.
> - **Revocation Effiency**: Access can be quickly revoked by deleting or invalidating the session record, ensuring up-to-date session validity.

#### Session-Based Authentication Disadvantages

> - **Performance Issue at Scale**: The dependency on database interactions for every session validation can introduce latency, particularly for high-traffic applications.
> - **Latency in Dynamic Environments**: In applications with dynamic clients, this latency can impact user experience, making session-based authentication less ideal in such scenarios.

![Session-Based Authentication](https://media2.dev.to/dynamic/image/width=800%2Cheight=%2Cfit=scale-down%2Cgravity=auto%2Cformat=auto/https%3A%2F%2Fdev-to-uploads.s3.amazonaws.com%2Fuploads%2Farticles%2F52i0zb81mjlfp95b27cw.jpg)