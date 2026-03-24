# HTTP Cookies

> A cookie (also known as a web cookie or browser cookie) is a small piece of data a server sends to a user's web browser. The browser may store cookies, create new cookies, modify existing ones, and send them back to the same server with later requests. Cookies enable web applications to store limited amounts of data and remember state information; by default the HTTP protocol is stateless.


## Cookie Usages

> - The user sends sign-in credentials to the server, for example via a form submission.
> - If the credentials are correct, the server updates the UI to indicate that the user is signed in, and responds with a cookie containing a session ID that records their sign-in status on the browser.
> - At a later time, the user moves to a different page on the same site. The browser sends the cookie containing the session ID along with the corresponding request to indicate that it still thinks the user is signed in.
> - The server checks the session ID and, if it is still valid, sends the user a personalized version of the new page. If it is not valid, the session ID is deleted and the user is shown a generic version of the page (or perhaps shown an "access denied" message and asked to sign in again).

![COOKIE](https://developer.mozilla.org/en-US/docs/Web/HTTP/Cookies/cookie-basic-example.png)

## Creating Removing and updating cookies

```HTTP
Set-Cookie: <cookie-name>=<cookie-value>


HTTP/2.0 200 OK
Content-Type: text/html
Set-Cookie: yummy_cookie=chocolate
Set-Cookie: tasty_cookie=strawberry
```

### Removal

```HTTP
Set-Cookie: id=a3fWa; Expires=Thu, 31 Oct 2021 07:28:00 GMT;
Set-Cookie: id=a3fWa; Max-Age=2592000
Set-Cookie: id=a3fWa; Max-Age=2592000; Secure
Set-Cookie: id=a3fWa; Max-Age=2592000; Secure; HttpOnly
Set-Cookie: id=a3fWa; Max-Age=2592000; Secure; HttpOnly; SameSite=Strict
```