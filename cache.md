## Why
Fetching something over the network is both slow and expensive. Large responses require many roundtrips between the client and server, which delays when they are available and when the browser can process them, and also incurs data costs for the visitor. As a result, the ability to cache and reuse previously fetched resources is a critical aspect of optimizing for performance.

## Definition
A web cache (or HTTP cache) is an information technology for the temporary storage (caching) of web documents, such as HTML pages and images, to reduce bandwidth usage, server load, and perceived lag. A web cache system stores copies of documents passing through it; subsequent requests may be satisfied from the cache if certain conditions are met.

## Mime support

In short terms, any response with any content type is cached, such as:

- application/json
- text/html
- text/plain
- text/xml
- text/javascript (JSONP)
- application/octet-stream
- application/pdf
- image/jpeg
- image/png
- image/gif
- image/bmp
- image/svg+xml
- image/x-icon

------

Example Code:

- https://github.com/get-ion/cache/blob/master/_examples/simple/main.go

> Read more at: https://github.com/get-ion/cache