When we talk about Middleware in Go, at its simplest, we are really talking about running code before and/or after our handler code in a HTTP request lifecycle. For example, logging middleware might write the incoming request details to a log, then call the handler code, before writing details about the response to the log. One of the cool things about middleware, if implemented correctly, is that these units are extremely flexible, reusable, and sharable.

Middleware is registered on routes, group of routes, subdomain or globally. The `github.com/get-ion/ion/context.Context`
is keeping a chain of these handlers which are called via `ctx.Next()`: 
```go
func(ctx context.Context){ 
    if condition {
        ctx.Next()
    }
}
```

In short terms: **Middleware is just a chain handlers which can be executed before or after the main handler, can transfer data between handlers and communicate with third-party libraries, they are just functions.**



**To a single route**
```go
app := ion.New()
app.Get("/mypath", myMiddleware1, myMiddleware2, func(ctx context.Context){ ctx.Next() }, myMainHandler, myDoneMiddleware)
```

**To a party of routes or subdomain**
```go

myparty := app.Party("/myparty", myMiddleware1,func(ctx context.Context){}, myMiddleware3)
{
	//....
}

```

**To all routes**
```go
app.Use(func(ctx context.Context){}, myMiddleware2)
```

**To global, all routes on all subdomains on all parties**
```go
app.UseGlobal(func(ctx context.Context){}, myMiddleware2)
```


## Can I use standard net/http handler with ion?

**Yes** you can, just pass the Handler inside the `handlerconv.FromStd` in order to be converted into ion.HandlerFunc and register it as you saw before.

### Convert handler which has the form of `http.Handler/HandlerFunc`

```go
package main

import (
	"github.com/get-ion/ion"
	"github.com/get-ion/ion/context"
	"github.com/get-ion/ion/core/handlerconv"
)

func main() {
	app := ion.New()

	sillyHTTPHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	     println(r.RequestURI)
	})

	sillyConvertedToIon := handlerconv.FromStd(sillyHTTPHandler)
	// FromStd can take (http.ResponseWriter, *http.Request, next http.Handler) too!
	app.Use(sillyConvertedToIon)

	app.Run(ion.Addr(":8080"))
}

```