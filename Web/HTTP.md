# What is HTTP

> HTTP is an extensible protocol which has evolved over time. It is an application layer protocol that is sent over TCP, or over a TLS-encrypted TCP connection, though any reliable transport protocol could theoretically be used. Due to its extensibility, it is used to not only fetch hypertext documents, but also images and videos or to post content to servers, like with HTML form results. HTTP can also be used to fetch parts of documents to update Web pages on demand.

## Components of HTTP-based systems



### Client

> he user-agent is any tool that acts on behalf of the user. This role is primarily performed by the Web browser
> 
> To display a Web page, the browser sends an original request to fetch the HTML document that represents the page. It then parses this file, making additional requests corresponding to execution scripts, layout information (CSS) to display, and sub-resources contained within the page (usually images and videos).
>

### Web Server

>On the opposite side of the communication channel is the server, which serves the document as requested by the client. A server appears as only a single machine virtually; but it may actually be a collection of servers sharing the load (load balancing), or other software (such as caches, a database server, or e-commerce servers), totally or partially generating the document on demand.

### Proxies between client and server

> In reality, there are more computers between a browser and the server handling the request: there are routers, modems, and more. Thanks to the layered design of the Web, these are hidden in the network and transport layers. HTTP is on top, at the application layer. Although important for diagnosing network problems, the underlying layers are mostly irrelevant to the description of HTTP.

![proxies](https://mdn.github.io/shared-assets/images/diagrams/http/overview/client-server-chain.svg)

> Between the Web browser and the server, numerous computers and machines relay the HTTP messages. Due to the layered structure of the Web stack, most of these operate at the transport, network or physical levels, becoming transparent at the HTTP layer and potentially having a significant impact on performance. Those operating at the application layers are generally called proxies. These can be transparent, forwarding on the requests they receive without altering them in any way, or non-transparent, in which case they will change the request in some way before passing it along to the server. Proxies may perform numerous functions:
> - caching (the cache can be public or private, like the browser cache)
> - filtering (like an antivirus scan or parental controls)
> - load balancing (to allow multiple servers to serve different requests)
> - authentication (to control access to different resources)
> - logging (allowing the storage of historical information)


## Stateless vs Sessionless

> HTTP is stateless which means there is no connection between 2 requests to check if they were successfully carried but HTTP can be sessionful meaning  by using cookies and som headers it can share same context or state between requests

## HTTP Connections

> HTTP is an application layer protocol , which means it doesn't rely on how it is trasported. 
> the connection and data transmission are controlled by the transport layer .
> HTTP needs a reliable transport protocol which doesn't loses data on transport.
> the most common types of data transmission protocols are tcp and udp . tcp is reliable and udp isn't.
> 
> The established tcp connection can be persistent and non-persistent.
> 
> The default and legacy tcp mode is non-persistent which opens several tcp connections and closes them after the transfer.
> in the persistent mode there is only one tcp connection which is used several times and only closes if a certain timeout is reached.


## HTTP Flow

> 1 - Open a TCP connection
> 2 - Send an HTTP message

```HTTP
GET / HTTP/1.1
Host: developer.mozilla.org
Accept-Language: fr
```

> 3 - Read the response

```HTTP
HTTP/1.1 200 OK
Date: Sat, 09 Oct 2010 14:28:02 GMT
Server: Apache
Last-Modified: Tue, 01 Dec 2009 20:18:22 GMT
ETag: "51142bc1-7449-479b075b2891b"
Accept-Ranges: bytes
Content-Length: 29769
Content-Type: text/html

<!doctype html>… (here come the 29769 bytes of the requested web page)
```

> 4 - close or reuse the connection for further requests.

## Pipelining

> If HTTP pipelining is activated, several requests can be sent without waiting for the first response to be fully received. HTTP pipelining has proven difficult to implement in existing networks, where old pieces of software coexist with modern versions. HTTP pipelining has been superseded in HTTP/2 with more robust multiplexing requests within a frame.

## HTTP Messages

### Requests
 
![requests](https://mdn.github.io/shared-assets/images/diagrams/http/overview/http-request.svg)

### Responses

![responses](https://mdn.github.io/shared-assets/images/diagrams/http/overview/http-response.svg)
