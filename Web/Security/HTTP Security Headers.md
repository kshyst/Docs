# Security Headers

Provides mitigation against attacks like `XSS`, `Click Jacking` and ...

## Content-Security-Policy (CSP)

Against XSS. It specifies what type of content from what websites can be used.

```HTTP
Content-Security-Policy: default-src 'self'; script-src 'self' https://cdn.tryhackme.com; style-src 'self'
```

Here self references the host the website is running on.

- `default-src`: which specifies the default policy of self, which means only the current website.
- `script-src`: which specifics the policy for where scripts can be loaded from, which is self along with scripts hosted on https://cdn.tryhackme.com
- `style-src`: which specifies the policy for where style CSS style sheets can be loaded from the current website (self)

## Strict-Transport-Security (HSTS)

Ensures that browser always connect through https.

```HTTP
Strict-Transport-Security: max-age=63072000; includeSubDomains; preload
```

- `preload`: This optional setting allows the website to be included in preload lists. Browsers can use preload lists to enforce HSTS before even having their first visit to a website.

## X-Content-Type-Options

Instructs browser not to guess the MIME type but only use the `Content-Type` header.

- `nosniff`: a directive that instructs browser not to sniff of guess the MIME type.

## Referrer-Policy

Controls the amount of information sent to the destination web server when a use is redirected from a source web server like clicking a hyperlink.

- `Referrer-Policy: no-referrer` : This completely disables any information being sent about the referrer
- `Referrer-Policy: same-origin` : This policy will only send referrer information when the destination is part of the same origin. This is helpful when you want referrer information passed when hyperlinks are within the same website but not outside to external websites.
- `Referrer-Policy: strict-origin` : This policy only sends the referrer as the origin when the protocol stays the same. So, a referrer is sent when an HTTPS connection goes to another HTTPS connection.
- `Referrer-Policy: strict-origin-when-cross-origin` : This is similar to strict-origin except for same-origin requests, where it sends the full URL path in the origin header.