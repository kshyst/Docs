# API

> APIs are mechanisms that enable two software components to communicate with each other using a set of definitions and protocols. For example, the weather bureau’s software system contains daily weather data. The weather app on your phone “talks” to this system via APIs and shows you daily weather updates on your phone.

## What does API stand for?

> API stands for Application Programming Interface. In the context of APIs, the word Application refers to any software with a distinct function. Interface can be thought of as a contract of service between two applications. This contract defines how the two communicate with each other using requests and responses. Their API documentation contains information on how developers are to structure those requests and responses.

## Types of APIs

### 1. Open APIs

> Open APIs are publicly available APIs that can be accessed by any developer. They are also known as external APIs or public APIs. Open APIs are designed to be consumed by external developers, and they are often used to enable third-party integrations with a service or platform.

### 2. Internal APIs

> Internal APIs are APIs that are designed to be consumed by internal developers. They are also known as private APIs. Internal APIs are used to enable communication between different internal services or systems within an organization.


### 3. Partner APIs

> Partner APIs are APIs that are shared with specific external partners. They are designed to be consumed by trusted partners or third-party developers who have a business relationship with the API provider. Partner APIs are often used to enable integration with specific partners or to provide access to premium features or data.


### 4. Composite APIs

> Composite APIs are APIs that combine multiple APIs into a single interface. They are used to simplify complex interactions between multiple services or systems. Composite APIs abstract the complexity of interacting with multiple APIs by providing a unified interface that hides the underlying implementation details.


### 5. RESTful APIs

> RESTful APIs are APIs that adhere to the principles of Representational State Transfer (REST). REST is an architectural style that defines a set of constraints for creating web services. RESTful APIs use standard HTTP methods (such as GET, POST, PUT, DELETE) to perform operations on resources. They are designed to be stateless, scalable, and easy to use.


### 6. Web APIs

> A Web API or Web Service API is an application processing interface between a web server and web browser. All web services are APIs but not all APIs are web services. REST API is a special type of Web API that uses the standard architectural style explained above.


### 7. SOAP APIs

> SOAP (Simple Object Access Protocol) APIs are APIs that use the SOAP protocol for communication. SOAP is a protocol for exchanging structured information in the implementation of web services. SOAP APIs use XML as the message format and can be more complex than RESTful APIs.


### 8. GraphQL APIs

> GraphQL APIs are APIs that use the GraphQL query language for communication. GraphQL is a query language for APIs and a runtime for executing those queries. GraphQL APIs allow clients to request only the data they need, which can reduce the amount of data transferred over the network.


### 9. gRPC APIs

> gRPC APIs are APIs that use the gRPC framework for communication. gRPC is a high-performance, open-source RPC (Remote Procedure Call) framework developed by Google. gRPC APIs use Protocol Buffers as the message format and can be more efficient than RESTful APIs.

## Restful API Structure

> RESTful APIs are designed to be simple, scalable, and easy to use. They use standard HTTP methods (such as GET, POST, PUT, DELETE) to perform operations on resources. RESTful APIs are stateless, meaning that each request from a client contains all the information needed to process the request. This makes RESTful APIs easy to scale and maintain.

> Example of a RESTful API request:

```http

GET /api/v1/users/123 HTTP/1.1

Host: example.com

Authorization

```

