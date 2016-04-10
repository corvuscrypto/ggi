# GGI

The Go Gateway Interface (GGI) is an experimental universal interface between servers and Go programs so that requests can be handled by modularized Go packages.

Basically, Go packages compile fast. Why not leverage this? This is an experimental project that takes advantage of that fact to handle web requests using the approach CGI basically takes. A request handled by GGI flows as such:

1. Request encountered by the GGI server
2. GGI server looks up registered routes to find a corresponding process
3. GGI server attempts to hand off the request to that process
  * if process is already running, hand off the request and go to step 6
4. If the process is not running the GGI server looks where the go files would be and attempts to compile them
5. The process is then started and the request is handed to it
6. The process returns a response
7. The response gets returned to the request origin

Pretty basic and a bit too simplistic, but it works.

It's still buggy, but it you want to try some things using this idea I have another repository that includes an implementation of GGI for "practical" (in quotes because it's still shit) use. https://github.com/corvuscrypto/gserve

## Route Registration

Route registration is handled by a single function, `RegisterRoute`. This is where you declare which package handles requests

E.g. if I want to have a single-file main package with the file `index.go` handle the top level of my site (i.e. '/') then I would just register the route using `ggi.RegisterRoute("/","index.go")`.

Now let's say I have a larger main package under the directory `test_mod/` and I want the route "/test/" to go there. Easy enough I just use `ggi.RegisterRoute("/test/", "test_mod/")`

## Request Handling

In order to pass off your incoming requests, use the `HandleRequest` method. The way I do this is to add it to my webserver as a handler using the `http.HandleFunc` method. This works well.

## Compilation

I took the lazy route with compiling main packages. I just execute the "go build" command :P. Obviously this causes issues with go packages that have special build flags that need to be applied but I will work on this at some point, just not yet.

## Current Goals

Basically this was step one: Get something, anything working.

Now I'm working on getting it refined so that request data is not just essentially a ping, and responses aren't just

## Benefits

The most obvious benefit is that when you have a go file that won't compile (not really an issue for those that use most modern IDE's and editors) it won't break your entire application.

A less bvious benefit is unexpected process failure due to an unexpected condition. E.g. a nil pointer error that results in an uncaught panic. While I haven't fully implemented the behavior for this situation, the idea that this process can fail without causing your entire application to go down is pretty comforting.

Also. This was done in 1.5 days as of this writing so give me a break eh? ;)
