# Parsers

REST framework includes a number of built-in Parser classes, that allow you to accept requests with various media types. There is also support for defining your own custom parsers, which gives you the flexibility to design the media types that your API accepts.

## How parsers are determined

The set of valid parsers for a view is always defined as a list of classes. When request.data is accessed, REST framework will examine the Content-Type header on the incoming request, and determine which parser to use to parse the request content.

